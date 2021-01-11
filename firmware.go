package main

import "github.com/naueramant/go-3d-printer/serial"

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

func DetectFirmware(serial *serial.Connection) (firmware Firmware, err error) {
	return FirmwareGeneric, nil
}
