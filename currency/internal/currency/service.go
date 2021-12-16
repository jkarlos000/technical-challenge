package currency

// BeerBox is a beer box price of business logic.
type BeerBox struct {
	Price	float32
}

// Service defines the interface exposed by this package.
type Service interface {
	GetPrice(id int64, base, destination float32) (BeerBox, error)
}

func (b *BeerBox) GetPrice(id int64, base, destination float32) (BeerBox, error) {
	panic("implement me")
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{r, l, make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest)}
	go c.handleUpdates()

	return c
}

