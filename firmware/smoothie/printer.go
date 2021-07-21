package smoothie

import (
	"context"

	"github.com/naueramant/go-3d-printer/firmware/generic"
	"github.com/naueramant/go-3d-printer/serial"
)

// Printer is the smoothie Firmware implementation of the Printer interface.
type Printer struct {
	// The generic printer implementation which should not be used directly.
	generic.Printer
}

func New(ctx context.Context, connection *serial.Connection) (p *Printer) {
	return &Printer{
		Printer: generic.Printer{
			Context:    ctx,
			Connection: connection,
		},
	}
}
