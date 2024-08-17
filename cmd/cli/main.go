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

	logrus.Infof("Machine: %s", info.MachineType)
	logrus.Infof("Firmware Name: %s", info.FirmwareName)
	logrus.Infof("Source Code URL: %s", info.SourceCodeURL)
	logrus.Infof("Protocol: %s", info.ProtocolVersion)
	logrus.Infof("Extruders: %d", info.ExtruderCount)
	logrus.Infof("UUID:     %s", info.UUID)
}
