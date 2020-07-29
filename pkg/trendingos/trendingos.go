package trendingos

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const repositoriesPath string = "repositories"

var baseURL = url.URL{
	Scheme: "https",
	Host:   "ghapi.huchen.dev",
	Path:   "/",
}

// Client is the main object/struct used to connect this app with the github-trending-api service
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient returns a pointer to a new instance of the Client
func NewClient() *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		baseURL:    &baseURL,
		httpClient: c,
	}
}

// TrendingRepositories returns a list of trending repositories given certain parameters
func (c *Client) TrendingRepositories(args TrendingRepositoryArgs) ([]TrendingRepository, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: repositoriesPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var respositories []TrendingRepository
	if err := json.NewDecoder(res.Body).Decode(&respositories); err != nil {
		return nil, err
	}

	return respositories, nil
}
