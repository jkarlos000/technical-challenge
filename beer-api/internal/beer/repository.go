package beer

import (
	"context"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/entity"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/dbcontext"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
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
	Update(ctx context.Context, beer entity.Beer) error
	// Delete removes the beer with given ID from the storage.
	Delete(ctx context.Context, id int) error
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new beer repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
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
				return 0, errors.BadRequest("The request is already saved")
			default:
				return 0, err
			}
		}
		return 0, err
	}
	return beer.ID, nil
}

func (r repository) GetPrice(ctx context.Context, id int, srcCurrency, dstCurrency string) (float32, error) {
	panic("Microservice consume here!")
}

func (r repository) Update(ctx context.Context, beer entity.Beer) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

