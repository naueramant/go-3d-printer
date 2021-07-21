package main

import (
	"context"

	factory "github.com/naueramant/go-3d-printer"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.TODO()

	p, err := factory.AutoConnect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	p.AutoHome()
}
