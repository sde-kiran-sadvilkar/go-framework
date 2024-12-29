package main

import (
	"myapp/controllers"
	"myapp/data"
	"myapp/middleware"

	"github.com/sde-kiran-sadvilkar/sope"
)

type app struct {
	Core        *sope.Sope
	Controllers *controllers.Controllers
	Models      data.Models
	Middleware  *middleware.Middleware
}

func main() {

	s := initApp()
	s.Core.CreateServer()
}
