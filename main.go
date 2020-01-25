package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/danvergara/newsapigo"
	"github.com/gin-gonic/gin"
	"github.com/mattevans/dinero"
)

func main() {
	r := gin.Default()

	APIKey := os.Getenv("NEWSAPI_KEY")

	r.GET("/currency-exchange", func(c *gin.Context) {
		currencyClient := dinero.NewClient(
			os.Getenv("OPEN_EXCHANGE_APP_ID"),
			"USD",
			20*time.Minute,
		)

		rsp, err := currencyClient.Rates.Get("MXN")

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"dolar-peso": rsp})
	})

	r.GET("/top-news", func(c *gin.Context) {
		client := newsapigo.NewsClient{
			APIKey: APIKey,
		}

		params := url.Values{}
		params.Add("conutry", "mx")
		params.Add("category", "general")
		response, err := client.GetTopHeadlines(params)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"news": response.Articles})
	})

	r.Run()
}
