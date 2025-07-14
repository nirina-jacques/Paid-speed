package handler

import (
	"github.com/gorilla/mux"
)

func (cf *ServerConfig) InitRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", cf.HomePage)
	router.HandleFunc("/login", cf.Signin).Methods("POST")
	router.HandleFunc("/register", cf.Register).Methods("POST")
	//router.HandleFunc("/ws", cf.WsHandler)
	cf.Router = router
}
