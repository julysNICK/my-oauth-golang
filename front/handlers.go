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

func (app *Application) PostLoginPage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	app.InfoLog.Println(email, password)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) Auth(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "auth", &templateData{}); err != nil {
		app.ErrorLog.Println(err)
	}

}
func (app *Application) AuthTemp(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "auth-temp", &templateData{}); err != nil {
		app.ErrorLog.Println(err)
	}

}

func (app *Application) PostAuthPage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	token := r.Form.Get("token")
	idUser := r.Form.Get("id")

	app.InfoLog.Println(email, password, token, idUser)

	http.Redirect(w, r, "/auth", http.StatusSeeOther)
}
