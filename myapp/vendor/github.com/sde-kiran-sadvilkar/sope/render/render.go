package render

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews *jet.Set
}

type TemplateData struct {
	isAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (s *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface {}) error{
	switch strings.ToLower(s.Renderer){
	case "go":
		return s.RenderGoPage(w,r,view,data)

	case "jet":
		return s.RenderJetPage(w,r,view,variables,data)

	default :

	}

	return errors.New("no rendering engine specified.Please set rendering engine in .env file")
}

func (s *Render) RenderGoPage( w http.ResponseWriter, r *http.Request, view string,data interface{}) error{
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl",s.RootPath,view))

	if err!=nil {
		return err
	}

	td := &TemplateData{}

	if data!= nil{
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)

	if err != nil {
		return err
	}

	return nil
}

func (s *Render) RenderJetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error{
	var vars jet.VarMap

	if variables == nil{
		vars = make(jet.VarMap)
	}else{
		vars= variables.(jet.VarMap)
	}


	td := &TemplateData{}

	if data != nil{
		td = data.(*TemplateData)
	}

	t,err:= s.JetViews.GetTemplate(fmt.Sprintf("%s.jet",templateName))

	if err!= nil{
		log.Println(err)
		return err
	}

	err = t.Execute(w, vars,td)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}