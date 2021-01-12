package generic

import (
	"context"

	"github.com/naueramant/go-3d-printer/serial"
	"github.com/pkg/errors"
)

var (
	ErrDisconnect = errors.New("Failed to disconnect printer")
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

func (p *Printer) Disconnect() (err error) {
	if err := p.Connection.Port.Flush(); err != nil {
		return errors.Wrap(
			errors.Wrap(err, "Failed to flush printer connection"),
			ErrDisconnect.Error(),
		)
	}

	if err := p.Connection.Disconnect(); err != nil {
		return errors.Wrap(err, ErrDisconnect.Error())
	}

	return nil
}
