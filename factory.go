package main

import (
	"context"
	"time"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/naueramant/go-3d-printer/printer/generic"
	"github.com/naueramant/go-3d-printer/printer/marlin"
	"github.com/naueramant/go-3d-printer/printer/prusa"
	"github.com/naueramant/go-3d-printer/printer/reprap"
	"github.com/naueramant/go-3d-printer/printer/smoothie"
	"github.com/naueramant/go-3d-printer/serial"
	"github.com/pkg/errors"
)

var (
	DetectionTimeout = 1 * time.Second
)

var (
	ErrFirmwareNotSupported = errors.New("Firmware not supported")
)

func AutoConnect(ctx context.Context) (p printer.Printer, err error) {
	s, err := serial.NewConnection("/dev/ttyUSB0")
	if err != nil {
		return nil, err
	}

	f, err := DetectFirmware(ctx, s, DetectionTimeout)
	if err != nil {
		return nil, err
	}

	return New(ctx, s, f)
}

func New(ctx context.Context, connection *serial.Connection, firmware Firmware) (p printer.Printer, err error) {
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
