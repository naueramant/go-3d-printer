package generic

import (
	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

func (p *Printer) SetBedTemperature(temperature int) (err error) {
	return errors.New("Not implemented")
}

func (p *Printer) SetHotendTemperature(temperature int) (err error) {
	return errors.New("Not implemented")
}

func (p *Printer) GetTemperatures() (temp *printer.Temperature, err error) {
	return nil, errors.New("Not implemented")
}

func (p *Printer) GetTemperaturesContinuesly(seconds int) (ch chan *printer.Temperature, stop func(), err error) {
	return nil, nil, errors.New("Not implemented")
}
