package camera

import (
	"raspberrypi-car/pca"
	"testing"
)

func TestInitCamera(t *testing.T) {
	c := InitCamera(pca.PCADev)
	c.SetAngle(10, 20)
}
