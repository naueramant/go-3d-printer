package firmware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/naueramant/go-3d-printer/serial"
	"github.com/pkg/errors"
)

var (
	ErrFailedToDetectFirmware = errors.New("Failed to detect firmware")
	ErrDetectFirmwareTimeout  = errors.New("Timed outed trying to detect firmware")
	ErrUnknownFirmware        = errors.New("Unknown firmware")
)

func Detect(ctx context.Context, connection *serial.Connection, timeout time.Duration) (firmware printer.FirmwareType, err error) {
	timedCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	firmwareChan := make(chan printer.FirmwareType, 1)
	firmwareErr := make(chan error, 1)

	go func() {
		if err := connection.WriteString("M115\n"); err != nil {
			firmwareErr <- errors.Wrap(err, ErrFailedToDetectFirmware.Error())
		}

		res, err := connection.ReadString()
		if err != nil {
			firmwareErr <- errors.Wrap(err, ErrFailedToDetectFirmware.Error())
		}

		if strings.Contains(res, "Marlin") {
			firmwareChan <- printer.FirmwareTypeMarlin
		}
		if strings.Contains(res, "RepRap") {
			firmwareChan <- printer.FirmwareTypeRepRap
		}
		if strings.Contains(res, "Repetier") {
			firmwareChan <- printer.FirmwareTypeRepetier
		}
		if strings.Contains(res, "Smoothie") {
			firmwareChan <- printer.FirmwareTypeSmoothie
		}
		if strings.Contains(res, "Prusa") {
			firmwareChan <- printer.FirmwareTypePrusa
		}

		firmwareErr <- ErrUnknownFirmware
	}()

	select {
	case <-timedCtx.Done():
		return 0, errors.Wrap(
			errors.New(fmt.Sprintf(
				"Printer did not respond after %.f second(s)",
				timeout.Seconds(),
			)),
			ErrDetectFirmwareTimeout.Error(),
		)
	case f := <-firmwareChan:
		return f, nil
	case err := <-firmwareErr:
		return 0, err
	}
}
