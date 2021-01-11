package marlin

import (
	"github.com/naueramant/go-3d-printer/printer/generic"
	"github.com/naueramant/go-3d-printer/serial"
)

// Printer is the Marlin Firmware implementation of the Printer interface.
type Printer struct {
	generic.Printer
	Connection *serial.Connection
}

func New(connection *serial.Connection) (p *Printer) {
	return &Printer{
		Connection: connection,
	}
}
