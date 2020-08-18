package githubjobs

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPositionsDescriptionPython(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pythonJobs, err := ioutil.ReadFile("python_jobs.json")
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(pythonJobs))
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
	}

	params := PositionsArgs{
		Description: "python",
	}

	res, err := c.Positions(params)

	assert.NoError(t, err)
	assert.Equal(t, "Full Time", res[0].Type)
	assert.Equal(t, 50, len(res))
	assert.Equal(t, "Senior Data Engineer", res[0].Title)
	assert.Equal(t, "ASML", res[2].Company)
	assert.Equal(t, "Veldhoven", res[2].Location)
	assert.Equal(t, "a184a967-12b1-4c89-a935-1b2c952dd3ed", res[3].ID)
}

func TestPositionsFullTime(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pythonJobs, err := ioutil.ReadFile("python_full_time.json")
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(pythonJobs))
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
	}

	params := PositionsArgs{
		Description: "python",
		FullTime:    true,
	}

	res, err := c.Positions(params)

	assert.NoError(t, err)
	assert.Equal(t, 34, len(res))
	assert.Equal(t, "Senior Python Engineer", res[0].Title)
	assert.Equal(t, "SteepRock, Inc", res[0].Company)
	assert.Equal(t, "bde7d20c-e0b0-4050-9c81-21f7bf9de87f", res[1].ID)
	assert.Equal(t, "Ludwigshafen am Rhein", res[2].Location)

	for i, job := range res {
		if job.Type != "Full Time" {
			t.Errorf("The job %d has %s type", i, job.Type)
		}
	}
}

func TestPositionsWrongDescription(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pythonJobs, err := ioutil.ReadFile("wrong_description.json")
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(pythonJobs))
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
	}

	params := PositionsArgs{
		Description: "pytngjnva",
	}

	res, err := c.Positions(params)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
}
