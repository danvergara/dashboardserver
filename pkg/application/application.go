package application

import (
	"github.com/danvergara/dashboardserver/pkg/config"
)

// Application is the main object of the app
type Application struct {
	Cfg *config.Config
}

// New returns an instance of the Application struct
func New() (*Application, error) {
	cfg := config.NewConfig()

	return &Application{Cfg: cfg}, nil
}
