package githubjobs

import (
	"net/url"
	"strconv"
)

// LatLon is an special struct that stores the latitude and longitude of the job
// A specific latitude. If used, you must also send long and must not send location.
// A specific longitude. If used, you must also send lat and must not send location.
type LatLon struct{ Lat, Lon float64 }

// PositionsArgs Search for jobs by term, location, full time vs part time, or any combination of the three. All parameters are optional
type PositionsArgs struct {
	// A search term, such as "ruby" or "java". This parameter is aliased to search.
	Description string
	// A city name, zip code, or other location search term.
	Location string
	// If present, latitude and longitude coordinates we are requesting
	// a list of jobs in that location
	LatLon *LatLon
	//  If you want to limit results to full time positions set this parameter to 'true'.
	FullTime bool
}

// QueryParams returns an instance of url.Values with the query params passed by the user
func (args PositionsArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.Description != "" {
		q.Add("description", args.Description)
	}

	if args.Location != "" {
		q.Add("location", args.Location)
	}

	if args.LatLon != nil {
		q.Add("lat", strconv.FormatFloat(args.LatLon.Lat, 'f', -1, 64))
		q.Add("lon", strconv.FormatFloat(args.LatLon.Lon, 'f', -1, 64))
	}

	if args.FullTime {
		q.Add("full_time", strconv.FormatBool(args.FullTime))
	}

	return q
}
