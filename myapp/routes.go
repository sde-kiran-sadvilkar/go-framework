package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (core *core) routes() *chi.Mux{
	// middleware must come before any routes


	// add routes here
	core.App.Routes.Get("/",core.Controllers.Home)
	core.App.Routes.Get("/go-page", core.Controllers.GoPage)
	core.App.Routes.Get("/jet-page", core.Controllers.JetPage)
	core.App.Routes.Get("/sessions",core.Controllers.SessionTest)
	

	//static routes 
	fileServer := http.FileServer(http.Dir("./public"))
	core.App.Routes.Handle("/public/*", http.StripPrefix("/public",fileServer))

	return core.App.Routes
}