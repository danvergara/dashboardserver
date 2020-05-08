package openweather

import (
	"encoding/json"
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
	resp, err := c.get(weatherPath, queryParams)

	if err != nil {
		return Weather{Message: err.Error()}, err
	}

	defer resp.Body.Close()

	var weatherResponse = Weather{}
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	weatherResponse.StatusCode = resp.StatusCode

	return weatherResponse, err
}

// GetWeatherForecast gets the forecasting of the weather of the next 5 days
func (c *Client) GetWeatherForecast(queryParams url.Values) (Forecast, error) {
	resp, err := c.get(forecastPath, queryParams)

	if err != nil {
		return Forecast{ErrorMessage: err.Error()}, err
	}

	defer resp.Body.Close()

	var forecastResponse = Forecast{}

	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&forecastResponse)
	} else {
		var forecastError = ForecastError{}
		err = json.NewDecoder(resp.Body).Decode(&forecastError)
		forecastResponse.ErrorMessage = forecastError.Message
	}

	forecastResponse.StatusCode = resp.StatusCode

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

func (c *Client) get(path string, params url.Values) (resp *http.Response, err error) {
	c.setClient()
	u := c.buildURL(path, params)
	resp, err = c.httpClient.Get(u.String())
	return
}

func (c *Client) setClient() {
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: time.Second * 3}
	}
}

func (c *Client) setBaseURL() string {
	if c.BaseURL == "" {
		return defaultBaseURL
	}
	return c.BaseURL
}
