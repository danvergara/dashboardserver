package currencyexchange

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPassingWrongBase(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorResponse := []byte(`
			{
				"error": "Base 'USsD' is not supported."
			}
		`)

		w.WriteHeader(http.StatusBadRequest)

		_, err := w.Write(errorResponse)

		if err != nil {
			t.Error(err.Error())
		}

		if r.URL.Path != "/latest" {
			t.Error("Bad latest paht")
		}

	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}

	c := Client{
		baseURL:    rawURL,
		httpClient: testClient,
	}

	// Wrong base value
	params := LatestArgs{
		Base: "USsD",
	}

	_, err := c.LatestCurrencyExchange(params)

	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "Base 'USsD' is not supported."))
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
		err := encoder.Encode(currencyExchangeData)
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

	params := LatestArgs{
		Base:    "USD",
		Symbols: []string{"MXN"},
	}

	resp, err := c.LatestCurrencyExchange(params)

	assert.Nil(t, err)
	assert.Equal(t, 18.9966669753, resp.Rates["MXN"])
}
