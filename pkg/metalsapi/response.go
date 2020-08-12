package metalsapi

// Response is the struct that the maps the values from the metals-api
type Response struct {
	Success    bool                          `json:"success"`
	Timeseries bool                          `json:"timeseries,omitempty"`
	StartDate  string                        `json:"start_date,omitempty"`
	EndDate    string                        `json:"end_date,omitempty"`
	Date       string                        `json:"date,omitempty"`
	Base       string                        `json:"base"`
	Rates      map[string]map[string]float64 `json:"rates"`
	Unit       string                        `json:"unit"`
	Error      *Error                        `json:"error,omitempty"`
}

// Error object has all the relevant information when an errors occurs
type Error struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}
