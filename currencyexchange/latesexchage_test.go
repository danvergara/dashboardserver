package currencyexchange

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassingWrongBase(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorCurrencyData := ExchangeRateData{
			Error: "Base 'USsD' is not supported.",
		}

		if r.URL.Path != "/latest" {
			t.Error("Bad latest paht")
		}

		econder := json.NewEncoder(w)
		econder.Encode(errorCurrencyData)
	}))

	defer sv.Close()
	c := Client{
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("base", "WRONGBASE")
	_, err := c.GetLatestCurrencyExchange(queryParams)

	assert.NotNil(t, err)
	assert.Equal(t, "Base 'USsD' is not supported.", err.Error())
}

func TestGetLatestCurrencyExchange(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currencyExchangeData := ExchangeRateData{
			Date:  "2020-02-21",
			Base:  "USD",
			Rates: map[string]float64{"MXN": 18.9966669753},
		}

		if r.URL.Path != "/latest" {
			t.Error("Bad latest path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(currencyExchangeData)
	}))

	defer sv.Close()

	c := Client{
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("base", "USD")
	queryParams.Add("symbols", "MXN")

	resp, err := c.GetLatestCurrencyExchange(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 18.9966669753, resp.Rates["MXN"])
}
