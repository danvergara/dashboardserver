package trendingos

// TrendingRepository is the response github-trending-api for /repositories
type TrendingRepository struct {
	Author             string        `json:"author"`
	Name               string        `json:"name"`
	Avatar             string        `json:"avatar"`
	URL                string        `json:"url"`
	Description        string        `json:"description"`
	Language           string        `json:"language"`
	LanguageColor      string        `json:"languageColor"`
	Stars              int           `json:"stars"`
	Forks              int           `json:"forks"`
	CurrentPeriodStars int           `json:"currentPeriodStars"`
	BuiltBy            []Contributor `json:"builtBy"`
}

// Contributor is the struct used to store a contributor of a given project
type Contributor struct {
	Href     string `json:"href"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}
