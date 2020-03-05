package openweather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentWeatherWithoutCityID(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Weather{
			Message: "Nothing to geocode",
		}

		if r.URL.Path != "/data/2.5/weather" {
			t.Error("Bad weather path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(response)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	_, err := c.GetCurrentWeather(queryParams)

	assert.NotNil(t, err)
	assert.Equal(t, "Nothing to geocode", err.Error())
}

func TestGetCurrentWeatherPassingCityID(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentResponse := Weather{
			Weather: []MetaWeather{
				{
					Main:        "Clear",
					Description: "clear sky",
					Icon:        "01n",
				},
			},
			Main: Main{
				Temp:      15.17,
				FeelsLike: 11.17,
				TempMin:   9.0,
				TempMax:   19.0,
				Pressure:  1025,
				Humidity:  37,
			},
			Wind: Wind{
				Speed: 3.1,
			},
			Clouds: Clouds{
				All: 5,
			},
			Name: "Mexico City",
			Cod:  200,
		}

		if r.URL.Path != "/data/2.5/weather" {
			t.Error("Bad current weather path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(currentResponse)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	// Mexico City ID
	queryParams.Add("id", "3527646")
	queryParams.Add("unit", "metric")
	resp, err := c.GetCurrentWeather(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 15.17, resp.Main.Temp)
	assert.Equal(t, 37, resp.Main.Humidity)
	assert.Equal(t, 3.1, resp.Wind.Speed)
	assert.Equal(t, "01n", resp.Weather[0].Icon)
	assert.Equal(t, 9.0, resp.Main.TempMin)
	assert.Equal(t, 19.0, resp.Main.TempMax)
	assert.Equal(t, 200, resp.Cod)
}
