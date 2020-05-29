package healthcheck

import (
	"io"
	"log"
	"net/http"

	"github.com/danvergara/dashboardserver/pkg/application"
)

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
