package timeseries

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/logger"
	"github.com/danvergara/dashboardserver/pkg/metalsapi"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

// The exchange rates time series from metals api
// swagger:response timeSeriesResponse
type timeSeriesResponseWrapper struct {
	// in: body
	Body timeSeriesResponse
}

type timeSeriesResponse struct {
	TimeSeries *metalsapi.Response `json:"time_series"`
}

// The necessary params to specify the base
// swagger:parameters timeseries
type timeSeriesParameters struct {
	// The currency code or metal code of your preferred base currency
	// in: query
	Base string `json:"base"`
}

// swagger:route GET /v1/timeseries timeseries
//
// Returns the exchange rates in time series in dollars given a base currency
//
// Produces:
// - application/json
//
// Responses:
//	200: timeSeriesResponse

// Get returns the exchange rates from metals api given a base currency
func Get(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		base := r.URL.Query().Get("base")

		c := metalsapi.NewClient(os.Getenv("METALS_API_KEY"))

		now := time.Now()

		params := metalsapi.TimeSeriesArgs{
			Base:      base,
			Symbols:   []string{"USD"},
			EndDate:   now,
			StartDate: now.AddDate(0, 0, -5),
		}

		res, err := c.TimeSeries(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		series := timeSeriesResponse{TimeSeries: res}

		encoder := json.NewEncoder(w)
		if err = encoder.Encode(series); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error.Println(err)
		}
	}
}
