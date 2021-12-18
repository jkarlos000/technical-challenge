package beer_test

import (
	"context"
	errors2 "errors"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer/beertesting"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/entity"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateBeerRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     beer.CreateBeerRequest
		wantError bool
	}{
		{"success", beer.CreateBeerRequest{
			Name:     "test",
			Brewery:  "testBrewery",
			Country:  "Chile",
			Price:    9.99,
			Currency: "USD",
		}, false},
		{"required", beer.CreateBeerRequest{
			Name: "test",
		}, true},
		{"too long", beer.CreateBeerRequest{Name: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			Brewery:  "testBrewery",
			Country:  "Chile",
			Price:    9.99,
			Currency: "CLP"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	repo := &beertesting.FakeRepository{}
	s := beer.NewService(repo, logger)

	ctx := context.Background()
	now := time.Now()

	beers := []entity.Beer{
		{
			ID:        1,
			Name:      "testName",
			Brewery:   "testBrewery",
			Country:   "testCountry",
			Price:     10.05,
			Currency:  "CLP",
			CreatedAt: now,
			UpdatedAt: &now,
		},
		{
			ID:        2,
			Name:      "testName2",
			Brewery:   "testBrewery2",
			Country:   "testCountry2",
			Price:     9.25,
			Currency:  "CLP",
			CreatedAt: now,
			UpdatedAt: &now,
		},
	}

	repo.CountReturns(0, nil)
	// initial count
	count, _ := s.Count(ctx)
	assert.Equal(t, 0, count)

	// successful creation
	repo.CreateReturns(1, nil)
	repo.GetReturns(beers[0], nil)
	beerItem, err := s.Create(ctx, beer.CreateBeerRequest{
		Name:     "testName",
		Brewery:  "testBrewery",
		Country:  "Chile",
		Price:    9.99,
		Currency: "USD",
	})
	repo.CountReturns(1, nil)
	assert.Nil(t, err)
	assert.Equal(t, "testName", beerItem.Name)
	assert.NotEmpty(t, beerItem.CreatedAt)
	assert.NotEmpty(t, beerItem.UpdatedAt)
	repo.CountReturns(1, nil)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)
	_, err = s.Create(ctx, beer.CreateBeerRequest{
		Name:     "testName",
		Brewery:  "testBrewery",
		Country:  "Chile",
		Price:    9.99,
		Currency: "USDX",
	})
	assert.NotNil(t, err)
	assert.Equal(t, 1, count)
	_, err = s.Create(ctx, beer.CreateBeerRequest{
		Name:     "testName",
		Brewery:  "testBrewery",
		Country:  "Ñameco",
		Price:    9.99,
		Currency: "USD",
	})
	assert.NotNil(t, err)
	repo.CreateReturns(0, errors2.New("have problems with db"))
	_, err = s.Create(ctx, beer.CreateBeerRequest{
		Name:     "testName",
		Brewery:  "testBrewery",
		Country:  "Bolivia",
		Price:    9.99,
		Currency: "BOB",
	})
	assert.NotNil(t, err)

	// Get
	repo.GetReturns(entity.Beer{}, errors2.New("no record found"))
	_, err = s.Get(ctx, 0)
	assert.NotNil(t, err)

	// Query
	repo.QueryReturns(beers, nil)
	_, err = s.Query(ctx, 0, 0, nil)
	assert.Nil(t, err)
	repo.QueryReturns([]entity.Beer{}, errors2.New("problem with repository"))
	_, err = s.Query(ctx, 0, 0, nil)
	assert.NotNil(t, err)

	// GetPrice
	repo.GetPriceReturns(9.99, nil)
	_, err = s.GetPrice(ctx, 1, "BOB", 0)
	assert.NotNil(t, err)
	_, err = s.GetPrice(ctx, 1, "BOBX", 0)
	assert.NotNil(t, err)
	repo.GetReturns(beers[0], nil)
	_, err = s.GetPrice(ctx, 1, "BOB", 0)
	assert.Nil(t, err)
	repo.GetPriceReturns(0.00, errors2.New("problems with data repository"))
	_, err = s.GetPrice(ctx, 1, "BOB", 0)
	assert.Nil(t, err)
}
