package main

import (
	"devops/app"
	"devops/handler"
)

func main() {
	handler.InitConfig()
	handler.InitDB(handler.ConfigContent.Database)
	app.Start(":8080")
}
