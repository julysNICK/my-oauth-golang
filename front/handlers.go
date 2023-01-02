package main

import "net/http"

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.ErrorLog.Println(err)
	}
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "register", &templateData{}); err != nil {
		app.ErrorLog.Println(err)
	}
}