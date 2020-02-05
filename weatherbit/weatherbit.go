package weatherbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Client is the client used to connect to the Weatherbit API
type Client struct {
	APIKey     string
	BaseURL    string
	httpClient *http.Client
}

// GetCurrentWeather gets the current weather in a specific zones given some parameters
func (c *Client) GetCurrentWeather(queryParams url.Values) (CurrentWeatherResponse, error) {
	req, err := c.newRequest("GET", currentPath, queryParams, nil)
	if err != nil {
		return CurrentWeatherResponse{Error: err.Error()}, err
	}

	if c.APIKey == "" {
		return CurrentWeatherResponse{Error: "API key is required."}, fmt.Errorf("api key is required")
	}

	var currentResponse = CurrentWeatherResponse{}
	_, err = c.do(req, &currentResponse)

	if currentResponse.Error != "" {
		return currentResponse, fmt.Errorf(currentResponse.Error)
	}

	return currentResponse, err
}

// GetWeatherForecast returns a forecast in 1 day intervals from any point on the planet, up to 16 days in the future
func (c *Client) GetWeatherForecast(queryParams url.Values) (ForecastWeatherResponse, error) {
	req, err := c.newRequest("GET", forecastPath, queryParams, nil)
	if err != nil {
		return ForecastWeatherResponse{Error: err.Error()}, err
	}

	if c.APIKey == "" {
		return ForecastWeatherResponse{Error: "API key is required."}, fmt.Errorf("api key is required")
	}

	var forecastResponse = ForecastWeatherResponse{}
	_, err = c.do(req, &forecastResponse)

	if forecastResponse.Error != "" {
		return forecastResponse, fmt.Errorf(forecastResponse.Error)
	}

	return forecastResponse, err
}

func (c *Client) buildURL(path string, queryParams url.Values) *url.URL {
	u, err := url.Parse(c.urlBase())

	if err != nil {
		log.Fatal(err)
	}

	queryParams.Add("key", c.APIKey)

	u.Path = path
	u.RawQuery = queryParams.Encode()
	return u
}

func (c *Client) newRequest(method, path string, queryParams url.Values, body interface{}) (*http.Request, error) {
	u := c.buildURL(path, queryParams)

	var buf io.ReadWriter = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if body != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		c.setHeader(req)
	}

	return req, nil

}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	c.setClient()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) setClient() {
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: 3 * time.Second}
	}
}

func (c *Client) setHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func (c *Client) urlBase() string {
	if c.BaseURL == "" {
		return defaultBaseURL
	}

	return c.BaseURL
}
