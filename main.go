package main

import (
	"fmt"
	"log"
	"net/http"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/pca9685"
	"periph.io/x/host/v3"
	"raspberrypi-car/camera"
	"raspberrypi-car/controller"
	"raspberrypi-car/wheel"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	Bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bus: %+v", Bus)

	PCADev, err := pca9685.NewI2C(Bus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}

	c := camera.InitCamera(PCADev)
	c.SetAngle(10, 20)

	carWheel := wheel.InitCarWheel()
	webController := controller.NewWebController(carWheel)
	http.ListenAndServe(":8080", webController)
}
