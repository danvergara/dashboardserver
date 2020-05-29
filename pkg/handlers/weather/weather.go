package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/openweather"
)

// CurrentWeatherResponse {"current-weather": openweather.Weather}
type CurrentWeatherResponse struct {
	CurrentWeather openweather.Weather `json:"current-weather"`
}

// ForecastResponse {"weather-forecast": openweather.Forecast}
type ForecastResponse struct {
	WeatherForecast openweather.Forecast `json:"weather-forecast"`
}

// CurrentWeather Returns the main data of the current Weather
func CurrentWeather(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		weatherClient := openweather.Client{
			APIKey: os.Getenv("OPENWEATHER_KEY"),
		}

		params := url.Values{}
		params.Add("id", "3527646")
		params.Add("units", "metric")

		response, err := weatherClient.GetCurrentWeather(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		encoder := json.NewEncoder(w)
		currentWeatherResponse := CurrentWeatherResponse{CurrentWeather: response}
		if err = encoder.Encode(currentWeatherResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}

// Forecast returns the forecast of the next 5 days
func Forecast(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		weatherClient := openweather.Client{
			APIKey: os.Getenv("OPENWEATHER_KEY"),
		}

		params := url.Values{}
		params.Add("id", "3527646")
		params.Add("units", "metric")

		response, err := weatherClient.GetWeatherForecast(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		forecastWeatherResponse := ForecastResponse{WeatherForecast: response}
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(forecastWeatherResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
