package controllers

import (
	"myapp/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/sde-kiran-sadvilkar/sope"
)

type Controllers struct {
	Core *sope.Sope
	Models data.Models
}

func (c *Controllers) Home (w http.ResponseWriter, r *http.Request){
	err:= c.Core.Render.Page(w,r, "home", nil, nil)

	if err != nil {
		c.Core.ErrorLog.Println("Error Rendering", err)
	}
}

func (c *Controllers) GoPage (w http.ResponseWriter, r *http.Request){
	err:= c.Core.Render.RenderGoPage(w,r, "home", nil)

	if err != nil {
		c.Core.ErrorLog.Println("Error Rendering", err)
	}
}

func (c *Controllers) JetPage (w http.ResponseWriter, r *http.Request){
	err:= c.Core.Render.RenderJetPage(w,r, "jet-template", nil, nil)

	if err != nil {
		c.Core.ErrorLog.Println("Error Rendering", err)
	}
}

func (c *Controllers) SessionTest (w http.ResponseWriter, r *http.Request){
	
	testData := "bar"

	c.Core.Session.Put(r.Context(),"foo", testData)

	recievedData := c.Core.Session.GetString(r.Context(),"foo")

	vars:= make(jet.VarMap)
	vars.Set("foo",recievedData)
	
	err:= c.Core.Render.RenderJetPage(w,r, "sessions", vars, nil)

	if err != nil {
		c.Core.ErrorLog.Println("Error Rendering", err)
	}
}