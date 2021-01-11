package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	p, err := AutoConnect()
	if err != nil {
		logrus.Fatal(err)
	}

	if err := p.SetFanSpeed(0, 0); err != nil {
		logrus.Fatal(err)
	}
}
