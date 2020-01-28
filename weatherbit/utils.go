package weatherbit

import (
	"net/http"
	"net/url"
)

func buildRequest(url *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return req, err
	}
	return req, nil
}
