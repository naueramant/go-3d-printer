package main

import (
	printer "github.com/naueramant/go-3d-printer/printer/generic"
	com "github.com/naueramant/go-3d-printer/serial"

	"github.com/sirupsen/logrus"
)

func main() {
	c, err := com.NewConnection("/dev/ttyUSB0")
	if err != nil {
		logrus.Fatal(err)
	}

	p := printer.New(c)

	if err := p.SetFanSpeed(0, 255); err != nil {
		logrus.Fatal(err)
	}
}
