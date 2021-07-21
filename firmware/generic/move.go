package generic

import (
	"fmt"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

var (
	ErrEnableSteppers  = errors.New("Enable steppe motors failed")
	ErrDisableSteppers = errors.New("Disable steppe motors failed")
	ErrMoveAbsolute    = errors.New("Absolute move failed")
	ErrMoveRelative    = errors.New("Relative move failed")
	ErrAutoHome        = errors.New("Auto home failed")
	ErrEmergencyStop   = errors.New("Emergency stop failed")
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

func (p *Printer) MoveAbsolute(x, y, z, rate int, mode printer.MoveMode) (err error) {
	m, err := printer.MoveModeToGCode(mode)
	if err != nil {
		return err
	}

	if _, err := p.SendGCodes([]string{
		"G90",
		fmt.Sprintf("%s X%d Y%d Z%d F%d", m, x, y, z, rate),
	}); err != nil {
		return errors.Wrap(err, ErrMoveAbsolute.Error())
	}

	return nil
}

func (p *Printer) MoveRelative(x, y, z, rate int, mode printer.MoveMode) (err error) {
	m, err := printer.MoveModeToGCode(mode)
	if err != nil {
		return err
	}

	if _, err := p.SendGCodes([]string{
		"G91",
		fmt.Sprintf("%s X%d Y%d Z%d F%d", m, x, y, z, rate),
	}); err != nil {
		return errors.Wrap(err, ErrMoveRelative.Error())
	}

	return nil
}

func (p *Printer) Extrude(amount, rate int) (err error) {
	return errors.New("Not implemented")
}

func (p *Printer) Retract(amount, rate int) (err error) {
	return errors.New("Not implemented")
}

func (p *Printer) AutoHome() (err error) {
	if _, err := p.SendGCode("G28"); err != nil {
		return errors.Wrap(err, ErrAutoHome.Error())
	}

	return nil
}

func (p *Printer) EmergencyStop() (err error) {
	if _, err := p.SendGCode("M112"); err != nil {
		return errors.Wrap(err, ErrEmergencyStop.Error())
	}

	return nil
}

func (p *Printer) GetCurrentPosition() (pos *printer.Position, err error) {
	// TODO: re-evaluate reading of output from printer....

	return nil, errors.New("Not implemented")
}
