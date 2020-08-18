package router

import (
	"net/http"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/handlers/economics"
	"github.com/danvergara/dashboardserver/pkg/handlers/healthcheck"
	"github.com/danvergara/dashboardserver/pkg/handlers/jobs"
	"github.com/danvergara/dashboardserver/pkg/handlers/topnews"
	"github.com/danvergara/dashboardserver/pkg/handlers/trendingrepositories"
	"github.com/danvergara/dashboardserver/pkg/handlers/weather"
	auth "github.com/danvergara/dashboardserver/pkg/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	openmiddleware "github.com/go-openapi/runtime/middleware"
)

// Response struct
type Response struct {
	Message string `json:"message"`
}

// New returns a pointer to chi Mux Router instance
func New(app *application.Application) *chi.Mux {
	mux := chi.NewRouter()
	jwtMiddleware := auth.NewJWTMiddlerware()

	ops := openmiddleware.RedocOpts{SpecURL: "./swagger.yaml"}
	sh := openmiddleware.Redoc(ops, nil)

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

	mux.Get("/_healthcheck", healthcheck.Healthcheck(app))
	mux.Handle("/docs", sh)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	mux.Route("/v1", func(r chi.Router) {
		r.Use(jwtMiddleware.Handler)
		r.Get("/top-news", topnews.TopNews(app))
		r.Get("/current-weather", weather.CurrentWeather(app))
		r.Get("/weather-forecast", weather.Forecast(app))
		r.Get("/historical-currency-rates", economics.HistoricalCurrencyRates(app))
		r.Get("/currency-exchange", economics.CurrencyExchange(app))
		r.Get("/repositories", trendingrepositories.Get(app))
		r.Get("/jobs", jobs.Get(app))
	})
	return mux
}
