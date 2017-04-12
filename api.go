package main

import (
	"github.com/siel/ApiGo/config"
)


func main() {
	config.InitApp(config.GetConfig())
}
