package beer

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/jkarlos000/technical-challenge/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}


	r.Get("/beers", res.searchBeers)
	r.Get("/beers/<id>", res.searchBeerById)
	r.Get("/beers/<id>/boxprice", res.boxBeerPriceById)
	r.Post("/beers", res.addBeers)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) searchBeerById(c *routing.Context) error {
	panic("implement me")
}

func (r resource) boxBeerPriceById(c *routing.Context) error {
	panic("implement me")
}

func (r resource) searchBeers(c *routing.Context) error {
	panic("implement me")
}

func (r resource) addBeers(c *routing.Context) error {
	panic("implement me")
}
