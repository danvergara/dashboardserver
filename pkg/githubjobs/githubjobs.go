package githubjobs

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const positionsPath string = "positions.json?"

var baseURL = url.URL{
	Scheme: "https",
	Host:   "jobs.github.com",
	Path:   "/",
}

// Client is public interface who interacts with the REST API
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

// Positions return a list of Jobs
func (c *Client) Positions(args PositionsArgs) ([]Job, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: positionsPath})

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
	var jobs []Job

	if err := json.NewDecoder(res.Body).Decode(&jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}
