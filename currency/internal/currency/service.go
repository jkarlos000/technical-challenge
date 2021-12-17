package currency

import (
	"context"
	proto "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	"github.com/jkarlos000/technical-challenge/currency/internal/entity"
	"github.com/jkarlos000/technical-challenge/currency/pkg/log"
	"strings"
	"time"
)

type service struct {
	proto.UnimplementedCurrencyServer
	repo   Repository
	logger log.Logger
}

var ApiKey string
const Url = "http://api.currencylayer.com/live?access_key="

// NewCurrency creates a new Currency server
func NewCurrency(repo Repository, logger log.Logger) proto.CurrencyServer {
	return service{repo: repo, logger: logger}
}

func (s service) GetPrice(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	var currency entity.Currency
	var err error
	base := strings.ToUpper(request.Base)
	destination := strings.ToUpper(request.Destination)
	if currency, err = s.repo.Get(ctx, base, destination); err != nil {
		return &proto.Response{}, err
	}
	return &proto.Response{Rate: currency.Rate}, nil
}

func (s service) GetPriceStream(request *proto.Request, server proto.Currency_GetPriceStreamServer) error {
	panic("implement me")
}

func (s service) handleRatesUpdate(ctx context.Context) {
	update := s.repo.MonitorRates(ctx, 30 * time.Minute)
	for range update {}
}
