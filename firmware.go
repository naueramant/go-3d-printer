package main

import (
	"strings"

	"github.com/naueramant/go-3d-printer/serial"
	"github.com/pkg/errors"
)

type Firmware uint8

const (
	FirmwareGeneric = iota
	FirmwareMarlin
	FirmwareRepRap
	FirmwareRepetier
	FirmwareSmoothie
	FirmwarePrusa
)

var FirmwareNameMap = map[Firmware]string{
	0: "Generic",
	1: "Marlin",
	2: "RepRap",
	3: "Repetier",
	4: "Smoothie",
	5: "Prusa",
}

var (
	ErrFailedToDetectFirmware = errors.New("Failed to detect firmware")
	ErrUnknownFirmware        = errors.New("Unknown firmware")
)

func DetectFirmware(serial *serial.Connection) (firmware Firmware, err error) {
	if err := serial.WriteString("M115\n"); err != nil {
		return 0, errors.Wrap(err, ErrFailedToDetectFirmware.Error())
	}

	res, err := serial.ReadString()
	if err != nil {
		return 0, errors.Wrap(err, ErrFailedToDetectFirmware.Error())
	}

	if strings.Contains(res, "Marlin") {
		return FirmwareMarlin, nil
	}
	if strings.Contains(res, "RepRap") {
		return FirmwareRepRap, nil
	}
	if strings.Contains(res, "Repetier") {
		return FirmwareRepetier, nil
	}
	if strings.Contains(res, "Smoothie") {
		return FirmwareSmoothie, nil
	}
	if strings.Contains(res, "Prusa") {
		return FirmwarePrusa, nil
	}

	return 0, ErrUnknownFirmware
}
