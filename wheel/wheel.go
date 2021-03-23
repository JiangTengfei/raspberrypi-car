package wheel

import (
	"fmt"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/rpi"
)

type CarWheel struct {
	MotorA1 gpio.PinIO
	MotorA2 gpio.PinIO
	MotorB1 gpio.PinIO
	MotorB2 gpio.PinIO
	MotorC1 gpio.PinIO
	MotorC2 gpio.PinIO
	MotorD1 gpio.PinIO
	MotorD2 gpio.PinIO
}

type Direct int

const (
	FORWARD  Direct = 1
	BACKWARD Direct = 2
	STOPPED  Direct = 3
)

func InitCarWheel() *CarWheel {
	c := &CarWheel{}
	c.MotorA1 = rpi.P1_11
	c.MotorA2 = rpi.P1_12
	c.MotorB1 = rpi.P1_13
	c.MotorB2 = rpi.P1_15
	c.MotorC1 = rpi.P1_16
	c.MotorC2 = rpi.P1_18
	c.MotorD1 = rpi.P1_19
	c.MotorD2 = rpi.P1_21
	_ = c.MotorA1.Out(gpio.Low)
	_ = c.MotorA2.Out(gpio.Low)
	_ = c.MotorB1.Out(gpio.Low)
	_ = c.MotorB2.Out(gpio.Low)
	_ = c.MotorC1.Out(gpio.Low)
	_ = c.MotorC2.Out(gpio.Low)
	_ = c.MotorD1.Out(gpio.Low)
	_ = c.MotorD2.Out(gpio.Low)
	return c
}

func (c *CarWheel) Forward() {
	wheel(c.MotorA1, c.MotorA2, FORWARD)
	wheel(c.MotorB1, c.MotorB2, FORWARD)
	wheel(c.MotorC1, c.MotorC2, FORWARD)
	wheel(c.MotorD1, c.MotorD2, FORWARD)
	fmt.Printf("car move forward.\n")
}

func (c *CarWheel) Backward() {
	wheel(c.MotorA1, c.MotorA2, BACKWARD)
	wheel(c.MotorB1, c.MotorB2, BACKWARD)
	wheel(c.MotorC1, c.MotorC2, BACKWARD)
	wheel(c.MotorD1, c.MotorD2, BACKWARD)
	fmt.Printf("car move backward.\n")
}

func (c *CarWheel) Stopped() {
	wheel(c.MotorA1, c.MotorA2, STOPPED)
	wheel(c.MotorB1, c.MotorB2, STOPPED)
	wheel(c.MotorC1, c.MotorC2, STOPPED)
	wheel(c.MotorD1, c.MotorD2, STOPPED)
	fmt.Printf("car move stopped.\n")
}

func (c *CarWheel) ParallelLeft() {
	wheel(c.MotorA1, c.MotorA2, BACKWARD)
	wheel(c.MotorB1, c.MotorB2, FORWARD)
	wheel(c.MotorC1, c.MotorC2, FORWARD)
	wheel(c.MotorD1, c.MotorD2, BACKWARD)
	fmt.Printf("car move parallel left.\n")
}

func (c *CarWheel) ParallelRight() {
	wheel(c.MotorA1, c.MotorA2, FORWARD)
	wheel(c.MotorB1, c.MotorB2, BACKWARD)
	wheel(c.MotorC1, c.MotorC2, BACKWARD)
	wheel(c.MotorD1, c.MotorD2, FORWARD)
	fmt.Printf("car move parallel right.\n")
}

func (c *CarWheel) TurnLeft() {

}

func (c *CarWheel) TurnRight() {

}

func wheel(motor1, motor2 gpio.PinIO, direct Direct) {
	switch direct {
	case FORWARD:
		_ = motor1.Out(gpio.High)
		_ = motor2.Out(gpio.Low)
		fmt.Printf("motor1: %v - high; motor2: %v - low. \n", motor1.Name(), motor2.Name())
	case BACKWARD:
		_ = motor1.Out(gpio.Low)
		_ = motor2.Out(gpio.High)
		fmt.Printf("motor1: %v - low; motor2: %v - high. \n", motor1.Name(), motor2.Name())
	case STOPPED:
		_ = motor1.Out(gpio.Low)
		_ = motor2.Out(gpio.Low)
		fmt.Printf("motor1: %v - low; motor2: %v - low. \n", motor1.Name(), motor2.Name())
	}
}
