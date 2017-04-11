package controllers

import (
	"net/http"
	"fmt"
	"html"
	"log"
)

func InitApp(config map[string]string)  {
	fmt.Println("Servidor iniciado, presione ctrl+c para finalizar.")
	fmt.Println("puerto: ",config["port"])
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Se ha recibido una solicitud de:",r.RemoteAddr)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(config["port"], nil))

}
