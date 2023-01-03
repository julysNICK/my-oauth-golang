package controllers

import "github.com/go-chi/cors"

func (c ServerConfig) Get() {

	cF := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: false,
		})

	c.Routes.Use(cF.Handler)
	c.Routes.HandleFunc("/users", GetUsers).Methods("GET")
	c.Routes.HandleFunc("/register", c.Register).Methods("POST")
	c.Routes.HandleFunc("/login", c.Login).Methods("POST")

}
