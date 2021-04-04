package pca

import (
	"log"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/pca9685"
	"raspberrypi-car/i2c"
)

var PCA9685Dev *pca9685.Dev

func init() {
	var err error
	PCA9685Dev, err = pca9685.NewI2C(i2c.I2cBus, pca9685.I2CAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err := PCA9685Dev.SetPwmFreq(50 * physic.Hertz); err != nil {
		log.Fatal(err)
	}
	if err := PCA9685Dev.SetAllPwm(0, 0); err != nil {
		log.Fatal(err)
	}
}