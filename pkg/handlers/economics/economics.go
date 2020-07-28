package economics

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/currencyexchange"
)

// The current currency exchange between MXN-USD
// swagger:response currencyExchangeResponse
type currencyExchangeResponseWrapper struct {
	// in: body
	Body CurrencyExchangeResponse
}

// CurrencyExchangeResponse {"currency-exchange": currencyexchange.ExchangeRateData}
type CurrencyExchangeResponse struct {
	CurrencyExchange currencyexchange.ExchangeRateData `json:"currency-exchange"`
}

//The Historical Currency Response between MXN-USD
// swagger:response historicalCurrencyResponse
type historicalCurrencyResponseWrapper struct {
	// in: body
	Body HistoricalCurrencyResponse
}

// HistoricalCurrencyResponse struct {"historical-currency-rates": currencyexchange.HistoricalExchangeRateData}
type HistoricalCurrencyResponse struct {
	HistoricalCurrencyRates currencyexchange.HistoricalExchangeRateData `json:"historical-currency-rates"`
}

// swagger:route GET /v1/currency-exchange currency-exchange
//
// Returns a the current currency exchange rate between MXN-USD
//
// Produces:
// - application/json
//
// Responses:
//	200: currencyExchangeResponse

// CurrencyExchange returns the currency exchange between the dollar and the mexican peso
func CurrencyExchange(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := currencyexchange.NewClient()

		params := currencyexchange.LatestArgs{
			Base:    "USD",
			Symbols: []string{"MXN"},
		}

		res, err := c.LatestCurrencyExchange(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		currencyExchangeResponse := CurrencyExchangeResponse{CurrencyExchange: *res}
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(currencyExchangeResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}

// swagger:route GET /v1/historical-currency-rates historical-currency-rates
//
// Returns a the historical currency exchange rates between MXN-USD of the previous 20 days
//
// Produces:
// - application/json
//
// Responses:
//	200: historicalCurrencyResponse

// HistoricalCurrencyRates returns the historical currency rates given start and end dates
func HistoricalCurrencyRates(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := currencyexchange.NewClient()
		now := time.Now()

		params := currencyexchange.HistoryArgs{
			Base:    "USD",
			Symbols: []string{"MXN"},
			EndAt:   now,
			StartAt: now.AddDate(0, 0, -20),
		}

		res, err := c.HistoricalCurrencyRate(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		historicalCurrencyResponse := HistoricalCurrencyResponse{HistoricalCurrencyRates: *res}

		encoder := json.NewEncoder(w)
		if err = encoder.Encode(historicalCurrencyResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
