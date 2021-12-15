package beer

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jkarlos000/technical-challenge/internal/entity"
	"github.com/jkarlos000/technical-challenge/pkg/log"
)

// Service encapsulates use case logic for beers.
type Service interface {
	Get(ctx context.Context, id string) (Beer, error)
	GetPrice(ctx context.Context, id, currency string, quantity uint32) (Beer, error)
	Create(ctx context.Context, input CreateBeerRequest) (Beer, error)
	Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]Beer, error)
	Count(ctx context.Context, term string, filters map[string]interface{}) (int, error)
}

const (
	MaxCashAmount = 999999999.9
	MinCashAmount = 0.1
)

// Beer represents the data about an beer.
type Beer struct {
	entity.Beer
}

// CreateBeerRequest represents an beer creation request.
type CreateBeerRequest struct {
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Country   string    `json:"country"`
	Price     float32   `json:"price"`
	Currency  string    `json:"currency"`
}

// Validate validates the CreateBeerRequest fields.
func (m CreateBeerRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&m.Brewery, validation.Required, validation.Length(1, 50)),
		validation.Field(&m.Country, validation.Required, validation.Length(1, 70)),
		validation.Field(&m.Price, validation.Required, validation.Min(MinCashAmount), validation.Max(MaxCashAmount)),
		validation.Field(&m.Currency, validation.Required, validation.Length(1, 3)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new user service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Get(ctx context.Context, id string) (Beer, error) {
	panic("implement me")
}

func (s service) GetPrice(ctx context.Context, id, currency string, quantity uint32) (Beer, error) {
	panic("implement me")
}

func (s service) Create(ctx context.Context, input CreateBeerRequest) (Beer, error) {
	panic("implement me")
}

func (s service) Query(ctx context.Context, offset, limit int, term string, filters map[string]interface{}) ([]Beer, error) {
	panic("implement me")
}

func (s service) Count(ctx context.Context, term string, filters map[string]interface{}) (int, error) {
	panic("implement me")
}

