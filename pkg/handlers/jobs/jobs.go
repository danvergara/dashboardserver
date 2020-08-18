package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/githubjobs"
	"github.com/danvergara/dashboardserver/pkg/logger"
)

type jobsResponse struct {
	Jobs []githubjobs.Job `json:"jobs"`
}

// Get returns a list of job positions for software engineers from Github Jobs API
func Get(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		desc := r.URL.Query().Get("description")

		c := githubjobs.NewClient()
		params := githubjobs.PositionsArgs{
			Description: desc,
			FullTime:    true,
		}

		jobs, err := c.Positions(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res := jobsResponse{Jobs: jobs}

		encoder := json.NewEncoder(w)
		if err = encoder.Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error.Println(err)
		}
	}
}
