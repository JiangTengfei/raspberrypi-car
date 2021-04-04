package i2c

import (
	"log"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
)

var Bus i2c.BusCloser

func init() {
	var err error
	Bus, err = i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
}
