package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/openweather"
)

// The Current Weather data in json format
// swagger:response currentWeather
type currentWeatherWrapper struct {
	// The neede weather data
	// in: body
	Body CurrentWeatherResponse
}

// CurrentWeatherResponse {"current-weather": openweather.Weather}
type CurrentWeatherResponse struct {
	CurrentWeather openweather.Weather `json:"current-weather"`
}

// The Forecast Weather data in json format
// swagger:response weatherForecast
type weatherForecastWrapper struct {
	// The weather forecast data for the next 5 days
	// in: body
	Body ForecastResponse
}

// ForecastResponse {"weather-forecast": openweather.Forecast}
type ForecastResponse struct {
	WeatherForecast openweather.Forecast `json:"weather-forecast"`
}

// swagger:route GET /v1/current-weather current-weather
//
// Returns the current weather data
//
// Produces:
// - application/json
//
// Responses:
//	200: currentWeather

// CurrentWeather Returns the main data of the current Weather
func CurrentWeather(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := openweather.NewClient(os.Getenv("OPENWEATHER_KEY"))

		params := openweather.WeatherArgs{
			ID:    3527646,
			Units: "metric",
		}

		res, err := c.CurrentWeather(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		encoder := json.NewEncoder(w)
		currentWeatherResponse := CurrentWeatherResponse{CurrentWeather: res}
		if err = encoder.Encode(currentWeatherResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}

// swagger:route GET /v1/weather-forecast weather-forecast
//
// Returns the weather forecast data for the next 5 days
//
// Produces:
// - application/json
//
// Responses:
//	200: weatherForecast

// Forecast returns the forecast of the next 5 days
func Forecast(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := openweather.NewClient(os.Getenv("OPENWEATHER_KEY"))

		params := openweather.WeatherArgs{
			ID:    3527646,
			Units: "metric",
		}

		res, err := c.WeatherForecast(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		forecastWeatherResponse := ForecastResponse{WeatherForecast: res}
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(forecastWeatherResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
