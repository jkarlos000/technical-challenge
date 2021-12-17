package currency

import (
	"context"
	"encoding/json"
	"github.com/jkarlos000/technical-challenge/currency/internal/entity"
	"github.com/jkarlos000/technical-challenge/currency/pkg/dbcontext"
	"github.com/jkarlos000/technical-challenge/currency/pkg/log"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Get returns the currency with the specified base and destination currency.
	Get(ctx context.Context, base, destination string) (entity.Currency, error)
	// Count returns the number of currencies relation.
	Count(ctx context.Context) (int, error)
	// Create saves a new relation of currencies in the storage.
	Create(ctx context.Context, currency entity.Currency) error
	// Update updates the currencies relation with given base and destination in the storage.
	Update(ctx context.Context, currency entity.Currency) error
	// Delete removes the currencies relation with given base and destination from the storage.
	Delete(ctx context.Context, id int) error
	// MonitorRates update every 30 minutes the actual rates values.
	MonitorRates(ctx context.Context, interval time.Duration) chan struct{}
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

type Currencylayer struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		USDAED float32 `json:"USDAED"`
		USDAFN float32 `json:"USDAFN"`
		USDALL float32 `json:"USDALL"`
		USDAMD float32 `json:"USDAMD"`
		USDANG float32 `json:"USDANG"`
		USDAOA float32 `json:"USDAOA"`
		USDARS float32 `json:"USDARS"`
		USDAUD float32 `json:"USDAUD"`
		USDAWG float32 `json:"USDAWG"`
		USDAZN float32 `json:"USDAZN"`
		USDBAM float32 `json:"USDBAM"`
		USDBBD float32 `json:"USDBBD"`
		USDBDT float32 `json:"USDBDT"`
		USDBGN float32 `json:"USDBGN"`
		USDBHD float32 `json:"USDBHD"`
		USDBIF float32 `json:"USDBIF"`
		USDBMD float32 `json:"USDBMD"`
		USDBND float32 `json:"USDBND"`
		USDBOB float32 `json:"USDBOB"`
		USDBRL float32 `json:"USDBRL"`
		USDBSD float32 `json:"USDBSD"`
		USDBTC float32 `json:"USDBTC"`
		USDBTN float32 `json:"USDBTN"`
		USDBWP float32 `json:"USDBWP"`
		USDBYN float32 `json:"USDBYN"`
		USDBYR float32 `json:"USDBYR"`
		USDBZD float32 `json:"USDBZD"`
		USDCAD float32 `json:"USDCAD"`
		USDCDF float32 `json:"USDCDF"`
		USDCHF float32 `json:"USDCHF"`
		USDCLF float32 `json:"USDCLF"`
		USDCLP float32 `json:"USDCLP"`
		USDCNY float32 `json:"USDCNY"`
		USDCOP float32 `json:"USDCOP"`
		USDCRC float32 `json:"USDCRC"`
		USDCUC float32 `json:"USDCUC"`
		USDCUP float32 `json:"USDCUP"`
		USDCVE float32 `json:"USDCVE"`
		USDCZK float32 `json:"USDCZK"`
		USDDJF float32 `json:"USDDJF"`
		USDDKK float32 `json:"USDDKK"`
		USDDOP float32 `json:"USDDOP"`
		USDDZD float32 `json:"USDDZD"`
		USDEGP float32 `json:"USDEGP"`
		USDERN float32 `json:"USDERN"`
		USDETB float32 `json:"USDETB"`
		USDEUR float32 `json:"USDEUR"`
		USDFJD float32 `json:"USDFJD"`
		USDFKP float32 `json:"USDFKP"`
		USDGBP float32 `json:"USDGBP"`
		USDGEL float32 `json:"USDGEL"`
		USDGGP float32 `json:"USDGGP"`
		USDGHS float32 `json:"USDGHS"`
		USDGIP float32 `json:"USDGIP"`
		USDGMD float32 `json:"USDGMD"`
		USDGNF float32 `json:"USDGNF"`
		USDGTQ float32 `json:"USDGTQ"`
		USDGYD float32 `json:"USDGYD"`
		USDHKD float32 `json:"USDHKD"`
		USDHNL float32 `json:"USDHNL"`
		USDHRK float32 `json:"USDHRK"`
		USDHTG float32 `json:"USDHTG"`
		USDHUF float32 `json:"USDHUF"`
		USDIDR float32 `json:"USDIDR"`
		USDILS float32 `json:"USDILS"`
		USDIMP float32 `json:"USDIMP"`
		USDINR float32 `json:"USDINR"`
		USDIQD float32 `json:"USDIQD"`
		USDIRR float32 `json:"USDIRR"`
		USDISK float32 `json:"USDISK"`
		USDJEP float32 `json:"USDJEP"`
		USDJMD float32 `json:"USDJMD"`
		USDJOD float32 `json:"USDJOD"`
		USDJPY float32 `json:"USDJPY"`
		USDKES float32 `json:"USDKES"`
		USDKGS float32 `json:"USDKGS"`
		USDKHR float32 `json:"USDKHR"`
		USDKMF float32 `json:"USDKMF"`
		USDKPW float32 `json:"USDKPW"`
		USDKRW float32 `json:"USDKRW"`
		USDKWD float32 `json:"USDKWD"`
		USDKYD float32 `json:"USDKYD"`
		USDKZT float32 `json:"USDKZT"`
		USDLAK float32 `json:"USDLAK"`
		USDLBP float32 `json:"USDLBP"`
		USDLKR float32 `json:"USDLKR"`
		USDLRD float32 `json:"USDLRD"`
		USDLSL float32 `json:"USDLSL"`
		USDLTL float32 `json:"USDLTL"`
		USDLVL float32 `json:"USDLVL"`
		USDLYD float32 `json:"USDLYD"`
		USDMAD float32 `json:"USDMAD"`
		USDMDL float32 `json:"USDMDL"`
		USDMGA float32 `json:"USDMGA"`
		USDMKD float32 `json:"USDMKD"`
		USDMMK float32 `json:"USDMMK"`
		USDMNT float32 `json:"USDMNT"`
		USDMOP float32 `json:"USDMOP"`
		USDMRO float32 `json:"USDMRO"`
		USDMUR float32 `json:"USDMUR"`
		USDMVR float32 `json:"USDMVR"`
		USDMWK float32 `json:"USDMWK"`
		USDMXN float32 `json:"USDMXN"`
		USDMYR float32 `json:"USDMYR"`
		USDMZN float32 `json:"USDMZN"`
		USDNAD float32 `json:"USDNAD"`
		USDNGN float32 `json:"USDNGN"`
		USDNIO float32 `json:"USDNIO"`
		USDNOK float32 `json:"USDNOK"`
		USDNPR float32 `json:"USDNPR"`
		USDNZD float32 `json:"USDNZD"`
		USDOMR float32 `json:"USDOMR"`
		USDPAB float32 `json:"USDPAB"`
		USDPEN float32 `json:"USDPEN"`
		USDPGK float32 `json:"USDPGK"`
		USDPHP float32 `json:"USDPHP"`
		USDPKR float32 `json:"USDPKR"`
		USDPLN float32 `json:"USDPLN"`
		USDPYG float32 `json:"USDPYG"`
		USDQAR float32 `json:"USDQAR"`
		USDRON float32 `json:"USDRON"`
		USDRSD float32 `json:"USDRSD"`
		USDRUB float32 `json:"USDRUB"`
		USDRWF float32 `json:"USDRWF"`
		USDSAR float32 `json:"USDSAR"`
		USDSBD float32 `json:"USDSBD"`
		USDSCR float32 `json:"USDSCR"`
		USDSDG float32 `json:"USDSDG"`
		USDSEK float32 `json:"USDSEK"`
		USDSGD float32 `json:"USDSGD"`
		USDSHP float32 `json:"USDSHP"`
		USDSLL float32 `json:"USDSLL"`
		USDSOS float32 `json:"USDSOS"`
		USDSRD float32 `json:"USDSRD"`
		USDSTD float32 `json:"USDSTD"`
		USDSVC float32 `json:"USDSVC"`
		USDSYP float32 `json:"USDSYP"`
		USDSZL float32 `json:"USDSZL"`
		USDTHB float32 `json:"USDTHB"`
		USDTJS float32 `json:"USDTJS"`
		USDTMT float32 `json:"USDTMT"`
		USDTND float32 `json:"USDTND"`
		USDTOP float32 `json:"USDTOP"`
		USDTRY float32 `json:"USDTRY"`
		USDTTD float32 `json:"USDTTD"`
		USDTWD float32 `json:"USDTWD"`
		USDTZS float32 `json:"USDTZS"`
		USDUAH float32 `json:"USDUAH"`
		USDUGX float32 `json:"USDUGX"`
		USDUSD float32 `json:"USDUSD"`
		USDUYU float32 `json:"USDUYU"`
		USDUZS float32 `json:"USDUZS"`
		USDVEF float32 `json:"USDVEF"`
		USDVND float32 `json:"USDVND"`
		USDVUV float32 `json:"USDVUV"`
		USDWST float32 `json:"USDWST"`
		USDXAF float32 `json:"USDXAF"`
		USDXAG float32 `json:"USDXAG"`
		USDXAU float32 `json:"USDXAU"`
		USDXCD float32 `json:"USDXCD"`
		USDXDR float32 `json:"USDXDR"`
		USDXOF float32 `json:"USDXOF"`
		USDXPF float32 `json:"USDXPF"`
		USDYER float32 `json:"USDYER"`
		USDZAR float32 `json:"USDZAR"`
		USDZMK float32 `json:"USDZMK"`
		USDZMW float32 `json:"USDZMW"`
		USDZWL float32 `json:"USDZWL"`
	} `json:"quotes"`
}

