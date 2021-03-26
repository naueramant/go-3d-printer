package main

import (
	"context"

	"github.com/naueramant/go-3d-printer/pkg/factory"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.TODO()

	_, err := factory.Connect(ctx, "/dev/ttyUSB0")
	if err != nil {
		logrus.Fatal(err)
	}
}
