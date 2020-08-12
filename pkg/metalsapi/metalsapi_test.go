package metalsapi

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeSeriesXAU(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := ioutil.ReadFile("response.json")

		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(res))
		if err != nil {
			t.Error(err)
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)
	testClient := &http.Client{Timeout: time.Minute}

	c := Client{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "API_KEY",
	}

	startDate, _ := time.Parse(dateLayout, "2020-08-06")
	endDate, _ := time.Parse(dateLayout, "2020-08-11")

	params := TimeSeriesArgs{
		StartDate: startDate,
		EndDate:   endDate,
		Base:      "XAU",
		Symbols:   []string{"USD"},
	}

	resp, err := c.TimeSeries(params)

	assert.NoError(t, err)
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, true, resp.Timeseries)
	assert.Equal(t, "2020-08-06", resp.StartDate)
	assert.Equal(t, "2020-08-11", resp.EndDate)
	assert.Equal(t, "XAU", resp.Base)
	assert.Equal(t, 6, len(resp.Rates))
	assert.Equal(t, 2068.5316962146076, resp.Rates["2020-08-06"]["USD"])
	assert.Equal(t, 1910.4029425707276, resp.Rates["2020-08-11"]["USD"])
	assert.Equal(t, "per ounce", resp.Unit)
}

func TestWrongAPIKey(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := ioutil.ReadFile("api_key_error.json")

		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(res))
		if err != nil {
			t.Error(err)
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)
	testClient := &http.Client{Timeout: time.Minute}

	c := Client{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "WRONG_API_KEY",
	}

	startDate, _ := time.Parse(dateLayout, "2020-08-06")
	endDate, _ := time.Parse(dateLayout, "2020-08-11")

	params := TimeSeriesArgs{
		StartDate: startDate,
		EndDate:   endDate,
		Base:      "XAU",
		Symbols:   []string{"USD"},
	}

	resp, err := c.TimeSeries(params)

	assert.NoError(t, err)
	assert.Equal(t, false, resp.Success)
	assert.Equal(t, 101, resp.Error.Code)
	assert.Equal(t, "invalid_access_key", resp.Error.Type)
	assert.Equal(t, "No API Key was specified or an invalid API Key was specified.", resp.Error.Info)
}

func TestWrongSymbols(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := ioutil.ReadFile("api_key_error.json")

		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(res))
		if err != nil {
			t.Error(err)
		}
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)
	testClient := &http.Client{Timeout: time.Minute}

	c := Client{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "API_KEY",
	}

	startDate, _ := time.Parse(dateLayout, "2020-08-06")
	endDate, _ := time.Parse(dateLayout, "2020-08-11")

	params := TimeSeriesArgs{
		StartDate: startDate,
		EndDate:   endDate,
		Base:      "XAU",
		Symbols:   []string{"USD"},
	}

	resp, err := c.TimeSeries(params)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp.Rates))
}
