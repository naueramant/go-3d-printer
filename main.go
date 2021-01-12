package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	p, err := AutoConnect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := p.PowerOn(); err != nil {
		logrus.Fatal(err)
	}
}
