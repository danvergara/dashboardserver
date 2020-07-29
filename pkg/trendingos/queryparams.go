package trendingos

import "net/url"

// TrendingRepositoryArgs is object used to map the query params passed by the user
type TrendingRepositoryArgs struct {
	// language: optional, list trending repositories of certain programming languages,
	// possible values are listed here:
	// https://github.com/huchenme/github-trending-api/blob/master/src/languages.json
	Language string
	// since: optional, default to daily, possible values: daily, weekly and monthly
	Since string
	// spoken_language_code: optional, list trending repositories of certain spoken languages
	// (e.g English, Chinese), possible values are listed here:
	// https://github.com/huchenme/github-trending-api/blob/master/src/spoken-languages.json
	SpokenLanguageCode string
}

// QueryParams returns an instance of url.Values with the query params passed by the user
func (args TrendingRepositoryArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.Language != "" {
		q.Add("language", args.Language)
	}

	if args.Since != "" {
		q.Add("since", args.Since)
	}

	if args.SpokenLanguageCode != "" {
		q.Add("spoken_language_code", args.SpokenLanguageCode)
	}

	return q
}
