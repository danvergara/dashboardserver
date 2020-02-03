package weatherbit

// CurrentWeatherResponse stores the data received from the request to the current endpoint
type CurrentWeatherResponse struct {
	Data  []CurrentWeatherData `json:"data"`
	Error string               `json:"error,omitempty"`
}

// CurrentWeatherData has the main data of the current weather
type CurrentWeatherData struct {
	ObTime   string  `json:"ob_time"`
	Clouds   int     `json:"clouds"`
	SolarRad float64 `json:"solar_rad"`
	CityName string  `json:"city_name"`
	WinSpd   float64 `json:"win_spd"`
	Precip   int     `json:"precip,omitempty"`
	Temp     float64 `json:"temp"`
	DateTime string  `json:"datetime"`
}

// ForecastWeatherResponse stores the data received from the forescast endpoint
type ForecastWeatherResponse struct {
	CityName string         `json:"city_name"`
	Data     []ForecastData `json:"data,omitempty"`
	Error    string         `json:"error,omitempty"`
}

// ForecastData represents the forescast of a specific day in the future
type ForecastData struct {
	ValidDate  string  `json:"valid_date"`
	WinSpd     float64 `json:"win_spd"`
	HighTemp   float64 `json:"high_temp"`
	LowTemp    float64 `json:"low_temp"`
	MaxTemp    float64 `json:"max_temp"`
	MinTemp    float64 `json:"min_temp"`
	AppMaxTemp float64 `json:"app_max_temp"`
	AppMinTemp float64 `json:"app_min_temp"`
	Precip     float64 `json:"precip"`
	Clouds     int     `json:"clouds"`
	Pop        int     `json:"pop"`
}
