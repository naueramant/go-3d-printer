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

	info, err := p.GetFirmwareInformation()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Machine:  %s", info.MachineType)
	logrus.Infof("Firmware: %s", info.FirmwareName)
	logrus.Infof("UUID:     %s", info.UUID)
}
