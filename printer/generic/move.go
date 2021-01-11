package generic

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrEnableSteppers  = errors.New("Enable steppe motors failed")
	ErrDisableSteppers = errors.New("Disable steppe motors failed")
	ErrMoveAbsolute    = errors.New("Absolute move failed")
	ErrMoveRelative    = errors.New("Relative move failed")
	ErrAutoHome        = errors.New("Auto home failed")
)

func (p *Printer) EnableSteppers() (err error) {
	if _, err := p.SendGCode("M17"); err != nil {
		return errors.Wrap(err, ErrEnableSteppers.Error())
	}

	return nil
}

func (p *Printer) DisableSteppers() (err error) {
	if _, err := p.SendGCode("M18"); err != nil {
		return errors.Wrap(err, ErrDisableSteppers.Error())
	}

	return nil
}

func (p *Printer) MoveAbsolute(x, y, z, rate int) (err error) {
	if _, err := p.SendGCodes([]string{
		"G90",
		fmt.Sprintf("G0 X%d Y%d Z%d F%d", x, y, z, rate),
	}); err != nil {
		return errors.Wrap(err, ErrMoveAbsolute.Error())
	}

	return nil
}

func (p *Printer) MoveRelative(x, y, z, rate int) (err error) {
	if _, err := p.SendGCodes([]string{
		"G91",
		fmt.Sprintf("G0 X%d Y%d Z%d F%d", x, y, z, rate),
	}); err != nil {
		return errors.Wrap(err, ErrMoveRelative.Error())
	}

	return nil
}

func (p *Printer) Extrude(amount, rate int) (err error) {
	if _, err := p.SendGCodes([]string{
		"G92 E0",
		fmt.Sprintf("G0 E%d F%d", amount, rate),
	}); err != nil {
		return errors.Wrap(err, ErrMoveRelative.Error())
	}

	return nil
}

func (p *Printer) AutoHome() (err error) {
	if _, err := p.SendGCode("G28"); err != nil {
		return errors.Wrap(err, ErrAutoHome.Error())
	}

	return nil
}
