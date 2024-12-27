package main

import (
	"myapp/controllers"
	"myapp/data"

	"github.com/sde-kiran-sadvilkar/sope"
)

type core struct {
	App *sope.Sope
	Controllers *controllers.Controllers
	Models data.Models

}

func main() {

	s :=initApp()
	s.App.CreateServer()
}