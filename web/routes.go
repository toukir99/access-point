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
			http.HandlerFunc(handlers.SignUpUser),
			
		),
	)
	mux.Handle(
		"POST /users/verify-otp", 
		manager.With(
			http.HandlerFunc(handlers.VerifyOTP),
		),
	)
	mux.Handle(
		"POST /users/signin", 
		manager.With(
			http.HandlerFunc(handlers.SignInUser),
		),
	)
	mux.Handle(
		"GET /users", 
		manager.With(
			http.HandlerFunc(handlers.GetAllUsers),
		),
	)
	mux.Handle(
		"GET /users/", 
		manager.With(
			http.HandlerFunc(handlers.GetUserById),
			middlewares.AuthMiddleware,
		),
	)
	mux.Handle(
		"PUT /users", 
		manager.With(
			http.HandlerFunc(handlers.UpdateUser),
			middlewares.AuthMiddleware,
		),
	)
	mux.Handle(
		"DELETE /users", 
		manager.With(
			http.HandlerFunc(handlers.DeleteUser),
			middlewares.AuthMiddleware,
		),
	)
}