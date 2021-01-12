package generic

import (
	"context"

	"github.com/naueramant/go-3d-printer/serial"
)

// Printer can be seen as the most generic implementation of commands for an FDM printer across all firmwares.
type Printer struct {
	Context    context.Context
	Connection *serial.Connection
}

func New(ctx context.Context, connection *serial.Connection) (p *Printer) {
	return &Printer{
		Context:    ctx,
		Connection: connection,
	}
}
