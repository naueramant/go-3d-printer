package generic

import (
	"fmt"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

var (
	ErrTemperatureValue     = errors.New("Invalid temperature value")
	ErrSetBedTemperature    = errors.New("Failed to set bed temperature")
	ErrSetHotendTemperature = errors.New("Failed to set hotend temperature")
)

func (p *Printer) SetBedTemperature(temperature int) (err error) {
	if temperature < 0 {
		return errors.Wrap(
			errors.New("Temperature can not be negative"),
			ErrTemperatureValue.Error(),
		)
	}

	if _, err := p.SendGCode(fmt.Sprintf("M140 S%d", temperature)); err != nil {
		return errors.Wrap(err, ErrSetBedTemperature.Error())
	}

	return nil
}

func (p *Printer) SetHotendTemperature(hotendIndex, temperature int) (err error) {
	if temperature < 0 {
		return errors.Wrap(
			errors.New("Temperature can not be negative"),
			ErrTemperatureValue.Error(),
		)
	}

	if _, err := p.SendGCode(fmt.Sprintf("M104 T%d S%d", hotendIndex, temperature)); err != nil {
		return errors.Wrap(err, ErrSetHotendTemperature.Error())
	}

	return nil
}

func (p *Printer) GetTemperatures() (temp *printer.Temperature, err error) {
	return nil, errors.New("Not implemented")
}
