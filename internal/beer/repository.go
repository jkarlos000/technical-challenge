package beer

import (
	"context"
	"github.com/jkarlos000/technical-challenge/internal/entity"
	"github.com/jkarlos000/technical-challenge/pkg/dbcontext"
	"github.com/jkarlos000/technical-challenge/pkg/log"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Get returns the user with the specified user ID.
	Get(ctx context.Context, id string) (entity.Beer, error)
	// Count returns the number of users.
	Count(ctx context.Context) (int, error)
	// Query returns the list of users with the given offset and limit.
	Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]entity.Beer, error)
	// Create saves a new user in the storage.
	Create(ctx context.Context, user entity.Beer, admin entity.Beer) error
	// Update updates the user with given ID in the storage.
	Update(ctx context.Context, user entity.Beer, admin entity.Beer, excludes []string) error
	// Delete removes the user with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Beer, error) {
	panic("implement me")
}

func (r repository) Count(ctx context.Context) (int, error) {
	panic("implement me")
}

func (r repository) Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]entity.Beer, error) {
	panic("implement me")
}

func (r repository) Create(ctx context.Context, user entity.Beer, admin entity.Beer) error {
	panic("implement me")
}

func (r repository) Update(ctx context.Context, user entity.Beer, admin entity.Beer, excludes []string) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

