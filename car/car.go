package car

import (
	"log"
	"net/http"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/pca9685"
	"raspberrypi-car/camera"
	"raspberrypi-car/controller"
	"raspberrypi-car/wheel"
)

var Bus i2c.BusCloser
var PCA9685 *pca9685.Dev

func InitCar() *RaspPiCar {
	var err error
	Bus, err = i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	PCA9685, err = pca9685.NewI2C(Bus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err := PCA9685.SetPwmFreq(50 * physic.Hertz); err != nil {
		log.Fatal(err)
	}
	if err := PCA9685.SetAllPwm(0, 0); err != nil {
		log.Fatal(err)
	}
	cam := camera.InitCamera(PCA9685)
	carWheel := wheel.InitCarWheel()
	webController := controller.NewWebController(carWheel, cam)

	return &RaspPiCar{
		Cam:       cam,
		CarWheels: carWheel,
		WebCtl:    webController,
	}
}

type RaspPiCar struct {
	Cam       *camera.Camera
	CarWheels *wheel.CarWheel
	WebCtl    *controller.WebController
}

func (car *RaspPiCar) Serv() {
	http.ListenAndServe(":8080", car.WebCtl)
}
