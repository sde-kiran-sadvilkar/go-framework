package main

import (
	"log"
	"myapp/controllers"
	"myapp/data"
	"myapp/middleware"
	"os"

	"github.com/sde-kiran-sadvilkar/sope"
)

func initApp() *app {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init app
	sop := &sope.Sope{}
	err = sop.New(path)

	if err != nil {
		log.Fatal(err)
	}

	sop.AppName = "myapp"

	appControllers := &controllers.Controllers{
		Core: sop,
	}

	appMiddleware := &middleware.Middleware{
		Core: sop,
	}

	sop.InfoLog.Println("Debug is set to", sop.Debug)
	sop.ErrorLog.Println("This is how error log looks", sop.Debug)

	app := &app{
		Core:        sop,
		Controllers: appControllers,
		Middleware:  appMiddleware,
	}

	app.Core.Routes = app.routes()
	app.Models = data.New(app.Core.DB.Pool)
	appControllers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
