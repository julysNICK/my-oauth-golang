package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
 	mux.Get("/", app.Home)
	mux.Get("/register", app.Register)
	return mux
}