package openweather

// Forecast has a list of weather objects
type Forecast struct {
	List         []Weather `json:"list"`
	ErrorMessage string    `json:"error_message"`
	StatusCode   int       `json:"status_code"`
}

// ForecastError has the message error
type ForecastError struct {
	Message string `json:"message"`
}
