package currency

import (
	"context"
	"errors"
	"github.com/jkarlos000/technical-challenge/currency/internal/entity"
	"reflect"
	"time"
)

func (r repository) registerNewRates(ctx context.Context, rates Currencylayer) {
	if count, err := r.Count(ctx); count == 0 {
		if err != nil {
			return
		}
		v := reflect.Indirect(reflect.ValueOf(rates.Quotes))

		for i := 0; i < v.NumField(); i++ {
			r.addNewCurrency(ctx, v.Type().Field(i).Name, v.Field(i).Interface().(float32))
		}
	} else {
		if err != nil {
			return
		}
		v := reflect.Indirect(reflect.ValueOf(rates.Quotes))

		for i := 0; i < v.NumField(); i++ {
			r.updateNewCurrency(ctx, v.Type().Field(i).Name, v.Field(i).Interface().(float32))
		}
	}
}

func (r repository) addNewCurrency(ctx context.Context, currency string, rate float32) {
	if len(currency) != 6 {
		return
	}
	now := time.Now()
	currencies := entity.Currency{
		Base:        "USD",
		Destination: currency[3:],
		Rate:        rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}
	if err := r.Create(ctx, currencies); err != nil {
		return
	}
	currencies = entity.Currency{
		Base:        currency[3:],
		Destination: "USD",
		Rate:        1 / rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}
	if err := r.Create(ctx, currencies); err != nil {
		return
	}
}

func (r repository) updateNewCurrency(ctx context.Context, currency string, rate float32) {
	var currencyA entity.Currency
	var currencyB entity.Currency
	var err error
	if currencyA, err = r.Get(ctx, "USD", currency[3:]); err != nil {
		return
	}
	if currencyB, err = r.Get(ctx, currency[3:], "USD"); err != nil {
		return
	}
	now := time.Now()
	currencies := entity.Currency{
		ID:          currencyA.ID,
		Base:        "USD",
		Destination: currency[3:],
		Rate:        rate,
		UpdatedAt:   &now,
	}
	if err := r.Update(ctx, currencies); err != nil {
		return
	}
	currencies = entity.Currency{
		ID:          currencyB.ID,
		Base:        currency[3:],
		Destination: "USD",
		Rate:        1 / rate,
		UpdatedAt:   &now,
	}
	if err := r.Update(ctx, currencies); err != nil {
		return
	}
}

func (r repository) addCurrency(ctx context.Context, base, destination string) error {
	var currencyA entity.Currency
	var currencyB entity.Currency
	var err error
	if base == "USD" && destination == "USD" {
		return errors.New("no data for this currency")
	}
	if currencyA, err = r.Get(ctx, "USD", base); err != nil {
		return err
	}
	if currencyB, err = r.Get(ctx, "USD", destination); err != nil {
		return err
	}
	now := time.Now()
	if err := r.Create(ctx, entity.Currency{
		Base:        base,
		Destination: destination,
		Rate:        currencyB.Rate / currencyA.Rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}); err != nil {
		return err
	}
	if err := r.Create(ctx, entity.Currency{
		Base:        destination,
		Destination: base,
		Rate:        currencyA.Rate / currencyB.Rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}); err != nil {
		return err
	}
	return nil
}

func (r repository) updateCurrency(ctx context.Context, base, destination string) error {
	var currencyA entity.Currency
	var currencyB entity.Currency
	var err error
	if currencyA, err = r.Get(ctx, "USD", base); err != nil {
		return err
	}
	if currencyB, err = r.Get(ctx, "USD", destination); err != nil {
		return err
	}
	now := time.Now()
	if err := r.Update(ctx, entity.Currency{
		ID:          currencyA.ID,
		Base:        base,
		Destination: destination,
		Rate:        currencyB.Rate / currencyA.Rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}); err != nil {
		return err
	}
	if err := r.Update(ctx, entity.Currency{
		ID:          currencyB.ID,
		Base:        destination,
		Destination: base,
		Rate:        currencyA.Rate / currencyB.Rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}); err != nil {
		return err
	}
	return nil
}
