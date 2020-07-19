package currencyexchange

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

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
		err := encoder.Encode(historicalExchangeRate)
		if err != nil {
			t.Error(err.Error())
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}

	c := Client{
		baseURL:    rawURL,
		httpClient: testClient,
	}

	startAt, _ := time.Parse(dateLayout, "2020-02-18")
	endAt, _ := time.Parse(dateLayout, "2020-02-20")

	params := HistoryArgs{
		Base:    "USD",
		Symbols: []string{"MXN"},
		StartAt: startAt,
		EndAt:   endAt,
	}

	resp, err := c.HistoricalCurrencyRate(params)

	assert.Nil(t, err)
	assert.Equal(t, 18.638683432, resp.Rates["2020-02-18"]["MXN"])
	assert.Equal(t, 18.5824074074, resp.Rates["2020-02-19"]["MXN"])
}
