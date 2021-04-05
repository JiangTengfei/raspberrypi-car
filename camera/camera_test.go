package camera

import (
	"raspberrypi-car/car"
	"testing"
)

func TestInitCamera(t *testing.T) {
	c := InitCamera(car.PCA9685)
	c.SetAngle(10, 20)
}
