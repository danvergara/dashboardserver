package router

import (
	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/handlers/economics"
	"github.com/danvergara/dashboardserver/pkg/handlers/healthcheck"
	"github.com/danvergara/dashboardserver/pkg/handlers/topnews"
	"github.com/danvergara/dashboardserver/pkg/handlers/weather"

	"github.com/go-chi/chi"
)

// New returns a pointer to chi Mux Router instance
func New(app *application.Application) *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/_healthcheck", healthcheck.Healthcheck(app))
	mux.Get("/v1/top-news", topnews.TopNews(app))
	mux.Get("/v1/current-weather", weather.CurrentWeather(app))
	mux.Get("/v1/weather-forecast", weather.Forecast(app))
	mux.Get("/v1/historical-currency-rates", economics.HistoricalCurrencyRates(app))
	mux.Get("/v1/currency-exchange", economics.CurrencyExchange(app))
	return mux
}
