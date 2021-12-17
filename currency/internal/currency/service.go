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
	ctx    context.Context
}

var ApiKey string

const Url = "http://api.currencylayer.com/live?access_key="

// NewCurrency creates a new Currency server
func NewCurrency(repo Repository, logger log.Logger, ctx context.Context) proto.CurrencyServer {
	s := &service{repo: repo, logger: logger, ctx: ctx}
	go s.handleRatesUpdate()
	return s
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
	for {
		select {
		case <-server.Context().Done():
			return nil
		default:
			res, err := s.GetPrice(server.Context(), request)
			switch err.(type) {
			case nil:
			default:
				return err
			}
			if err = server.Send(res); err != nil {
				return err
			}
		}
	}
}

func (s service) handleRatesUpdate() {
	update := s.repo.MonitorRates(s.ctx, 30*time.Minute)
	for range update {
	}
}
