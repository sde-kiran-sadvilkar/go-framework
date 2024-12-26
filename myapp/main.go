package main

import (
	"myapp/controllers"

	"github.com/sde-kiran-sadvilkar/sope"
)

type core struct {
	App *sope.Sope
	Controllers *controllers.Controllers

}

func main() {

	s :=initApp()
	s.App.CreateServer()
}