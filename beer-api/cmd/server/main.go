package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	f "github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/config"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/healthcheck"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/accesslog"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/dbcontext"
	_ "github.com/jkarlos000/technical-challenge/beer-api/pkg/doc"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	protos "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"time"
)

//go:generate swagger generate spec -o ../../swaggerui/swagger.yml

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// connect to the database
	db, err := dbx.MustOpen("postgres", cfg.DSN)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	db.QueryLogFunc = logDBQuery(logger)
	db.ExecLogFunc = logDBExec(logger)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error(err)
		}
	}()

	m, err := migrate.New(cfg.MigrationURL, cfg.DSN)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error(err)
	}

	// make grpc Dials
	conn, err := grpc.Dial(cfg.CurrencyService, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, dbcontext.New(db), cfg, conn),
	}

	// start the HTTP server with graceful shutdown
	go routing.GracefulShutdown(hs, 10*time.Second, logger.Infof)
	logger.Infof("server %v is running at %v", Version, address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err)
		os.Exit(-1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, db *dbcontext.DB, cfg *config.Config, conn *grpc.ClientConn) http.Handler {
	router := routing.New()

	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)

	healthcheck.RegisterHandlers(router, Version)

	rg := router.Group("/v1")

	beer.RegisterHandlers(rg.Group(""),
		beer.NewService(beer.NewRepository(protos.NewCurrencyClient(conn), db, logger), logger),
		logger,
	)

	// Swaggerui - serve index file
	router.Get("/swaggerui", f.Content("swaggerui/index.html"))

	router.Get("/*", f.Server(f.PathMap{
		"/swaggerui":     "/swaggerui/",
	}))

	return router
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(logger log.Logger) dbx.QueryLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB query error: %v", err)
		}
	}
}

// logDBExec returns a logging function that can be used to log SQL executions.
func logDBExec(logger log.Logger) dbx.ExecLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB execution successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB execution error: %v", err)
		}
	}
}
