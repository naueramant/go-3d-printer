package generic

import (
	"fmt"

	"github.com/pkg/errors"
)

func (p *Printer) SetFanSpeed(fanIndex, speed int) (err error) {
	if speed < 0 || speed > 255 {
		return errors.New("Print value must be in the range 0-255")
	}

	if _, err := p.SendGCode(fmt.Sprintf("M106 P%d S%d", fanIndex, speed)); err != nil {
		return errors.Wrap(err, "Failed to set fan speed")
	}

	return nil
}
