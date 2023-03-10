package main

import (
	"log"
	"net/http"

	internal "serviceOauth/controllers"

	"github.com/joho/godotenv"
)

func main() {

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
	http.ListenAndServe(":9090", ServerConfig.Routes)

}
