package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	protos "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	"github.com/jkarlos000/technical-challenge/currency/internal/config"
	"github.com/jkarlos000/technical-challenge/currency/internal/currency"
	"github.com/jkarlos000/technical-challenge/currency/pkg/dbcontext"
	"github.com/jkarlos000/technical-challenge/currency/pkg/log"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

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

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// manage updates with context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create an instance of the Currency server
	c := currency.NewCurrency(currency.NewRepository(dbcontext.New(db), logger, cfg.CurrencyLayerApiKey), logger, ctx)

	// register the currency server
	protos.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// build GRPC server
	address := fmt.Sprintf(":%v", cfg.ServerPort)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		logger.Errorf("got signal %v, attempting graceful shutdown", s)
		gs.GracefulStop()
		cancel()
		wg.Done()
	}()

	err = gs.Serve(l)
	if err != nil {
		logger.Errorf("could not serve: %v", err)
	}
	wg.Wait()
	logger.Info("clean shutdown")
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
