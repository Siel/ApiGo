package controllers

import (
	"net/http"
	"fmt"
	//"html"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//http.FileServer(http.Dir("./static"))
	fmt.Fprintln(w, "Welcome controller!")
}

func Mensaje(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL, "requested. From:",r.RemoteAddr)
	fmt.Println("Con parámetros: ",mux.Vars(r))

	mensaje := mux.Vars(r)["mensaje"]
	fmt.Fprintln(w, "Mensaje:", mensaje)
}
