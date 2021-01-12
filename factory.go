package main

import (
	"context"

	printer "github.com/naueramant/go-3d-printer/printer"
	"github.com/naueramant/go-3d-printer/printer/generic"
	"github.com/naueramant/go-3d-printer/printer/marlin"
	"github.com/naueramant/go-3d-printer/printer/prusa"
	"github.com/naueramant/go-3d-printer/printer/reprap"
	"github.com/naueramant/go-3d-printer/printer/smoothie"
	"github.com/naueramant/go-3d-printer/serial"
	"github.com/pkg/errors"
)

var (
	ErrInvalidFirmwareEnum  = errors.New("Invalid firmware enum")
	ErrFirmwareNotSupported = errors.New("Firmware not supported")
)

func AutoConnect(ctx context.Context) (p printer.Printer, err error) {
	s, err := serial.NewConnection("/dev/ttyUSB0")
	if err != nil {
		return nil, err
	}

	f, err := DetectFirmware(s)
	if err != nil {
		return nil, err
	}

	return New(ctx, s, f)
}

func New(ctx context.Context, connection *serial.Connection, firmware Firmware) (p printer.Printer, err error) {
	if int(firmware) > len(FirmwareNameMap) {
		return nil, ErrInvalidFirmwareEnum
	}

	switch firmware {
	case FirmwareGeneric:
		return generic.New(ctx, connection), nil
	case FirmwareMarlin:
		return marlin.New(ctx, connection), nil
	case FirmwareRepRap:
		return reprap.New(ctx, connection), nil
	case FirmwareRepetier:
		return reprap.New(ctx, connection), nil
	case FirmwareSmoothie:
		return smoothie.New(ctx, connection), nil
	case FirmwarePrusa:
		return prusa.New(ctx, connection), nil
	}

	return nil, ErrFirmwareNotSupported
}
