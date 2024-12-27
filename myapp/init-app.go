package main

import (
	"log"
	"myapp/controllers"
	"myapp/data"
	"os"

	"github.com/sde-kiran-sadvilkar/sope"
)

func initApp() *app {
	path,err := os.Getwd()
	if err!=nil {
		log.Fatal(err)
	}

	//init app
	sop := &sope.Sope{}
	err = sop.New(path)

	if err!=nil{
		log.Fatal(err)
	}

	sop.AppName = "myapp"

	appControllers := &controllers.Controllers{
		Core: sop,
	} 

	sop.InfoLog.Println("Debug is set to", sop.Debug)
	sop.ErrorLog.Println("Debug is set to", sop.Debug)

	app:= &app{
		Core: sop,
		Controllers: appControllers,
	}

	app.Core.Routes = app.routes()
	app.Models = data.New(app.Core.DB.Pool)
	appControllers.Models = app.Models
	
	return app
}