package camera

import (
	"raspberrypi-car/pca"
	"testing"
)

func TestInitCamera(t *testing.T) {
	c := InitCamera(pca.PCA9685Dev)
	c.SetAngle(10, 20)
}
