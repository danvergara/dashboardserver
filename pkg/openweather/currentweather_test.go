package openweather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
			Name:       "Mexico City",
			StatusCode: 200,
		}

		if r.URL.Path != "/weather" {
			t.Error("Bad current weather path")
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(currentResponse)
		if err != nil {
			t.Error(err.Error())
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := Client{
		apiKey:     "FAKE_API_KEY",
		baseURL:    rawURL,
		httpClient: testClient,
	}

	params := WeatherArgs{
		ID:    3527646,
		Units: "metric",
	}

	resp, err := c.CurrentWeather(params)

	assert.Nil(t, err)
	assert.Equal(t, 15.17, resp.Main.Temp)
	assert.Equal(t, 37, resp.Main.Humidity)
	assert.Equal(t, 3.1, resp.Wind.Speed)
	assert.Equal(t, "01n", resp.Weather[0].Icon)
	assert.Equal(t, 9.0, resp.Main.TempMin)
	assert.Equal(t, 19.0, resp.Main.TempMax)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCurrentWeatherWithoutCityID(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorResponse := []byte(`
			{
				"cod": 400,
				"message": "Nothing to geocode"
			}
		`)

		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(errorResponse)
		if err != nil {
			t.Error(err.Error())
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := Client{
		apiKey:     "FAKE_API_KEY",
		baseURL:    rawURL,
		httpClient: testClient,
	}

	params := WeatherArgs{}

	resp, err := c.CurrentWeather(params)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "Nothing to geocode", resp.Message)
}

func TestFailApiKey(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorResponse := []byte(`
			{
				"cod": 401,
				"message": "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."
			}
		`)

		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write(errorResponse)
		if err != nil {
			t.Error(err.Error())
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := Client{
		apiKey:     "FAKE_API_KEY",
		baseURL:    rawURL,
		httpClient: testClient,
	}

	params := WeatherArgs{}
	resp, err := c.CurrentWeather(params)

	assert.Nil(t, err)
	assert.Equal(t, "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info.", resp.Message)
}
