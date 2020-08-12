package metalsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	timeseriesPath string = "timeseries"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "metals-api.com/api/",
}

// Client is the main object used to Connect this app with metals-api API
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	apiKey     string
}

// NewClient returns a pointer of a new instance of the Client
func NewClient(apiKey string) *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		apiKey:     apiKey,
		baseURL:    &baseURL,
		httpClient: c,
	}
}

// TimeSeries returns the response from metals-api in struct format
func (c *Client) TimeSeries(args TimeSeriesArgs) (*Response, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: timeseriesPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	params := args.QueryParams()
	params.Add("access_key", c.apiKey)

	req.URL.RawQuery = params.Encode()
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
