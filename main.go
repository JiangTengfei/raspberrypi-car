package main

import (
	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
	"log"
	"periph.io/x/host/v3"
	"time"
)

func main() {
	host.Init()

	//carWheel := wheel.InitCarWheel()
	//webController := controller.NewWebController(carWheel)
	//http.ListenAndServe(":8080", webController)

	// Create new connection to i2c-bus on 1 line with address 0x40.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.New(pca9685.Address, 1)
	if err != nil {
		log.Fatal(err)
	}

	pca0, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Sets a single PWM channel 0
	pca0.SetChannel(0, 0, 130)

	// Servo on channel 0
	servo0 := pca0.ServoNew(0, nil)

	// Angle in degrees. Must be in the range `0` to `Range`
	for i := 0; i < 130; i++ {
		servo0.Angle(i)
		time.Sleep(10 * time.Millisecond)
	}

	// Fraction as pulse width expressed between 0.0 `MinPulse` and 1.0 `MaxPulse`
	servo0.Fraction(0.5)
}
