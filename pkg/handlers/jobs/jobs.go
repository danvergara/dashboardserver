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

//lint:file-ignore U1000 Ignore all unused code, it's generated

// A list of jobs for software engineers  from the Github Jobs API
// swagger:response jobsResponse
type jobsResponseWrapper struct {
	// in: body
	Body jobsResponse
}

// The necessary params to filter the list of jobs
// swagger:parameters jobs
type jobsParameters struct {
	// The desired job description, mostly used to specify the programming language
	// in: query
	Description string `json:"description"`
}

// swagger:route GET /v1/jobs jobs
//
// Returns a jobs list for software engineers from Github Jobs API
//
// Produces:
// - application/json
//
// Responses:
//	200: jobsResponse

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
