package weatherbit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentWeatherWithoutCity(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorCurrentResponse := CurrentWeatherResponse{
			Error: "Invalid Parameters supplied.",
		}

		if r.URL.Path != "/v2.0/current" {
			t.Error("Bad current weather path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(errorCurrentResponse)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	_, err := c.GetCurrentWeather(queryParams)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Parameters supplied.", err.Error())
}

func TestGetCurrentWeatherPassingACity(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentResponse := CurrentWeatherResponse{
			Data: []CurrentWeatherData{
				{

					ObTime:   "2020-01-29 05:53",
					Clouds:   88,
					SolarRad: 0,
					CityName: "Mexico City",
					WinSpd:   1.08563,
					Temp:     14,
					DateTime: "2020-01-29:05",
				},
			},
		}

		if r.URL.Path != "/v2.0/current" {
			t.Error("Bad current weather path")
		}

		enconder := json.NewEncoder(w)
		enconder.Encode(currentResponse)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("city", "Mexico")

	resp, err := c.GetCurrentWeather(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 14.0, resp.Data[0].Temp)
	assert.Equal(t, 88, resp.Data[0].Clouds)
}
