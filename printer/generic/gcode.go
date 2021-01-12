package generic

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrSendGCode = errors.New("Failed to send GCODE")
)

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
