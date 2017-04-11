package config

import (
	"github.com/gorilla/mux"
	"github.com/siel/apiTest/controllers"
)

func GetRoutes() *mux.Router{
	var router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/mensaje/{mensaje}", controllers.Mensaje)

	return router
}

