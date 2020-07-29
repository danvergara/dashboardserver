package trendingos

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTrendingRepositoriesPython(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pythonResponse, err := ioutil.ReadFile("python_response.json")
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(pythonResponse))
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

	params := TrendingRepositoryArgs{
		Language: "python",
		Since:    "weekly",
	}

	res, err := c.TrendingRepositories(params)
	assert.NoError(t, err)
	assert.Equal(t, 25, len(res))
	assert.Equal(t, "vt-vl-lab", res[0].Author)
	assert.Equal(t, "python-cheatsheet", res[1].Name)
	assert.Equal(t, "Face Analysis Project on MXNet", res[2].Description)
	assert.Equal(t, 5, len(res[2].BuiltBy))
}

func TestTrendingRepositoriesRuby(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rubyResponse, err := ioutil.ReadFile("ruby_response.json")
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(rubyResponse))
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

	params := TrendingRepositoryArgs{
		Language: "ruby",
		Since:    "weekly",
	}

	res, err := c.TrendingRepositories(params)
	assert.NoError(t, err)
	assert.Equal(t, 25, len(res))
	assert.Equal(t, "https://github.com/faker-ruby.png", res[1].Avatar)
	assert.Equal(t, "https://github.com/elastic/logstash", res[4].URL)
	assert.Equal(t, 14487, res[5].Stars)
	assert.Equal(t, 1981, res[6].Forks)
}

func TestTrendingRepositoriesGolang(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goResponse, err := ioutil.ReadFile("go_response.json")
		if err != nil {
			t.Error(err)
		}

		_, err = w.Write([]byte(goResponse))
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

	params := TrendingRepositoryArgs{
		Language: "golang",
		Since:    "weekly",
	}

	res, err := c.TrendingRepositories(params)
	assert.NoError(t, err)
	assert.Equal(t, 25, len(res))
	assert.Equal(t, "Go", res[1].Language)
	assert.Equal(t, "#00ADD8", res[4].LanguageColor)
	assert.Equal(t, 337, res[1].CurrentPeriodStars)
	assert.Equal(t, "Dreamacro", res[1].BuiltBy[0].Username)
}
