package generic

import (
	"fmt"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

func (p *Printer) SetBedTemperature(temperature int) (err error) {
	if temperature < 0 {
		return errors.New("Temperature can not be negative")
	}

	if _, err := p.SendCommand(fmt.Sprintf("M140 S%d", temperature)); err != nil {
		return errors.Wrap(err, "Failed to set bed temperature")
	}

	return nil
}

func (p *Printer) SetHotendTemperature(hotendIndex, temperature int) (err error) {
	if temperature < 0 {
		return errors.New("Hotend temperature can not be negative")
	}

	if _, err := p.SendCommand(fmt.Sprintf("M104 T%d S%d", hotendIndex, temperature)); err != nil {
		return errors.Wrap(err, "Failed to set hotend temperature")
	}

	return nil
}

func (p *Printer) GetTemperatures() (ch <-chan printer.Temperatures, err error) {
	return nil, errors.New("Not implemented")
}
