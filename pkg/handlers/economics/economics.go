package economics

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/currencyexchange"
)

// CurrencyExchangeResponse {"currency-exchange": currencyexchange.ExchangeRateData}
type CurrencyExchangeResponse struct {
	CurrencyExchange currencyexchange.ExchangeRateData `json:"currency-exchange"`
}

// HistoricalCurrencyResponse struct {"historical-currency-rates": currencyexchange.HistoricalExchangeRateData}
type HistoricalCurrencyResponse struct {
	HistoricalCurrencyRates currencyexchange.HistoricalExchangeRateData `json:"historical-currency-rates"`
}

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
