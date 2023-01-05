package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serviceOauth/dto"
	"serviceOauth/models"
	"serviceOauth/services"
)

var users = []dto.User{
	{
		Id:       1,
		Email:    "test@test.com",
		Password: "test",
		Token:    "test",
	},
	{
		Id:       2,
		Email:    "test2@test.com",
		Password: "test",
		Token:    "test",
	},
	{
		Id:       3,
		Email:    "test3@test.com",
		Password: "test",
		Token:    "test",
	},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (ServerConfig *ServerConfig) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (ServerConfig *ServerConfig) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	user, err := services.RegisterUser(ServerConfig.DB, user)
	fmt.Println(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (ServerConfig *ServerConfig) Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	user, err := services.LoginUser(ServerConfig.DB, user.Email, user.Password)

	if err != nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
