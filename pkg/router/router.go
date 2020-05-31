package router

import (
	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/handlers/economics"
	"github.com/danvergara/dashboardserver/pkg/handlers/healthcheck"
	"github.com/danvergara/dashboardserver/pkg/handlers/topnews"
	"github.com/danvergara/dashboardserver/pkg/handlers/weather"
	auth "github.com/danvergara/dashboardserver/pkg/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// Response struct
type Response struct {
	Message string `json:"message"`
}

// New returns a pointer to chi Mux Router instance
func New(app *application.Application) *chi.Mux {
	mux := chi.NewRouter()
	jwtMiddleware := auth.NewJWTMiddlerware()

	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(jwtMiddleware.Handler)

	mux.Get("/_healthcheck", healthcheck.Healthcheck(app))
	mux.Get("/v1/top-news", topnews.TopNews(app))
	mux.Get("/v1/current-weather", weather.CurrentWeather(app))
	mux.Get("/v1/weather-forecast", weather.Forecast(app))
	mux.Get("/v1/historical-currency-rates", economics.HistoricalCurrencyRates(app))
	mux.Get("/v1/currency-exchange", economics.CurrencyExchange(app))
	return mux
}
