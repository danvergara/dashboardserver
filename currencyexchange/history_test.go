package currencyexchange

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHistoricalCurrencyRates(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		historicalExchangeRate := HistoricalExchangeRateData{
			Base:    "USD",
			StartAt: "2020-02-18",
			EndAt:   "2020-02-20",
			Rates: map[string]map[string]float64{
				"2020-02-18": {"MXN": 18.638683432},
				"2020-02-19": {"MXN": 18.5824074074},
				"2020-02-20": {"MXN": 18.6918443003},
			},
		}

		if r.URL.Path != "/history" {
			t.Error("Bad history path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(historicalExchangeRate)
	}))

	defer sv.Close()

	c := Client{
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("start_at", "2020-02-18")
	queryParams.Add("end_at", "2020-02-20")
	queryParams.Add("base", "USD")
	queryParams.Add("symbols", "MXN")

	resp, err := c.GetHistoricalCurrencyRate(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 18.638683432, resp.Rates["2020-02-18"]["MXN"])
	assert.Equal(t, 18.5824074074, resp.Rates["2020-02-19"]["MXN"])
}
