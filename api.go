package main

import (
	"github.com/siel/apiTest/controllers"
	"github.com/siel/apiTest/config"
)


func main() {
	controllers.InitApp(config.GetConfig())
}
