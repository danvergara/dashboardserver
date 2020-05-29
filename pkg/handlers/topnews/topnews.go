package topnews

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/cors"
	"github.com/danvergara/newsapigo"
)

// ArticlesReponse substitutes the old response {"news": response.Articles}
type ArticlesReponse struct {
	News []newsapigo.Article `json:"news"`
}

// TopNews returns the top news in Mexico
func TopNews(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.EnableCors(&w)

		client := newsapigo.NewsClient{
			APIKey: os.Getenv("NEWSAPI_KEY"),
		}

		params := url.Values{}
		params.Add("country", "mx")
		params.Add("category", "business")
		response, err := client.GetTopHeadlines(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		encoder := json.NewEncoder(w)
		articlesResponse := ArticlesReponse{News: response.Articles}
		if err = encoder.Encode(articlesResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
