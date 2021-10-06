package ultrasonic

import (
	"fmt"
	"periph.io/x/host/v3/rpi"
	"raspberrypi-car/controller"
	"raspberrypi-car/drivers"
	"time"
)

type Sensor struct {
	hcsr04        drivers.HCSR04
	webController *controller.WebController
}

func NewSensor(w *controller.WebController) *Sensor {
	return &Sensor{
		hcsr04:        drivers.NewHCSR04(rpi.P1_22, rpi.P1_23),
		webController: w,
	}
}

func (s *Sensor) Measure() {
	for true {
		distance := s.hcsr04.MeasureDistance()
		if distance < 10 {
			s.webController.CarWheel.Stopped()
		}
		fmt.Printf("Ultrasonic sensor's distance is: %0.2f\n", distance)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
