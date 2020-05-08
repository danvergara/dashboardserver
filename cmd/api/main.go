package main

import (
	"github.com/danvergara/dashboardserver/pkg/handlers/economics"
	"github.com/danvergara/dashboardserver/pkg/handlers/topnews"
	"github.com/danvergara/dashboardserver/pkg/handlers/weather"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/v1")
	{
		v1.GET("/currency-exchange", economics.CurrrencyExchange)
		v1.GET("/historical-currency-rates", economics.HistoricalCurrencyRates)
		v1.GET("/top-news", topnews.TopNews)
		v1.GET("/current-weather", weather.CurrentWeather)
		v1.GET("/weather-forecast", weather.WeatherForecast)
	}

	r.Run(":8000")
}
