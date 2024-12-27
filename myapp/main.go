package main

import (
	"myapp/controllers"
	"myapp/data"

	"github.com/sde-kiran-sadvilkar/sope"
)

type app struct {
	Core *sope.Sope
	Controllers *controllers.Controllers
	Models data.Models

}

func main() {

	s :=initApp()
	s.Core.CreateServer()
}