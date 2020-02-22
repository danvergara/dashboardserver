package currencyexchange

// HistoricalExchangeRateData stores the hisotircal exchage currency rate
type HistoricalExchangeRateData struct {
	Rates   map[string]map[string]float64 `json:"rates"`
	StartAt string                        `json:"start_at"`
	Base    string                        `json:"base"`
	EndAt   string                        `json:"end_at"`
	Error   string                        `json:"error,omitempty"`
}

// ExchangeRateData stores the exchage rate between two currencies
type ExchangeRateData struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Error string             `json:"error,omitempty"`
}
