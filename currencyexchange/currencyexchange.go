package currencyexchange

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

// Client is the client used to connecto the exchangeratesapi API
type Client struct {
	BaseURL    string
	httpClient *http.Client
}

// GetLatestCurrencyExchange gets the lates currency exchange
func (c *Client) GetLatestCurrencyExchange(queryParams url.Values) (ExchangeRateData, error) {
	req, err := c.newRequest("GET", latestExchangeRatePath, queryParams, nil)
	if err != nil {
		return ExchangeRateData{Error: err.Error()}, err
	}

	var latestExchangeRateResponse = ExchangeRateData{}
	_, err = c.do(req, &latestExchangeRateResponse)

	if latestExchangeRateResponse.Error != "" {
		return latestExchangeRateResponse, fmt.Errorf(latestExchangeRateResponse.Error)
	}

	return latestExchangeRateResponse, err
}

// GetHistoricalCurrencyRate gets the historical currency rates given the start and end dates
func (c *Client) GetHistoricalCurrencyRate(queryParams url.Values) (HistoricalExchangeRateData, error) {
	req, err := c.newRequest("GET", historyExchangeRatePath, queryParams, nil)

	if err != nil {
		return HistoricalExchangeRateData{Error: err.Error()}, err
	}

	var historicalExchangeRateResponse = HistoricalExchangeRateData{}
	_, err = c.do(req, &historicalExchangeRateResponse)

	if historicalExchangeRateResponse.Error != "" {
		return historicalExchangeRateResponse, fmt.Errorf(historicalExchangeRateResponse.Error)
	}

	return historicalExchangeRateResponse, err
}

func (c *Client) buildURL(path string, queryParams url.Values) *url.URL {
	u, err := url.Parse(c.urlBase())

	if err != nil {
		log.Fatal(err)
	}

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
