package main

import (
	printer "github.com/naueramant/go-3d-printer/printer"
	"github.com/naueramant/go-3d-printer/printer/generic"
	"github.com/naueramant/go-3d-printer/printer/marlin"
	"github.com/naueramant/go-3d-printer/serial"
	"github.com/pkg/errors"
)

var (
	ErrInvalidFirmwareEnum  = errors.New("Invalid firmware enum")
	ErrFirmwareNotSupported = errors.New("Firmware not supported")
)

func AutoConnect() (p printer.Printer, err error) {
	s, err := serial.NewConnection("/dev/ttyUSB0")
	if err != nil {
		return nil, err
	}

	f, err := DetectFirmware(s)
	if err != nil {
		return nil, err
	}

	return New(s, f)
}

func New(connection *serial.Connection, firmware Firmware) (p printer.Printer, err error) {
	if int(firmware) > len(FirmwareNameMap) {
		return nil, ErrInvalidFirmwareEnum
	}

	switch firmware {
	case FirmwareGeneric:
		return generic.New(connection), nil
	case FirmwareMarlin:
		return marlin.New(connection), nil
	case FirmwareRepRap:
		return nil, errors.Wrap(ErrFirmwareNotSupported, FirmwareNameMap[FirmwareRepRap])
	case FirmwareRepetier:
		return nil, errors.Wrap(ErrFirmwareNotSupported, FirmwareNameMap[FirmwareRepetier])
	case FirmwareSmoothie:
		return nil, errors.Wrap(ErrFirmwareNotSupported, FirmwareNameMap[FirmwareSmoothie])
	case FirmwarePrusa:
		return nil, errors.Wrap(ErrFirmwareNotSupported, FirmwareNameMap[FirmwarePrusa])
	}

	return nil, ErrFirmwareNotSupported
}
