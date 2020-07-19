package openweather

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.openweathermap.org",
	Path:   "/data/2.5/",
}

// Client is the client used to connect this app with openweather API.
type Client struct {
	apiKey     string
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient returns a new pointer of an instance of the Client.
// It expects a valid API Key as a parameter.
func NewClient(apiKey string) *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		apiKey:     apiKey,
		baseURL:    &baseURL,
		httpClient: c,
	}
}

// CurrentWeather gets the current weather in a specific zones given some parameters
func (c *Client) CurrentWeather(args WeatherArgs) (Weather, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: weatherPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)

	if err != nil {
		return Weather{Message: err.Error()}, err
	}

	req.Header.Add("Accept", "application/json")

	params := args.QueryParams()
	params.Add("APPID", c.apiKey)

	req.URL.RawQuery = params.Encode()
	res, err := c.httpClient.Do(req)

	if err != nil {
		return Weather{Message: err.Error()}, err
	}
	defer res.Body.Close()

	var weatherResponse = Weather{}

	if err := json.NewDecoder(res.Body).Decode(&weatherResponse); err != nil {
		return Weather{Message: err.Error()}, err
	}

	return weatherResponse, nil
}

// WeatherForecast gets the forecasting of the weather of the next 5 days
func (c *Client) WeatherForecast(args WeatherArgs) (Forecast, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: forecastPath})
	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return Forecast{ErrorMessage: err.Error()}, err
	}

	req.Header.Add("Accept", "application/json")

	params := args.QueryParams()
	params.Add("APPID", c.apiKey)
	req.URL.RawQuery = params.Encode()

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Forecast{ErrorMessage: err.Error()}, err
	}

	defer res.Body.Close()

	var forecastResponse = Forecast{}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(&forecastResponse)
	default:
		var forecastError = ForecastError{}
		err = json.NewDecoder(res.Body).Decode(&forecastError)
		forecastResponse.ErrorMessage = forecastError.Message
	}

	forecastResponse.StatusCode = res.StatusCode

	return forecastResponse, err
}
