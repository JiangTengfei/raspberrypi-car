package main

import (
	"net/http"
	"periph.io/x/host/v3"
	"raspberrypi-car/controller"
	"raspberrypi-car/wheel"
)

func main() {
	host.Init()

	carWheel := wheel.InitCarWheel()
	webController := controller.NewWebController(carWheel)
	http.ListenAndServe(":8080", webController)
}
