package currency_test

import (
	"context"
	"errors"
	v1 "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	"github.com/jkarlos000/technical-challenge/currency/internal/currency"
	"github.com/jkarlos000/technical-challenge/currency/internal/currency/currencytesting"
	"github.com/jkarlos000/technical-challenge/currency/internal/entity"
	"github.com/jkarlos000/technical-challenge/currency/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	logger, _ := log.NewForTest()
	repo := &currencytesting.FakeRepository{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repo.MonitorRatesReturns(make(chan struct{}))
	s := currency.NewCurrency(repo, logger, ctx)

	var c entity.Currency
	repo.GetReturns(c, nil)
	request := v1.Request{Base: "CLP", Destination: "USD"}
	response, err := s.GetPrice(ctx, &request)
	assert.Nil(t, err)
	assert.NotNil(t, response.Rate)
	repo.GetReturns(c, errors.New("no data"))
	response, err = s.GetPrice(ctx, &request)
	assert.NotNil(t, err)

	repo.MonitorRatesReturns(nil)
	cancel()
	var r v1.Request
	var cgps currencytesting.FakeCurrency_GetPriceStreamServer
	cgps.ContextReturns(ctx)
	err = s.GetPriceStream(&r, &cgps)
	assert.Nil(t, err)

}
