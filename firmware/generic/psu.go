package generic

import "github.com/pkg/errors"

func (p *Printer) PowerOn() (err error) {
	if _, err := p.SendCommand("M80"); err != nil {
		return errors.Wrap(err, "Failed to turn power on")
	}

	return nil
}

func (p *Printer) PowerOff() (err error) {
	if _, err := p.SendCommand("M81"); err != nil {
		return errors.Wrap(err, "Failed to turn power off")
	}

	return nil
}
