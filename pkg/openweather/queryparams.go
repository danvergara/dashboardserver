package openweather

import (
	"net/url"
	"strconv"
	"strings"
)

// LatLon is a special design to store the geographic coordinates of the city
type LatLon struct{ Lat, Lon int }

// WeatherArgs represents the query params objetcs.
// The url.Values type is built on top of the map[string][]string,
// so it would be cumbersome for a developer to write.
// Furthermore, a developer could spell a query parameter's name wrong, leading to unexpected errors.
// When developers use this struct, they won't need to think about constructing a url.Values,
// or about whether they spelled a query param right.
type WeatherArgs struct {
	// q city name, state code and country code divided by comma, use ISO 3166 country codes.
	// You can specify the parameter not only in English. In this case,
	// the API response should be returned in the same language as the language of requested location name
	// if the location is in our predefined list of more than 200,000 locations.
	// Example of API calss:
	// api.openweathermap.org/data/2.5/forecast?q=London,us&mode=xml
	Q []string
	// You can search weather forecast for 5 days with data every 3 hours by city ID.
	// API responds with exact result. All weather data can be obtained in JSON and XML formats.
	// List of city ID city.list.json.gz can be downloaded here http://bulk.openweathermap.org/sample/
	// Parameters:
	// id city ID
	// Examples of API calls:
	// api.openweathermap.org/data/2.5/forecast?id=524901
	ID int
	// By geographic coordinates
	// Description:
	// You can search weather forecast for 5 days with data every 3 hours by geographic coordinates.
	// All weather data can be obtained in JSON and XML formats.
	// API call:
	// api.openweathermap.org/data/2.5/forecast?lat={lat}&lon={lon}&appid={your api key}
	// Parameters:
	// lat, lon coordinates of the location of your interest
	// Examples of API calls:
	// api.openweathermap.org/data/2.5/forecast?lat=35&lon=139
	LatLon *LatLon
	// By ZIP code
	// Description:
	// Please note if country is not specified then the search works for USA as a default.
	// API call:
	// api.openweathermap.org/data/2.5/forecast?zip={zip code},{country code}&appid={your api key}
	// Parameters:
	// zip zip code
	// Examples of API calls:
	// api.openweathermap.org/data/2.5/forecast?zip=94040,us
	Zip int
	// Units format
	// Description:
	// Standard, metric, and imperial units are available.
	// Parameters:
	// units metric, imperial. When you do not use units parameter, format is Standard by default.
	Units string
}

// QueryParams returns an instance of url.Values with just the not empty values of a WeatherArgs instances
func (args WeatherArgs) QueryParams() url.Values {
	// Creates an empty intance of url.Values
	q := make(url.Values)

	if len(args.Q) > 0 {
		q.Add("q", strings.Join(args.Q, ","))
	}

	if args.ID != 0 {
		q.Add("id", strconv.Itoa(args.ID))
	}

	if args.LatLon != nil {
		q.Add("lat", strconv.Itoa(args.LatLon.Lat))
		q.Add("lon", strconv.Itoa(args.LatLon.Lon))
	}

	if args.Zip != 0 {
		q.Add("zip", strconv.Itoa(args.Zip))
	}

	if args.Units != "" {
		q.Add("units", args.Units)
	}

	return q
}
