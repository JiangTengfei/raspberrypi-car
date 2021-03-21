package main

import (
	"fmt"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
	"time"
)

type Direct int

var (
	MotorA1 = rpi.P1_11
	MotorA2 = rpi.P1_12
	MotorB1 = rpi.P1_13
	MotorB2 = rpi.P1_15
	MotorC1 = rpi.P1_16
	MotorC2 = rpi.P1_18
	MotorD1 = rpi.P1_19
	MotorD2 = rpi.P1_21
)

const (
	FORWARD  = Direct(1)
	BACKWARD = Direct(2)
	STOPPED  = Direct(3)
)

func main() {
	host.Init()

	Forward()
	time.Sleep(5 * time.Second)
	Stopped()
	time.Sleep(5 * time.Second)
	Backward()
}

func Forward() {
	Wheel(MotorA1, MotorA2, FORWARD)
	Wheel(MotorB1, MotorB2, FORWARD)
	Wheel(MotorC1, MotorC2, FORWARD)
	Wheel(MotorD1, MotorD2, FORWARD)
}

func Backward() {
	Wheel(MotorA1, MotorA2, BACKWARD)
	Wheel(MotorB1, MotorB2, BACKWARD)
	Wheel(MotorC1, MotorC2, BACKWARD)
	Wheel(MotorD1, MotorD2, BACKWARD)
}

func Stopped() {
	Wheel(MotorA1, MotorA2, STOPPED)
	Wheel(MotorB1, MotorB2, STOPPED)
	Wheel(MotorC1, MotorC2, STOPPED)
	Wheel(MotorD1, MotorD2, STOPPED)
}

func Wheel(motor1, motor2 gpio.PinIO, direct Direct) {
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
