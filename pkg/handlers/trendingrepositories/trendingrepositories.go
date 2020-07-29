package trendingrepositories

import (
	"encoding/json"
	"net/http"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/logger"
	"github.com/danvergara/dashboardserver/pkg/trendingos"
)

// The treding repositories on Github, filtered by date or language
// swagger:response repositoriesResponse
type repositoriesResponseWrapper struct {
	// in: body
	Body RepositoriesResponse
}

// The necessary params to filter the list of treding repositories
// swagger:parameters repositories
type trendingRepositoriesParameters struct {
	// The programming language
	// in: query
	Language string `json:"language"`
	// since: optional, default to daily, possible values: daily, weekly and monthly
	// in: query
	Since string `json:"since"`
	// spoken_language_code: optional, list trending repositories of certain spoken languages
	// in: query
	SpokenLanguageCode string `json:"spoken_language_code"`
}

// RepositoriesResponse is the actual response {"respositories": []}
type RepositoriesResponse struct {
	Repositories []trendingos.TrendingRepository `json:"repositories"`
}

// swagger:route GET /v1/repositories repositories
//
// Returns the list of the treding repositories on Github
//
// Produces:
// - application/json
//
// Responses:
//	200: repositoriesResponse

// Get retuns a list of trending repositories on Github
func Get(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		language := r.URL.Query().Get("language")
		since := r.URL.Query().Get("since")
		spokenLanguage := r.URL.Query().Get("spoken_language_code")

		c := trendingos.NewClient()

		params := trendingos.TrendingRepositoryArgs{
			Language:           language,
			Since:              since,
			SpokenLanguageCode: spokenLanguage,
		}

		res, err := c.TrendingRepositories(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		repositoriesResponse := RepositoriesResponse{Repositories: res}

		encoder := json.NewEncoder(w)
		if err = encoder.Encode(repositoriesResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error.Println(err)
		}
	}
}
