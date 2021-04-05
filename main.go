package main

import (
	"log"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
	"raspberrypi-car/car"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	raspPiCar := car.InitCar()
	raspPiCar.Cam.SetAngle(physic.Angle(90), physic.Angle(90))

	raspPiCar.Serv()
}
