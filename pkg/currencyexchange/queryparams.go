package currencyexchange

import (
	"net/url"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// LatestArgs represents the query params used to filter a request to /latest.
type LatestArgs struct {
	// Rates are quoted against the Euro by default. Quote against a different currency by setting the base parameter in your request.
	// GET /latest?base=USD
	Base string
	// Request specific exchange rates by setting the symbols parameter.
	// GET /latest?symbols=USD,GBP
	Symbols []string
}

// QueryParams returns a url.Values instance with non-empty values of LatestArgs
func (args LatestArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.Base != "" {
		q.Add("base", args.Base)
	}

	if len(args.Symbols) > 0 {
		q.Add("symbols", strings.Join(args.Symbols, ","))
	}

	return q
}

// HistoryArgs represents the query params used to filter a reques to /history
type HistoryArgs struct {
	// Get historical rates for a time period.
	// GET /history?start_at=2018-01-01&end_at=2018-09-01
	StartAt time.Time
	EndAt   time.Time
	// Limit results to specific exchange rates to save bandwidth with the symbols parameter.
	// GET /history?start_at=2018-01-01&end_at=2018-09-01&symbols=ILS,JPY
	Symbols []string
	// Quote the historical rates against a different currency.
	// GET /history?start_at=2018-01-01&end_at=2018-09-01&base=USD
	Base string
}

// QueryParams returns a url.Values instance with non-empty values of HistoryArgs
func (args HistoryArgs) QueryParams() url.Values {
	q := make(url.Values)

	if !args.StartAt.IsZero() {
		q.Add("start_at", args.StartAt.Format(dateLayout))
	}

	if !args.EndAt.IsZero() {
		q.Add("end_at", args.EndAt.Format(dateLayout))
	}

	if args.Base != "" {
		q.Add("base", args.Base)
	}

	if len(args.Symbols) > 0 {
		q.Add("symbols", strings.Join(args.Symbols, ","))
	}

	return q
}
