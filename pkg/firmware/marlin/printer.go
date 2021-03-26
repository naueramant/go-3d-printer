package marlin

import (
	"context"

	"github.com/naueramant/go-3d-printer/pkg/firmware/generic"
	"github.com/naueramant/go-3d-printer/pkg/serial"
)

// Printer is the Marlin Firmware implementation of the Printer interface.
type Printer struct {
	generic.Printer
	Context    context.Context
	Connection *serial.Connection
}

func New(ctx context.Context, connection *serial.Connection) (p *Printer) {
	return &Printer{
		Printer: generic.Printer{
			Context:    ctx,
			Connection: connection,
		},
		Context:    ctx,
		Connection: connection,
	}
}
