package healthcheck

import (
	"io"
	"log"
	"net/http"

	"github.com/danvergara/dashboardserver/pkg/application"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

// swagger:response healthcheckResponse
type healthcheckResponseWrapper struct {
	// The expected healthcheck response to see if the service is running
	// in: body
	Body healthcheckResponse
}

type healthcheckResponse struct {
	Alive bool `json:"alive"`
}

// swagger:route GET /_healthcheck healthcheck
//
// Checks if the application is running
//
// Responses:
// 	200: healthcheckResponse

// Healthcheck handle func
func Healthcheck(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		_, err := io.WriteString(w, `{"alive": true}`)

		if err != nil {
			log.Println(err)
		}
	}
}
