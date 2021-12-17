package beer

import (
	"context"
	"github.com/biter777/countries"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/entity"
	"github.com/jkarlos000/technical-challenge/beer-api/internal/errors"
	"github.com/jkarlos000/technical-challenge/beer-api/pkg/log"
	"math"
	"strings"
	"time"
)

// Service encapsulates use case logic for beers.
type Service interface {
	Get(ctx context.Context, id int) (BeerItem, error)
	GetPrice(ctx context.Context, id int, currency string, quantity uint32) (BeerBox, error)
	Create(ctx context.Context, input CreateBeerRequest) (BeerItem, error)
	Query(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]BeerItem, error)
	Count(ctx context.Context) (int, error)
}

const (
	MaxCashAmount = 999999999.9
	MinCashAmount = 0.1
)

// BeerItem represents the data about a beer.
type BeerItem struct {
	entity.Beer
}

// CreateBeerRequest represents a beer creation request.
type CreateBeerRequest struct {
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}

// Validate validates the CreateBeerRequest fields.
func (m CreateBeerRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&m.Brewery, validation.Required, validation.Length(1, 50)),
		validation.Field(&m.Country, validation.Required, validation.Length(1, 70)),
		validation.Field(&m.Price, validation.Required, validation.Min(MinCashAmount), validation.Max(MaxCashAmount)),
		validation.Field(&m.Currency, validation.Required, is.CurrencyCode),
	)
}

// BeerBox represents the data about a price of beer box.
type BeerBox struct {
	PriceTotal float64 `json:"price_total"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new user service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Get(ctx context.Context, id int) (BeerItem, error) {
	album, err := s.repo.Get(ctx, id)
	if err != nil {
		return BeerItem{}, err
	}
	return BeerItem{album}, nil
}

func (s service) GetPrice(ctx context.Context, id int, currency string, quantity uint32) (BeerBox, error) {
	var beer BeerItem
	var err error
	if err = validation.Validate(currency, is.CurrencyCode); err != nil {
		return BeerBox{}, errors.BadRequest("The currency specified is not a valid Currency Code")
	}
	if quantity == 0 || quantity < 0 {
		quantity = 6
	}
	if beer, err = s.Get(ctx, id); err != nil {
		return BeerBox{}, err
	}
	rate, err := s.repo.GetPrice(ctx, id, beer.Currency, currency)
	if err != nil {
		return BeerBox{}, err
	}
	xrate := math.Floor(float64(rate)*100)/100
	return BeerBox{PriceTotal: float64(quantity) * xrate}, nil
}

func (s service) Create(ctx context.Context, req CreateBeerRequest) (BeerItem, error) {
	req.Country = strings.TrimSpace(req.Country)
	req.Name = strings.TrimSpace(req.Name)
	req.Currency = strings.TrimSpace(req.Currency)
	req.Brewery = strings.TrimSpace(req.Brewery)
	if err := req.Validate(); err != nil {
		return BeerItem{}, err
	}
	country := countries.ByName(req.Country)
	if country == 0 {
		return BeerItem{}, errors.BadRequest("Field country is a not valid name of country")
	}
	req.Country = country.Info().Name
	now := time.Now()
	id, err := s.repo.Create(ctx, entity.Beer{
		Name:      req.Name,
		Brewery:   req.Brewery,
		Country:   req.Country,
		Price:     req.Price,
		Currency:  req.Currency,
		CreatedAt: now,
		UpdatedAt: &now,
	})
	if err != nil {
		return BeerItem{}, err
	}
	return s.Get(ctx, id)
}

func (s service) Query(ctx context.Context, offset, limit int, filters map[string]interface{}) ([]BeerItem, error) {
	items, err := s.repo.Query(ctx, offset, limit, filters)
	if err != nil {
		return nil, err
	}
	var result []BeerItem
	for _, item := range items {
		result = append(result, BeerItem{item})
	}
	return result, nil
}

func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}
