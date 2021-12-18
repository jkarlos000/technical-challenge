package beer_test

import (
	"context"
	"database/sql"
	errors2 "errors"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/beer"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/entity"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/test"
	beerapifakes "github.com/jkarlos000/technical-challenge/beer-api/internal/test/apifakes/beer-apifakes"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	v1 "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "beers")
	cc := &beerapifakes.FakeCurrencyClient{}

	repo := beer.NewRepository(cc, db, logger)

	ctx := context.Background()
	var id int
	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// auxiliary time var
	now := time.Now()
	// create
	id, err = repo.Create(ctx, entity.Beer{
		Name:      "testName",
		Brewery:   "testBrewery",
		Country:   "testCountry",
		Price:     9.09,
		Currency:  "MXN",
		CreatedAt: now,
		UpdatedAt: &now,
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)
	// create
	_, err = repo.Create(ctx, entity.Beer{
		Name:      "testName",
		Brewery:   "testBrewery",
		Country:   "testCountry",
		Price:     9.09,
		Currency:  "MXN",
		CreatedAt: now,
		UpdatedAt: &now,
	})
	assert.NotNil(t, err)
	_, err = repo.Create(ctx, entity.Beer{
		Name:      "testName",
		Brewery:   "testBrewery",
		Country:   "testCountry",
		Price:     99.99,
		Currency:  "MXNF",
		CreatedAt: now,
		UpdatedAt: &now,
	})
	assert.NotNil(t, err)

	// get
	beer, err := repo.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "testName", beer.Name)
	_, err = repo.Get(ctx, 0)
	assert.Equal(t, sql.ErrNoRows, err)

	// query
	beers, err := repo.Query(ctx, 0, count2, nil)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(beers))

	// getPrice of the beer box
	_, err = repo.GetPrice(ctx, 0, "MXN", "CLP")
	assert.Equal(t, sql.ErrNoRows, err)
	cc.GetPriceReturns(&v1.Response{}, errors2.New("no service"))
	_, err = repo.GetPrice(ctx, id, "MXN", "CLP")
	assert.NotNil(t, err)
	cc.GetPriceReturns(&v1.Response{Rate: 99.90}, nil)
	_, err = repo.GetPrice(ctx, id, "MXN", "CLP")
	assert.Nil(t, err)

}
