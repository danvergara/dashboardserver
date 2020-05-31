package main

import (
	"log"

	"github.com/danvergara/dashboardserver/pkg/application"
	"github.com/danvergara/dashboardserver/pkg/exithandler"
	"github.com/danvergara/dashboardserver/pkg/logger"
	"github.com/danvergara/dashboardserver/pkg/router"
	"github.com/danvergara/dashboardserver/pkg/server"
)

func main() {
	app, err := application.New()

	if err != nil {
		log.Fatal(err)
	}

	srv := server.
		New().
		WithAddr(app.Cfg.APIAddr()).
		WithRouter(router.New(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server at %s", app.Cfg.AppPort)
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err)
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			logger.Error.Println(err)
		}
	})

}
