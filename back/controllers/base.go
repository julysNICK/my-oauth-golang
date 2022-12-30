package controllers

import (
	"os"

	"serviceOauth/internal"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ServerConfig struct {
	DB     *gorm.DB
	Routes *mux.Router
}

func (ServerConfig *ServerConfig) Inicialize() (*gorm.DB, *mux.Router) {
	ServerConfig.Routes = mux.NewRouter()

	ServerConfig.DB = internal.GetConnection(
		os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_HOST"), os.Getenv("DB_NAME"),
	)
	return ServerConfig.DB, ServerConfig.Routes
}
