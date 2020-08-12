package metalsapi

import (
	"net/url"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

// TimeSeriesArgs is the object used to map the query params passed byt the user
type TimeSeriesArgs struct {
	// [required] The start date of your preferred timeframe.
	StartDate time.Time
	// [required] The end date of your preferred timeframe.
	EndDate time.Time
	// [optional] Enter the three-letter currency code or metal code of your preferred base currency.
	Base string
	// [optional] Enter one currency or metal codes to limit the output.
	Symbols []string
}

// QueryParams returns a url.Values instance with non-empty values of TimeSeriesArgs
func (args TimeSeriesArgs) QueryParams() url.Values {
	q := make(url.Values)

	if !args.StartDate.IsZero() {
		q.Add("start_date", args.StartDate.Format(dateLayout))
	}

	if !args.EndDate.IsZero() {
		q.Add("end_date", args.StartDate.Format(dateLayout))
	}

	if args.Base != "" {
		q.Add("base", args.Base)
	}

	if len(args.Symbols) > 0 {
		q.Add("symbols", strings.Join(args.Symbols, ","))
	}

	return q
}
