package openweather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeatherForecastLongList(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorForecastResponse := Forecast{
			List: []Weather{
				{
					Main: Main{
						Temp:      15.17,
						FeelsLike: 11.17,
						TempMin:   9.0,
						TempMax:   19.0,
						Pressure:  1025,
						Humidity:  37,
					},
					Weather: []MetaWeather{
						{
							Main:        "Clear",
							Description: "clear sky",
							Icon:        "01n",
						},
					},
					Wind: Wind{
						Speed: 3.1,
					},
					Clouds: Clouds{
						All: 5,
					},
					DtTxt: "2020-03-03 06:00:00",
				},
				{
					Main: Main{
						Temp:      15.17,
						FeelsLike: 11.17,
						TempMin:   9.0,
						TempMax:   19.0,
						Pressure:  1025,
						Humidity:  37,
					},
					Weather: []MetaWeather{
						{
							Main:        "Clear",
							Description: "clear sky",
							Icon:        "01n",
						},
					},
					Wind: Wind{
						Speed: 3.1,
					},
					Clouds: Clouds{
						All: 5,
					},
					DtTxt: "2020-03-03 06:00:00",
				},
				{
					Main: Main{
						Temp:      15.17,
						FeelsLike: 11.17,
						TempMin:   9.0,
						TempMax:   19.0,
						Pressure:  1025,
						Humidity:  37,
					},
					Weather: []MetaWeather{
						{
							Main:        "Clear",
							Description: "clear sky",
							Icon:        "01n",
						},
					},
					Wind: Wind{
						Speed: 3.1,
					},
					Clouds: Clouds{
						All: 5,
					},
					DtTxt: "2020-03-03 06:00:00",
				},
				{
					Main: Main{
						Temp:      15.17,
						FeelsLike: 11.17,
						TempMin:   9.0,
						TempMax:   19.0,
						Pressure:  1025,
						Humidity:  37,
					},
					Weather: []MetaWeather{
						{
							Main:        "Clear",
							Description: "clear sky",
							Icon:        "01n",
						},
					},
					Wind: Wind{
						Speed: 3.1,
					},
					Clouds: Clouds{
						All: 5,
					},
					DtTxt: "2020-03-03 06:00:00",
				},
				{
					Main: Main{
						Temp:      15.17,
						FeelsLike: 11.17,
						TempMin:   9.0,
						TempMax:   19.0,
						Pressure:  1025,
						Humidity:  37,
					},
					Weather: []MetaWeather{
						{
							Main:        "Clear",
							Description: "clear sky",
							Icon:        "01n",
						},
					},
					Wind: Wind{
						Speed: 3.1,
					},
					Clouds: Clouds{
						All: 5,
					},
					DtTxt: "2020-03-03 06:00:00",
				},
			},
		}

		if r.URL.Path != "/data/2.5/forecast" {
			t.Error("Bad forecast Path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(errorForecastResponse)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("unit", "metric")
	queryParams.Add("id", "3527646")

	resp, err := c.GetWeatherForecast(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 5, len(resp.List))
	assert.Equal(t, 15.17, resp.List[0].Main.Temp)
	assert.Equal(t, 37, resp.List[0].Main.Humidity)
	assert.Equal(t, 3.1, resp.List[0].Wind.Speed)
	assert.Equal(t, "01n", resp.List[0].Weather[0].Icon)
	assert.Equal(t, 9.0, resp.List[0].Main.TempMin)
	assert.Equal(t, 19.0, resp.List[0].Main.TempMax)
}

func TestGetWeatherForecastWithIntMessage(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ForecastResponse := []byte(`
			{
			  "cod": "200",
			  "message": 0,
			  "cnt": 40,
			  "list": [
					{
						"dt": 1586476800,
						"main": {
							"temp": 24.08,
							"feels_like": 19.73,
							"temp_min": 23.34,
							"temp_max": 24.08,
							"pressure": 1010,
							"sea_level": 1010,
							"grnd_level": 765,
							"humidity": 14,
							"temp_kf": 0.74
						},
						"weather": [
							{
								"id": 803,
								"main": "Clouds",
								"description": "broken clouds",
								"icon": "04d"
							}
						],
						"clouds": {
							"all": 72
						},
						"wind": {
							"speed": 2.47,
							"deg": 225
						},
						"sys": {
							"pod": "d"
						},
						"dt_txt": "2020-04-10 00:00:00"
					}
				]
			}
		`)

		w.Write(ForecastResponse)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("unit", "metric")
	queryParams.Add("id", "3527646")

	resp, err := c.GetWeatherForecast(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp.List))
	assert.Equal(t, 24.08, resp.List[0].Main.Temp)
	assert.Equal(t, 14, resp.List[0].Main.Humidity)
	assert.Equal(t, 2.47, resp.List[0].Wind.Speed)
	assert.Equal(t, "04d", resp.List[0].Weather[0].Icon)
	assert.Equal(t, 23.34, resp.List[0].Main.TempMin)
	assert.Equal(t, 24.08, resp.List[0].Main.TempMax)
}

func TestGetWeatherForecastWithoutCountryID(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ForecastResponse := []byte(`
			{
				"cod": "400",
				"message": "Nothing to geocode"
			}
		`)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(ForecastResponse)
	}))

	defer sv.Close()

	c := Client{
		APIKey:  "FAKE_API_KEY",
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	queryParams.Add("unit", "metric")

	resp, err := c.GetWeatherForecast(queryParams)

	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	assert.Nil(t, err)
	assert.Equal(t, "Nothing to geocode", resp.ErrorMessage)
}

func TestGetWeatherForecastWithoutAPIKey(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ForecastResponse := []byte(`
			{
				"cod": "401",
				"message": "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."
			}
		`)

		w.WriteHeader(http.StatusUnauthorized)
		w.Write(ForecastResponse)
	}))

	defer sv.Close()

	c := Client{
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}

	resp, err := c.GetWeatherForecast(queryParams)

	assert.Equal(t, resp.StatusCode, http.StatusUnauthorized)
	assert.Nil(t, err)
	assert.Equal(t, "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info.", resp.ErrorMessage)
}
