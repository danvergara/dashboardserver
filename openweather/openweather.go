package openweather

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

// Client is the client used to connect this app with openweather API
type Client struct {
	APIKey     string
	BaseURL    string
	httpClient *http.Client
}

// GetCurrentWeather gets the current weather in a specific zones given some parameters
func (c *Client) GetCurrentWeather(queryParams url.Values) (Weather, error) {
	req, err := c.newRequest("GET", weatherPath, queryParams, nil)

	if err != nil {
		return Weather{Message: err.Error()}, err
	}

	if c.APIKey == "" {
		return Weather{Message: "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."}, fmt.Errorf("API key is required")
	}

	var weatherResponse = Weather{}

	_, err = c.do(req, &weatherResponse)

	if weatherResponse.Message != "" {
		return weatherResponse, fmt.Errorf(weatherResponse.Message)
	}

	return weatherResponse, err
}

// GetWeatherForecast gets the forecasting of the weather of the next 5 days
func (c *Client) GetWeatherForecast(queryParams url.Values) (Forecast, error) {
	req, err := c.newRequest("GET", forecastPath, queryParams, nil)

	if err != nil {
		return Forecast{Message: err.Error()}, err
	}

	if c.APIKey == "" {
		return Forecast{Message: "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."}, fmt.Errorf("API key is required")
	}

	var forecastResponse = Forecast{}

	_, err = c.do(req, &forecastResponse)

	if forecastResponse.Message != "" {
		return forecastResponse, fmt.Errorf(forecastResponse.Message)
	}

	return forecastResponse, err
}

func (c *Client) buildURL(path string, params url.Values) *url.URL {
	u, err := url.Parse(c.setBaseURL())

	if err != nil {
		log.Fatal(err)
	}

	params.Add("APPID", c.APIKey)

	u.Path = path
	u.RawQuery = params.Encode()
	return u
}

func (c *Client) newRequest(method, path string, params url.Values, body interface{}) (*http.Request, error) {
	u := c.buildURL(path, params)

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
		c.httpClient = &http.Client{Timeout: time.Second * 3}
	}
}

func (c *Client) setHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func (c *Client) setBaseURL() string {
	if c.BaseURL == "" {
		return defaultBaseURL
	}
	return c.BaseURL
}
