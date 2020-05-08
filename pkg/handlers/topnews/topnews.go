package topnews

import (
	"net/http"
	"net/url"
	"os"

	"github.com/danvergara/newsapigo"
	"github.com/gin-gonic/gin"
)

// TopNews returns the top news in Mexico
func TopNews(c *gin.Context) {
	client := newsapigo.NewsClient{
		APIKey: os.Getenv("NEWSAPI_KEY"),
	}

	params := url.Values{}
	params.Add("country", "mx")
	params.Add("category", "business")
	response, err := client.GetTopHeadlines(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"news": response.Articles})
}
