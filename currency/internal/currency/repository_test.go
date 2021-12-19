package currency_test

import (
	"context"
	"github.com/jkarlos000/technical-challenge/currency/internal/currency"
	"github.com/jkarlos000/technical-challenge/currency/internal/entity"
	"github.com/jkarlos000/technical-challenge/currency/internal/test"
	"github.com/jkarlos000/technical-challenge/currency/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "currencies")

	repo := currency.NewRepository(db, logger, "")

	ctx := context.Background()

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// auxiliary time var
	now := time.Now()
	later := now.AddDate(0, 0, -1)

	// create
	err = repo.Create(ctx, entity.Currency{
		Base:        "BOB",
		Destination: "USD",
		Rate:        6.98,
		CreatedAt:   now,
		UpdatedAt:   &later,
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)
	// create
	err = repo.Create(ctx, entity.Currency{
		Base:        "USD",
		Destination: "BOBX",
		Rate:        6.98,
		CreatedAt:   now,
		UpdatedAt:   &now,
	})
	assert.NotNil(t, err)
	// get
	cur, err := repo.Get(ctx, "BOB", "USD")
	assert.Nil(t, err)
	assert.Equal(t, "BOB", cur.Base)
	_, err = repo.Get(ctx, "USD", "USD")
	assert.NotNil(t, err)

	 err = repo.Update(ctx, entity.Currency{
		ID:          cur.ID,
		Base:        cur.Base,
		Destination: cur.Destination,
		Rate:        7.01,
		UpdatedAt:   &now,
	})
	 assert.Nil(t, err)

}