// NewRepository creates a new currency repository
func NewRepository(db *dbcontext.DB, logger log.Logger, apikeyDSN string) Repository {
	ApiKey = apikeyDSN
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, base, destination string) (entity.Currency, error) {
	panic("implement me")
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("currencies").Row(&count)
	return count, err
}

func (r repository) Create(ctx context.Context, currency entity.Currency) error {
	if err := r.db.With(ctx).Model(&currency).Insert(); err != nil {
		return err
	}
	return nil
}

func (r repository) Update(ctx context.Context, currency entity.Currency) error {
	return r.db.With(ctx).Model(&currency).Exclude("CreatedAt").Update()
}

func (r repository) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

// MonitorRates checks the rates in the CurrencyLayer, actual limit of 250 request by day.
func (r repository) MonitorRates(ctx context.Context, interval time.Duration) chan struct{} {
	ret := make(chan struct{})
	url := Url + ApiKey

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				var netClient = &http.Client{
					Timeout: time.Second * 10,
				}
				response, err := netClient.Get(url)
				if err != nil {
					ret <- struct{}{}
				}
				var currencyRates Currencylayer
				buf, _ := ioutil.ReadAll(response.Body)
				json.Unmarshal(buf, &currencyRates)
				// Use json.Decode for reading streams of JSON data
				if err := json.NewDecoder(response.Body).Decode(&currencyRates); err != nil {
					r.logger.Error(err)
					response.Body.Close()
					ret <- struct{}{}
				}
				r.logger.Info("Updating rates values from currencylayer")
				ret <- struct{}{}
				response.Body.Close()
			}
		}
	}()
	return ret
}

