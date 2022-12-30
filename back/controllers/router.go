package controllers

func (c ServerConfig) Get() {
	c.Routes.HandleFunc("/users", GetUsers).Methods("GET")
	c.Routes.HandleFunc("/register", c.Register).Methods("POST")
	c.Routes.HandleFunc("/login", c.Login).Methods("POST")
}
