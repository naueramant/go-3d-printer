package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	p, err := AutoConnect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := p.Disconnect(); err != nil {
		logrus.Fatal(err)
	}

	if err := p.Disconnect(); err != nil {
		logrus.Fatal(err)
	}

	files, err := p.ListFiles()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("Path: %s Size: %d\n", f.Path, f.Size)
	}

	p.SetBedTemperature(0)
	p.SetHotendTemperature(0, 0)
}
