package economics

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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
		currencyClient := currencyexchange.Client{}
		params := url.Values{}
		params.Add("base", "USD")
		params.Add("symbols", "MXN")

		resp, err := currencyClient.GetLatestCurrencyExchange(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		currencyExchangeResponse := CurrencyExchangeResponse{CurrencyExchange: resp}
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
		currencyClient := currencyexchange.Client{}
		now := time.Now()
		dateLayout := "2006-01-02"

		params := url.Values{}
		params.Add("base", "USD")
		params.Add("symbols", "MXN")
		params.Add("end_at", now.Format(dateLayout))
		params.Add("start_at", now.AddDate(0, 0, -20).Format(dateLayout))
		resp, err := currencyClient.GetHistoricalCurrencyRate(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		historicalCurrencyResponse := HistoricalCurrencyResponse{HistoricalCurrencyRates: resp}
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(historicalCurrencyResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
