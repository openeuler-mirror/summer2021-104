package main

import (
	"devops/app"
	"devops/handler"
)

func init() {
	handler.InitConfig()
	handler.InitDB(handler.ConfigContent.Database)
}

func Run() {
	app.Start("127.0.0.1:8080")
}

func main() {
	go handler.MonitorCircle()
	Run()
}
