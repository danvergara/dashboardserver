package openweather

// Forecast has a list of weather objects
type Forecast struct {
	List    []Weather `json:"list"`
	Cod     string    `json:"cod"`
	Message string    `json:"message:omitempty"`
}
