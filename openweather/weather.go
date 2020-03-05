package openweather

// Weather object
type Weather struct {
	Weather []MetaWeather `json:"weather,omitempty"`
	Main    Main          `json:"main"`
	Wind    Wind          `json:"wind"`
	Clouds  Clouds        `json:"clouds"`
	DtTxt   string        `json:"dt_txt,omitempty"`
	Name    string        `json:"name,omitempty"`
	Cod     int           `json:"cod"`
	Message string        `json:"message,omitempty"`
}

// MetaWeather struct
type MetaWeather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main info of the current weather
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// Wind data
type Wind struct {
	Speed float64 `json:"speed"`
}

// Clouds data
type Clouds struct {
	All int `json:"all"`
}
