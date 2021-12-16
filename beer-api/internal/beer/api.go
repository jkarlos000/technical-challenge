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

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	r.Get("/beers/<id>", res.searchBeerById)
	r.Get("/beers/<id>/boxprice", res.boxBeerPriceById)
	r.Get("/beers", res.searchBeers)
	r.Post("/beers", res.addBeers)
}

type resource struct {
	service Service
	logger  log.Logger
}

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

func (r resource) boxBeerPriceById(c *routing.Context) error {
	panic("implement me")
}

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
