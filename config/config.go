package config

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func InitApp(setup map[string]string)  {
	fmt.Println("Servidor iniciado, presione ctrl+c para finalizar.")
	fmt.Println("puerto: ",setup["port"])
	log.Fatal(http.ListenAndServe(setup["port"], GetRoutes()))//ciclo infinito app
}

func GetConfig()(config map[string]string)  {
	config = map[string]string{
		"port":":8080",
	}
	for i:=0; i<len(os.Args);i++ {
		if os.Args[i]=="-p" {//TODO: fix the leak of the space between -p and the port "-p3000"
			config["port"]=":"+os.Args[i+1]
		}
	}
	return
}

