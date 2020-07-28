package topnews

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/newsapigo"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

// The Top News response in json format
// swagger:response articlesReponse
type articlesReponseWrapper struct {
	// in: body
	Body ArticlesReponse
}

// ArticlesReponse substitutes the old response {"news": response.Articles}
type ArticlesReponse struct {
	News []newsapigo.Article `json:"news"`
}

// swagger:route GET /v1/top-news topnews
//
// Returns a list of the top business news in MX
//
// Produces:
// - application/json
//
// Responses:
//	200: articlesReponse

// TopNews returns the top news about business in Mexico
// As we can see, newsapigo is the library that does the hard work
func TopNews(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Using the brand new NewClient constructor function to create a newsapigo Client
		c := newsapigo.NewClient(os.Getenv("NEWSAPI_KEY"))

		// to create the query params, create an TopHeadlinesArgs instance an fil the requred fields
		// In this case: Country and Category
		queryParams := newsapigo.TopHeadlinesArgs{
			Country:  "mx",
			Category: "business",
		}

		// Perfom the http request using the TopHeadlines method
		response, err := c.TopHeadlines(queryParams)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Create the endpoint response using the previous newsapigo's response
		articlesResponse := ArticlesReponse{News: response.Articles}

		// Enconde the response
		encoder := json.NewEncoder(w)
		if err = encoder.Encode(articlesResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
