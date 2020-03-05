package controllers

import (
	"net/http"
	"net/url"
	"os"

	"github.com/danvergara/dashservergo/openweather"
	"github.com/gin-gonic/gin"
)

// CurrentWeather Returns the main data of the current Weather
func CurrentWeather(c *gin.Context) {
	weatherClient := openweather.Client{
		APIKey: os.Getenv("OPENWEATHER_KEY"),
	}

	params := url.Values{}
	params.Add("id", "3527646")
	params.Add("units", "metric")

	response, err := weatherClient.GetCurrentWeather(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"current-weather": response})
}

// WeatherForecast returns the forecast of the next 5 days
func WeatherForecast(c *gin.Context) {
	weatherClient := openweather.Client{
		APIKey: os.Getenv("OPENWEATHER_KEY"),
	}

	params := url.Values{}
	params.Add("id", "3527646")
	params.Add("units", "metric")

	response, err := weatherClient.GetWeatherForecast(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"weather-forecast": response})
}
