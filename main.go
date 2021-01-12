package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	p, err := AutoConnect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Print("mip")
	time.Sleep(5 * time.Second)

	files, err := p.ListFiles()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("Path: %s Size: %d\n", f.Path, f.Size)
	}
}
