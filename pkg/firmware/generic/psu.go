package generic

import "github.com/pkg/errors"

var (
	ErrPowerOn  = errors.New("Failed to turn power on")
	ErrPowerOff = errors.New("Failed to turn power off")
)

func (p *Printer) PowerOn() (err error) {
	if _, err := p.SendGCode("M80"); err != nil {
		return errors.Wrap(err, ErrPowerOn.Error())
	}

	return nil
}

func (p *Printer) PowerOff() (err error) {
	if _, err := p.SendGCode("M81"); err != nil {
		return errors.Wrap(err, ErrPowerOff.Error())
	}

	return nil
}
