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

	if err := p.MoveRelative(0, 10, 0, 2000); err != nil {
		logrus.Fatal(err)
	}
}
