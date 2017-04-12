package config

import (
	"os"
	"fmt"
	"log"
	"net/http"
)


func InitApp(setup map[string]string)  {
	fmt.Println("Service started at port:",setup["port"], "press ctrl+c to finish.")
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

