package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/danvergara/newsapigo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	APIKey := os.Getenv("NEWSAPI_KEY")

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
