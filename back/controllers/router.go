package controllers

import (
	"net/http"
)

func (c ServerConfig) Get() {
	c.Routes.HandleFunc("/users", GetUsers).Methods("GET", http.MethodOptions)
	c.Routes.HandleFunc("/register", c.Register).Methods("POST", http.MethodOptions)
	c.Routes.HandleFunc("/login", c.Login).Methods("POST", http.MethodOptions)

}
