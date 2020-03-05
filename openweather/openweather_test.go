package openweather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailApiKey(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorResponse := Weather{
			Cod:     401,
			Message: "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info.",
		}

		values := r.URL.Query()
		if values.Get("APPID") == "" {
			t.Error("API Key not provided")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse)
	}))

	defer sv.Close()

	c := Client{
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}

	_, err := c.GetCurrentWeather(queryParams)

	assert.NotNil(t, err)
}
