package currencyexchange

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/danvergara/dashboardserver/pkg/logger"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.exchangeratesapi.io",
}

// Client is the main object used to connect this with exchangerates API
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient returns a pointer of a new instance of Client.
func NewClient() *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		baseURL:    &baseURL,
		httpClient: c,
	}
}

// LatestCurrencyExchange returns an object with the response of the currencyexchange API
func (c *Client) LatestCurrencyExchange(args LatestArgs) (*ExchangeRateData, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: latestPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	params := args.QueryParams()

	req.URL.RawQuery = params.Encode()
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var exrd ExchangeRateData
		if err := json.NewDecoder(res.Body).Decode(&exrd); err != nil {
			return nil, err
		}
		return &exrd, nil
	case 400, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// HistoricalCurrencyRate returns an object with the response of the currencyexchange API
func (c *Client) HistoricalCurrencyRate(args HistoryArgs) (*HistoricalExchangeRateData, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: historyPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	params := args.QueryParams()

	logger.Info.Printf("url: %v", params)
	req.URL.RawQuery = params.Encode()
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var hexrd HistoricalExchangeRateData
		if err := json.NewDecoder(res.Body).Decode(&hexrd); err != nil {
			return nil, err
		}
		return &hexrd, nil
	case 400, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}
