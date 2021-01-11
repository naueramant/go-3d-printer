package generic

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrSetFanSpeedValue = errors.New("Fan speed needs to be in the range 0-255")
	ErrSetFanSpeed      = errors.New("Set fan speed failed")
)

func (p *Printer) SetFanSpeed(fanIndex, speed int) (err error) {
	if speed < 0 || speed > 255 {
		return ErrSetFanSpeedValue
	}

	if _, err := p.SendGCode(fmt.Sprintf("M116 S%d", speed)); err != nil {
		return errors.Wrap(err, ErrSetFanSpeed.Error())
	}

	return nil
}
