package main

import (
	"log"
	"net/http"
	"periph.io/x/host/v3"
	"raspberrypi-car/camera"
	"raspberrypi-car/controller"
	"raspberrypi-car/pca"
	"raspberrypi-car/wheel"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	c := camera.InitCamera(pca.PCA9685Dev)
	c.SetAngle(10, 20)

	carWheel := wheel.InitCarWheel()
	webController := controller.NewWebController(carWheel)
	http.ListenAndServe(":8080", webController)
}
