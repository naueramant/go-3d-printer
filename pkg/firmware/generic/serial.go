package generic

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrDisconnect = errors.New("Failed to disconnect printer")
	ErrSendGCode  = errors.New("Failed to send GCode")
)

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

func (p *Printer) SendGCode(gcode string) (result string, err error) {
	if err := p.Connection.WriteString(fmt.Sprintf("%s\n", gcode)); err != nil {
		return "", errors.Wrap(err, ErrSendGCode.Error())
	}

	result, err = p.Connection.ReadString()
	if err != nil {
		return "", errors.Wrap(err, ErrSendGCode.Error())
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

// GetChecksum calculates the checksum
// for a given command.
//
// The checksum is calculated by XOR
// all bytes in the command.
func GetChecksum(cmd string) (checksum int) {
	for _, c := range cmd {
		checksum ^= int(c)
	}

	return
}
