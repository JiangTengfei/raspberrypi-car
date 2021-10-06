package drivers

import (
	"periph.io/x/conn/v3/gpio"
	"time"
)

const HardStop = 1000000

type HCSR04 struct {
	EchoPin gpio.PinIO
	PingPin gpio.PinIO
}

func NewHCSR04(ping gpio.PinIO, echo gpio.PinIO) (result HCSR04) {
	result.PingPin = ping
	result.EchoPin = echo
	return
}

func (hcsr *HCSR04) MeasureDistance() float32 {

	_ = hcsr.EchoPin.Out(gpio.Low)
	_ = hcsr.PingPin.Out(gpio.Low)

	hcsr.EchoPin.Pull()

	strobeZero := 0
	strobeOne := 0

	// strobe
	delayUs(200)
	_ = hcsr.PingPin.Out(gpio.High)
	delayUs(15)
	_ = hcsr.PingPin.Out(gpio.Low)

	// wait until strobe back
	for strobeZero = 0; strobeZero < HardStop && hcsr.EchoPin.Read() != gpio.High; strobeZero++ {
	}
	start := time.Now()
	for strobeOne = 0; strobeOne < HardStop && hcsr.EchoPin.Read() != gpio.Low; strobeOne++ {
		delayUs(1)
	}
	end := time.Now()

	return float32(end.UnixNano()-start.UnixNano()) / (58.0 * 1000)
}

func delayUs(ms int) {
	time.Sleep(time.Duration(ms) * time.Microsecond)
}
