package openweather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeatherForecastWithoutCityID(t *testing.T) {
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
