package generic

import (
	"github.com/naueramant/go-3d-printer/serial"
)

// Printer can be seen as the most generic implementation of commands for an FDM printer across all firmwares.
type Printer struct {
	Connection *serial.Connection
}

func New(connection *serial.Connection) *Printer {
	return &Printer{
		Connection: connection,
	}
}
