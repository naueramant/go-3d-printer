package generic

import (
	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

func (p *Printer) StartPrint(path string) (err error) {
	// if _, err := p.SendCommands([]string{
	// 	fmt.Sprintf("M23 %s", path),
	// 	"M24",
	// }); err != nil {
	// 	return errors.Wrap(err, ErrStartPrint.Error())
	// }

	// TODO: re-evaluate reading of output from printer....

	return errors.New("Not implemented")
}

func (p *Printer) PausePrint() (err error) {
	if _, err := p.SendCommand("M25"); err != nil {
		return errors.Wrap(err, "Failed to pause print")
	}

	return nil
}

func (p *Printer) ResumePrint() (err error) {
	// TODO: Check if anything is printing

	if _, err := p.SendCommand("M24"); err != nil {
		return errors.Wrap(err, "Failed to resume print")
	}

	return nil
}

func (p *Printer) GetPrintProgress() (ch <-chan printer.PrintProgress, err error) {
	return nil, errors.New("Not implemented")
}
