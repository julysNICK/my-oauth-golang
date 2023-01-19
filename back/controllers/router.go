package controllers

import (
	"net/http"

	"github.com/go-chi/cors"
)

func (c ServerConfig) Get() {
	c.Routes.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	c.Routes.HandleFunc("/users", c.GetUsers).Methods("GET", http.MethodOptions)
	c.Routes.HandleFunc("/register", c.Register).Methods("POST", http.MethodOptions)
	c.Routes.HandleFunc("/login", c.Login).Methods("POST", http.MethodOptions)
	c.Routes.HandleFunc("/verifyAuth", c.VerifyAuth).Methods("GET", http.MethodOptions)

}
