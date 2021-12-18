package beer_test

import (
	errors2 "errors"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer/beertesting"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/entity"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/test"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	"net/http"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	service := &beertesting.FakeService{}
	beers := []beer.BeerItem{
		{entity.Beer{
			ID:        1,
			Name:      "testName",
			Brewery:   "testBrewery",
			Country:   "testCountry",
			Price:     10.05,
			Currency:  "CLP",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}},
		{entity.Beer{
			ID:        2,
			Name:      "testName2",
			Brewery:   "testBrewery2",
			Country:   "testCountry2",
			Price:     9.25,
			Currency:  "CLP",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}},
	}

	tests := []struct {
		setup func(fakeService *beertesting.FakeService)
		api   test.APITestCase
	}{
		{
			api: test.APITestCase{Name: "get 1", Method: "GET", URL: "/beers/1", WantStatus: http.StatusOK, WantResponse: `*testName*`},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.GetReturns(beers[0], nil)
			},
		},
		{
			api: test.APITestCase{Name: "get param error", Method: "GET", URL: "/beers/1f", WantStatus: http.StatusBadRequest},
			setup: func(fakeService *beertesting.FakeService) {

			},
		},
		{
			api: test.APITestCase{Name: "get unknown", Method: "GET", URL: "/beers/1234", WantStatus: http.StatusNotFound},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.GetReturns(beer.BeerItem{}, errors.NotFound(""))
			},
		},
		{
			api: test.APITestCase{Name: "create ok", Method: "POST", URL: "/beers", Body: `{"name":"testName2"}`, WantStatus: http.StatusCreated, WantResponse: "*testName2*"},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.CreateReturns(beers[1], nil)
			},
		},
		{
			api: test.APITestCase{Name: "create ok count", Method: "GET", URL: "/beers", WantStatus: http.StatusOK, WantResponse: `*"total_count":2*`},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.CountReturns(2, nil)
			},
		},
		{
			api: test.APITestCase{Name: "create input error", Method: "POST", URL: "/beers", Body: `"name":"test"}`, WantStatus: http.StatusBadRequest},
			setup: func(fakeService *beertesting.FakeService) {

			},
		},
		{
			api: test.APITestCase{Name: "create internal error", Method: "POST", URL: "/beers", Body: `{"name":"testName2"}`, WantStatus: http.StatusInternalServerError},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.CreateReturns(beer.BeerItem{}, errors2.New("problem with repository data"))
			},
		},
		{
			api: test.APITestCase{Name: "get all", Method: "GET", URL: "/beers", WantStatus: http.StatusOK, WantResponse: `*"total_count":2*`},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.CountReturns(2, nil)
				fakeService.QueryReturns(beers, nil)
			},
		},
		{
			api: test.APITestCase{Name: "get all internal error", Method: "GET", URL: "/beers", WantStatus: http.StatusInternalServerError},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.CountReturns(0, errors2.New("problem with repository data"))
			},
		},
		{
			api: test.APITestCase{Name: "get all internal error", Method: "GET", URL: "/beers", WantStatus: http.StatusInternalServerError},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.CountReturns(100, nil)
				fakeService.QueryReturns([]beer.BeerItem{}, errors2.New("problem with data unmarshalling"))
			},
		},
		{
			api: test.APITestCase{Name: "get 1 beer box price", Method: "GET", URL: "/beers/1/boxprice?quantity=10&currency=TMT", WantStatus: http.StatusOK, WantResponse: `*price*`},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.GetPriceReturns(beer.BeerBox{PriceTotal: 99.99}, nil)
			},
		},
		{
			api: test.APITestCase{Name: "get box price param error", Method: "GET", URL: "/beers/1f/boxprice?quantity=10&currency=TMT", WantStatus: http.StatusBadRequest},
			setup: func(fakeService *beertesting.FakeService) {

			},
		},
		{
			api: test.APITestCase{Name: "get box price no service", Method: "GET", URL: "/beers/1/boxprice?quantity=10&currency=TMT", WantStatus: http.StatusInternalServerError},
			setup: func(fakeService *beertesting.FakeService) {
				fakeService.GetPriceReturns(beer.BeerBox{}, errors.InternalServerError("No connection with external service"))
			},
		},
	}

	beer.RegisterHandlers(router.Group(""), service, logger)

	for _, tc := range tests {
		tc.setup(service)
		test.Endpoint(t, router, tc.api)
	}
}
