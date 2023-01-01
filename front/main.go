package main

import (
	"encoding/gob"
	"flag"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"os"

	"time"
)

const version = "1.0.0"
const cssVersion = "1"


 type Config struct {
 	Port int
 	Env string
 	Api string
 	Db struct {
 		Ds string
 	}
 	SecretKey string
 	Frontend string
 }

 type Application struct {
 	Config Config
 	InfoLog *log.Logger
 	ErrorLog *log.Logger
 	TemplateCache map[string]*template.Template
 	Version string
 }

func (app *Application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.Config.Port),
		 Handler: app.routes(),
		ErrorLog: app.ErrorLog,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.InfoLog.Printf("Starting server on port %d", app.Config.Port)
	return srv.ListenAndServe()
}
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}


func main() {
	gob.Register(User{})
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "HTTP network port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.Api, "api", "http://localhost:3000", "API base URL")
	flag.StringVar(&cfg.SecretKey, "secretkey", "bRWmrwNUTqNUuzckjxsFlHZjxHkjrzKP", "Secret key")
	flag.StringVar(&cfg.Frontend, "frontend", "http://localhost:3000", "Frontend base URL")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	tc := make(map[string]*template.Template)

	app := &Application{
		Config: cfg,
		InfoLog: infoLog,
		ErrorLog: errorLog,
		TemplateCache: tc,
		Version: version,
	}

	err := app.serve()

	if err != nil {
		errorLog.Fatal(err)
	}

}