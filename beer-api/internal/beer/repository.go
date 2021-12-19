package beer

import (
	"context"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/entity"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/dbcontext"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	protos "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	"github.com/lib/pq"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Get returns the beer with the specified user ID.
	Get(ctx context.Context, id int) (entity.Beer, error)
	// Count returns the number of beers.
	Count(ctx context.Context) (int, error)
	// Query returns the list of beers with the given offset and limit.
	Query(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]entity.Beer, error)
	// Create saves a new beer in the storage.
	Create(ctx context.Context, beer entity.Beer) (int, error)
	// GetPrice returns the price of beer box relative to currency
	GetPrice(ctx context.Context, id int, srcCurrency, dstCurrency string) (float32, error)
	// Update updates the beer with given ID in the storage.
	// Update(ctx context.Context, beer entity.Beer) error
	// Delete removes the beer with given ID from the storage.
	// Delete(ctx context.Context, id int) error
}

// repository persists users in database
type repository struct {
	cc		protos.CurrencyClient
	db     *dbcontext.DB
	logger log.Logger

}

// NewRepository creates a new beer repository
func NewRepository(cc protos.CurrencyClient,db *dbcontext.DB, logger log.Logger) Repository {
	return repository{cc,db, logger}
}

func (r repository) Get(ctx context.Context, id int) (entity.Beer, error) {
	var beer entity.Beer
	err := r.db.With(ctx).Select().Model(id, &beer)
	return beer, err
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("beers").Row(&count)
	return count, err
}

func (r repository) Query(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]entity.Beer, error) {
	var beers []entity.Beer
	err := r.db.With(ctx).
		Select().
		From("beers").
		OrderBy("beer_id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&beers)
	return beers, err
}

func (r repository) Create(ctx context.Context, beer entity.Beer) (int, error) {
	if err := r.db.With(ctx).Model(&beer).Insert(); err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return 0, errors.BadRequest("The request is already saved in database")
			default:
				return 0, err
			}
		}
		return 0, err
	}
	return beer.ID, nil
}

func (r repository) GetPrice(ctx context.Context, id int, srcCurrency, dstCurrency string) (float32, error) {
	var err error
	if _, err = r.Get(ctx, id); err != nil {
		return 0.0, err
	}
	rr := &protos.Request{
		Base:        srcCurrency,
		Destination: dstCurrency,
	}
	var rate *protos.Response
	if rate, err = r.cc.GetPrice(ctx, rr); err != nil {
		r.logger.Error(err)
		return 0.0, errors.InternalServerError("")
	}
	return rate.Rate, nil
}
