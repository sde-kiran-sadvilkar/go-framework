package sope

import "net/http"

func (s *Sope) SessionLoad(next http.Handler) http.Handler{
	s.InfoLog.Println("Inside Middleware")
	return s.Session.LoadAndSave(next)
}