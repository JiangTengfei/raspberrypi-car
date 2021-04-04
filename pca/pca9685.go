package pca

import (
	"log"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/pca9685"
	"raspberrypi-car/i2c"
)

var PCADev *pca9685.Dev

func init() {
	var err error
	PCADev, err = pca9685.NewI2C(i2c.Bus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err := PCADev.SetPwmFreq(50 * physic.Hertz); err != nil {
		log.Fatal(err)
	}
	if err := PCADev.SetAllPwm(0, 0); err != nil {
		log.Fatal(err)
	}
}