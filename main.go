package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	p, err := AutoConnect()
	if err != nil {
		logrus.Fatal(err)
	}

	if err := p.MoveRelative(0, 10, 0, 2000); err != nil {
		logrus.Fatal(err)
	}
}
