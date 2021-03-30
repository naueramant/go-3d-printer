package reprap

import (
	"context"
	"github.com/naueramant/go-3d-printer/pkg/printer"
	"strings"

	"github.com/naueramant/go-3d-printer/pkg/firmware/generic"
	"github.com/naueramant/go-3d-printer/pkg/serial"
)

// Printer is the reprap Firmware implementation of the Printer interface.
type Printer struct {
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


var keys = []string{
	"FIRMWARE_NAME:",
	"FIRMWARE_VERSION:",
	"ELECTRONICS:",
	"FIRMWARE_DATE:",
}


func (p *Printer) GetFirmwareInformation() (info *printer.FirmwareInformation, err error) {
	info=&printer.FirmwareInformation{}
	err = p.Connection.WriteString("M115\n")
	if err!= nil{
		return
	}
	res, err := p.Connection.ReadString()
	if err!= nil{
		return
	}
	var splitPoints  []int
	for _, key := range keys{
		splitPoints = append(splitPoints, strings.Index(res, key))
	}
	// todo: sort splitPoints
	info.Props=map[string]string{}
	var i=0
	for _, n :=range splitPoints{
		if i==n{
			continue
		}
		s:=res[i: n]
		enteres:=strings.Split(s,":")
		info.Props[enteres[0]]=enteres[1]
		i=n
	}

	enteres:=strings.Split(res[i: len(res)],":")
	info.Props[enteres[0]]=strings.TrimRight(enteres[1],"\nok\n")

	return
}