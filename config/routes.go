package config

import (
	"github.com/gorilla/mux"
	"github.com/siel/ApiGo/controller"
)

func GetRoutes() *mux.Router{
	var router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.Index)
	//router.HandleFunc("/message", controller.Message).Methods("POST")
	router.HandleFunc("/suma", controller.Suma).Methods("POST")
	//router.HandleFunc("/message/{message}", controller.Message)

	return router
}


