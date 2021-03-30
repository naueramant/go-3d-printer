package main

import (
	"context"
	"github.com/alexsuslov/godotenv"
	"github.com/naueramant/go-3d-printer/pkg/factory"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning load .env err: %v", err)
	}
	ctx := context.TODO()

	port:=os.Getenv("SERIAL_PORT")
	if port ==""{
		port="/dev/ttyUSB0"
	}

	_, err := factory.Connect(ctx, port)
	if err != nil {
		logrus.Fatal(err)
	}
}
