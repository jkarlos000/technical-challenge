package beer

import (
	"encoding/json"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/pagination"
	"net/http"
	"strconv"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o beertesting/fake_service.gen.go . Service
//counterfeiter:generate -o beertesting/fake_repository.gen.go . Repository

// api fakes
//counterfeiter:generate -o ../test/apifakes/beer-apifakes github.com/jkarlos000/technical-challenge/currency/api/proto/v1.CurrencyClient

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/beers/<id>/boxprice", res.boxBeerPriceById)
	r.Get("/beers/<id>", res.searchBeerById)
	r.Get("/beers", res.searchBeers)
	r.Post("/beers", res.addBeers)
}

type resource struct {
	service Service
	logger  log.Logger
}

// swagger:route GET /beers/{beerID} cerveza SearchBeerById
//
// Lista el detalle de la marca de cervezas
//
// Busca una cerveza por su Id
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: beerItemResponse
//       404: notFound
func (r resource) searchBeerById(c *routing.Context) error {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return errors.BadRequest("")
	}
	beer, err := r.service.Get(c.Request.Context(), id)
	if err != nil {
		return err
	}
	return c.Write(beer)
}

// swagger:route GET /beers/{beerID}/boxprice cerveza BoxBeerPriceById
//
// Lista el precio de una caja de cervezas de una marca
//
// Obtiene el precio de una caja de cerveza por su Id
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: beerBoxResponse
//       404: notFound
func (r resource) boxBeerPriceById(c *routing.Context) error {
	currency := c.Query("currency")
	quantity, _ := strconv.Atoi(c.Query("quantity"))
	var id int
	if id, _ = strconv.Atoi(c.Param("id")); id == 0 {
		return errors.BadRequest("")
	}
	beerBox, err := r.service.GetPrice(c.Request.Context(), id, currency, uint32(quantity))
	if err != nil {
		return err
	}
	return c.Write(beerBox)
}

// swagger:route GET /beers cerveza SearchBeers
//
// Lista todas las cervezas
//
// Lista todas las cervezas que se encuentran en la base de datos
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: beersItemResponse
func (r resource) searchBeers(c *routing.Context) error {
	ctx := c.Request.Context()
	filters := make(map[string]interface{})

	// convert JSON string filters to map
	_ = json.Unmarshal([]byte(c.Query("filters")), &filters)
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	beers, err := r.service.Query(ctx, pages.Offset(), pages.Limit(), filters)
	if err != nil {
		return err
	}
	pages.Items = beers
	return c.Write(pages)
}

// swagger:route POST /beers cerveza AddBeers
//
// Ingresa una nueva cerveza
//
// Ingresa una nueva cerveza
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       201: beerItemCreatedResponse
//       400: badRequest
//       409: conflictError
func (r resource) addBeers(c *routing.Context) error {
	var input CreateBeerRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	beer, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(beer, http.StatusCreated)
}
