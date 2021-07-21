package generic

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func (p *Printer) Disconnect() (err error) {
	if err := p.Connection.Port.Flush(); err != nil {
		return errors.Wrap(err, "Failed to flush printer connection")
	}

	if err := p.Connection.Disconnect(); err != nil {
		return errors.Wrap(err, "Failed to disconnect printer")
	}

	return nil
}

func (p *Printer) SendGCode(gcode string) (result string, err error) {
	if err := p.Connection.WriteString(fmt.Sprintf("%s\n", gcode)); err != nil {
		return "", errors.Wrap(err, "Failed to write GCode to printer")
	}

	result, err = p.Connection.ReadString()
	if err != nil {
		return "", errors.Wrap(err, "Failed to read result from printer")
	}

	result = strings.ReplaceAll(result, "\r", "")

	if strings.Contains(result, "echo:Unknown command:") {
		return "", errors.Errorf("Unknown command \"%s\"", gcode)
	}

	return result, nil
}

func (p *Printer) SendGCodes(gcodes []string) (result []string, err error) {
	result = make([]string, len(gcodes))

	for i, c := range gcodes {
		r, err := p.SendGCode(c)
		if err != nil {
			return result, err
		}

		result[i] = r
	}

	return result, nil
}

func getXORChecksum(cmd string) (checksum int) {
	for _, c := range cmd {
		checksum ^= int(c)
	}

	return
}
