package main

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *app) routes() *chi.Mux{
	// middleware must come before any routes


	// add routes here
	app.Core.Routes.Get("/",app.Controllers.Home)
	app.Core.Routes.Get("/go-page", app.Controllers.GoPage)
	app.Core.Routes.Get("/jet-page", app.Controllers.JetPage)
	app.Core.Routes.Get("/sessions",app.Controllers.SessionTest)
	
	app.Core.Routes.Get("/users/login", app.Controllers.UserLogin)
	app.Core.Routes.Post("/users/login", app.Controllers.PostUserLogin)
	app.Core.Routes.Get("/users/logout", app.Controllers.Logout)


	app.Core.Routes.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Kiran",
			LastName: "Sadvilkar",
			Email: "test@example.com",
			Active: 1,
			Password: "password",
		}

		id, err := app.Models.Users.Insert(u)

		if err != nil {
			app.Core.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d: %s", id, u.FirstName)
	})


	app.Core.Routes.Get("/get-all-users", func(w http.ResponseWriter, r *http.Request){
		users, err := app.Models.Users.GetAll()
		if err != nil {
			app.Core.ErrorLog.Println(err)
			return
		}
		for _, x := range users {
			fmt.Fprint(w, x.LastName)
		}
	})

	app.Core.Routes.Get("/get-user/{id}", func(w http.ResponseWriter, r *http.Request){
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		u, err := app.Models.Users.Get(id)
		if err != nil {
			app.Core.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%s %s %s", u.FirstName, u.LastName, u.Email)
	})

	app.Core.Routes.Get("/update-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := app.Models.Users.Get(id)
		if err != nil {
			app.Core.ErrorLog.Println(err)
			return
		}

		u.LastName = app.Core.RandomString(10)
		err = u.Update(*u)
		if err != nil {
			app.Core.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "updated last name to %s", u.LastName)

	})
	

	//static routes 
	fileServer := http.FileServer(http.Dir("./public"))
	app.Core.Routes.Handle("/public/*", http.StripPrefix("/public",fileServer))

	return app.Core.Routes
}