func (r repository) registerRates(ctx context.Context, rates Currencylayer) {
	if count, err := r.Count(ctx); count == 0 {
		if err != nil {
			return
		}
		v := reflect.Indirect(reflect.ValueOf(rates.Quotes))

		for i := 0; i < v.NumField(); i++ {
			r.addCurrency(ctx, v.Type().Field(i).Name, v.Field(i).Interface().(float32))
		}
	} else {
		if err != nil {
			return
		}
		v := reflect.Indirect(reflect.ValueOf(rates.Quotes))

		for i := 0; i < v.NumField(); i++ {
			r.updateCurrency(ctx, v.Type().Field(i).Name, v.Field(i).Interface().(float32))
		}
	}
}

func (r repository) addCurrency(ctx context.Context, currency string, rate float32) {
	if len(currency) != 6 {
		return
	}
	now := time.Now()
	currencies := entity.Currency{
		Base:        "USD",
		Destination: currency[3:],
		Rate:        rate,
		UpdatedAt:   &now,
	}
	if err := r.Update(ctx, currencies); err != nil {
		return
	}
	currencies = entity.Currency{
		Base:        currency[3:],
		Destination: "USD",
		Rate:        1/rate,
		UpdatedAt:   &now,
	}
	if err := r.Update(ctx, currencies); err != nil {
		return
	}
}

func (r repository) updateCurrency(ctx context.Context, currency string, rate float32) {
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
		Rate:        1/rate,
		CreatedAt:   now,
		UpdatedAt:   &now,
	}
	if err := r.Create(ctx, currencies); err != nil {
		return
	}
}


