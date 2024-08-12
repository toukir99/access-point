package web

import (
	"access-point/web/handlers"
	"access-point/web/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /users/signup", 
		manager.With(
			http.HandlerFunc(handlers.SignUp),
			
		),
	)
	mux.Handle(
		"POST /users/login", 
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}