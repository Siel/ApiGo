package main

import (
	"github.com/siel/apiTest/config"
)


func main() {
	config.InitApp(config.GetConfig())
}
