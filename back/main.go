package main

import (
	"log"
	"net/http"

	internal "serviceOauth/controllers"

	"github.com/joho/godotenv"
)

func main() {
	// headersOk := handlers.AllowedHeaders([]string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Loaded .env file")
	}

	ServerConfig := internal.ServerConfig{}

	ServerConfig.Inicialize()
	ServerConfig.Get()
	http.ListenAndServe(":8080", ServerConfig.Routes)

}
