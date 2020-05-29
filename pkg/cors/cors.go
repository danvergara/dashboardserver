package cors

import "net/http"

// EnableCors Cross Origin Resource Sharing
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
