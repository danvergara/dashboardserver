package main

import (
	"github.com/danvergara/dashservergo/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/currency-exchange", controllers.CurrrencyExchange)
		v1.GET("/top-news", controllers.TopNews)
		v1.GET("/current-weather", controllers.CurrentWeather)
		v1.GET("/weather-forecast", controllers.WeatherForecast)
	}

	r.Run()
}
