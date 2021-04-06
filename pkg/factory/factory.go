package factory

import (
	"context"
	"time"

	"github.com/naueramant/go-3d-printer/pkg/firmware/generic"
	"github.com/naueramant/go-3d-printer/pkg/firmware/marlin"
	"github.com/naueramant/go-3d-printer/pkg/firmware/prusa"
	"github.com/naueramant/go-3d-printer/pkg/firmware/reprap"
	"github.com/naueramant/go-3d-printer/pkg/firmware/smoothie"
	"github.com/naueramant/go-3d-printer/pkg/printer"
	"github.com/naueramant/go-3d-printer/pkg/serial"
	"github.com/pkg/errors"
)

const (
	DetectionTimeout = 2 * time.Second
	DefaultBaudRate  = 115200
)

var (
	ErrFirmwareNotSupported = errors.New("Firmware not supported")
	ErrNoPrintersFound      = errors.New("No printers found")
)

func AutoConnect(ctx context.Context) (p printer.Printer, err error) {
	devices, err := serial.GetSerialDevices()
	if err != nil {
		return nil, err
	}

	for _, d := range devices {
		p, err := Connect(ctx, d, DefaultBaudRate)
		if err == nil {
			return p, nil
		}
	}

	return nil, ErrNoPrintersFound
}

func Connect(ctx context.Context, device string, baudrate int) (p printer.Printer, err error) {
	s, err := serial.NewConnection(device, baudrate)
	if err != nil {
		return nil, err
	}

	f, err := DetectFirmware(ctx, s, DetectionTimeout)
	if err != nil {
		return nil, err
	}

	return newPrinter(ctx, s, f)
}

func newPrinter(ctx context.Context, connection *serial.Connection, firmware printer.FirmwareType) (p printer.Printer, err error) {
	switch firmware {
	case printer.FirmwareTypeGeneric:
		return generic.New(ctx, connection), nil
	case printer.FirmwareTypeMarlin:
		return marlin.New(ctx, connection), nil
	case printer.FirmwareTypeRepRap:
		return reprap.New(ctx, connection), nil
	case printer.FirmwareTypeRepetier:
		return reprap.New(ctx, connection), nil
	case printer.FirmwareTypeSmoothie:
		return smoothie.New(ctx, connection), nil
	case printer.FirmwareTypePrusa:
		return prusa.New(ctx, connection), nil
	}

	return nil, ErrFirmwareNotSupported
}
