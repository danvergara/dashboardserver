package controllers

import (
	"net/http"
	"net/url"
	"os"

	"github.com/danvergara/dashservergo/weatherbit"
	"github.com/gin-gonic/gin"
)

// CurrentWeather Returns the main data of the current Weather
func CurrentWeather(c *gin.Context) {
	weatherClient := weatherbit.Client{
		APIKey: os.Getenv("WEATHERBIT_KEY"),
	}

	params := url.Values{}
	params.Add("city", "Mexico")

	response, err := weatherClient.GetCurrentWeather(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"current-weather": response})
}

// WeatherForecast returns the forecast of the next 5 days
func WeatherForecast(c *gin.Context) {
	weatherClient := weatherbit.Client{
		APIKey: os.Getenv("WEATHERBIT_KEY"),
	}

	params := url.Values{}
	params.Add("city", "Mexico")
	params.Add("days", "5")

	response, err := weatherClient.GetWeatherForecast(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"weather-forecast": response})
}
