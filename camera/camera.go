package camera

import (
	"fmt"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/pca9685"
)

var baseChan chan physic.Angle
var higherChan chan physic.Angle

type Camera struct {
	BaseServo   *pca9685.Servo
	HigherServo *pca9685.Servo
}

func InitCamera(p *pca9685.Dev) *Camera {
	fmt.Printf("call InitCamera method. dev: %+v \n", p)

	baseChan = make(chan physic.Angle)
	higherChan = make(chan physic.Angle)

	servos := pca9685.NewServoGroup(p, 50, 650, 10, 180)
	baseServo := servos.GetServo(0)
	higherServo := servos.GetServo(1)

	baseServo.SetMinMaxAngle(1, 180)
	higherServo.SetMinMaxAngle(1, 180)

	c := &Camera{
		BaseServo:   baseServo,
		HigherServo: higherServo,
	}

	go func(cam *Camera) {

		for {
			select {
			case a := <-baseChan:
				fmt.Printf("receive sig from baseChan: %d\n", int64(a))
				if err := c.BaseServo.SetAngle(a); err != nil {
					fmt.Printf("base servo SetAngle method return error: %+v\n", err)
				}
			case b := <-higherChan:
				fmt.Printf("receive sig from higherChan: %d\n", int64(b))
				if err := c.HigherServo.SetAngle(b); err != nil {
					fmt.Printf("higher servo SetAngle method return error: %+v\n", err)
				}
			}
		}

	}(c)

	// 初始化位置
	c.SetAngle(physic.Angle(45), physic.Angle(45))

	return c
}

func (c *Camera) SetAngle(horizontal, vertical physic.Angle) {
	baseChan <- horizontal
	higherChan <- vertical
}
