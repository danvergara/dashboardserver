package weatherbit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeatherForecastWithoutCity(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorCurrentResponse := ForecastWeatherResponse{
			Error: "Invalid Parameters supplied.",
		}

		if r.URL.Path != "/v2.0/forecast/daily" {
			t.Error("Bad forecast weather path")
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
	_, err := c.GetWeatherForecast(queryParams)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Parameters supplied.", err.Error())
}

func TestGetWeatherForecastPassingACity(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		forecastResponse := ForecastWeatherResponse{
			CityName: "Mexico",
			Data: []ForecastData{
				{
					ValidDate:  "2020-01-29",
					WinSpd:     3.46898,
					HighTemp:   20.7,
					LowTemp:    11,
					MaxTemp:    21,
					MinTemp:    10.4,
					AppMaxTemp: 19.1,
					AppMinTemp: 10.5,
					Precip:     0,
					Clouds:     0,
					Pop:        0,
				},
				{
					ValidDate:  "2020-01-30",
					WinSpd:     3.61541,
					HighTemp:   19.6,
					LowTemp:    10.7,
					MaxTemp:    19.7,
					MinTemp:    10.6,
					AppMaxTemp: 18.3,
					AppMinTemp: 10.7,
					Precip:     0,
					Clouds:     0,
					Pop:        0,
				},
			},
		}

		if r.URL.Path != "/v2.0/forecast/daily" {
			t.Error("Bad forecast weather path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(forecastResponse)

	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("city", "Mexico")
	queryParams.Add("days", "2")

	resp, err := c.GetWeatherForecast(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(resp.Data))
	assert.Equal(t, "Mexico", resp.CityName)

}
