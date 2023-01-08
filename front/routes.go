package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", app.Home)
	mux.Get("/register", app.Register)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/auth", app.Auth)
	mux.Get("/auth-temp", app.AuthTemp)
	mux.Post("/auth", app.PostAuthPage)
	return mux
}
