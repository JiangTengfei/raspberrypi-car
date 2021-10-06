package ultrasonic

import (
	"fmt"
	"periph.io/x/host/v3/rpi"
	"raspberrypi-car/drivers"
	"time"
)

type Sensor struct {
	hcsr04 drivers.HCSR04
}

func NewSensor() *Sensor {
	return &Sensor{
		hcsr04: drivers.NewHCSR04(rpi.P1_22, rpi.P1_23),
	}
}

func (s *Sensor) Measure() {
	for true {
		distance := s.hcsr04.MeasureDistance()
		fmt.Printf("Ultrasonic sensor's distance is: %0.2f\n", distance)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
