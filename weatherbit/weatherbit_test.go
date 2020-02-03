package weatherbit

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
		errorResponse := CurrentWeatherResponse{
			Error: "API key not valid, or not yet activated.",
		}

		encoder := json.NewEncoder(w)

		values := r.URL.Query()
		if values.Get("key") == "" {
			t.Error("API Key not provided")
		}

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
