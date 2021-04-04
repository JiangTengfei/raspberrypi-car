package main

import (
	"raspberrypi-car/camera"
	"raspberrypi-car/controller"
	"raspberrypi-car/wheel"
)

type RaspPiCar struct {
	Cam *camera.Camera
	CarWheels *wheel.CarWheel
	WebCtl    *controller.WebController
}
