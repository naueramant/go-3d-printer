package generic

import (
	"fmt"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

func (p *Printer) EnableStepperMotors() (err error) {
	if _, err := p.SendGCode("M17"); err != nil {
		return errors.Wrap(err, "Failed to enable stepper motors")
	}

	return nil
}

func (p *Printer) DisableStepperMotors() (err error) {
	if _, err := p.SendGCode("M18"); err != nil {
		return errors.Wrap(err, "Failed to disable stepper motors")
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
		return errors.Wrap(err, "Failed to move print head absolute")
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
		return errors.Wrap(err, "Failed to move print head relative")
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
		return errors.Wrap(err, "Failed to auto home printer")
	}

	return nil
}

func (p *Printer) EmergencyStop() (err error) {
	if _, err := p.SendGCode("M112"); err != nil {
		return errors.Wrap(err, "Emergency stop failed")
	}

	return nil
}

func (p *Printer) GetPosition() (pos *printer.Position, err error) {
	// TODO: re-evaluate reading of output from printer....

	return nil, errors.New("Not implemented")
}
