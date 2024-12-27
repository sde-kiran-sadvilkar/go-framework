package main

import (
	"log"
	"myapp/controllers"
	"myapp/data"
	"os"

	"github.com/sde-kiran-sadvilkar/sope"
)

func initApp() *core {
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
		App: sop,
	} 

	sop.InfoLog.Println("Debug is set to", sop.Debug)
	sop.ErrorLog.Println("Debug is set to", sop.Debug)

	app:= &core{
		App: sop,
		Controllers: appControllers,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	appControllers.Models = app.Models
	
	return app
}