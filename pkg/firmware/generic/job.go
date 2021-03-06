package generic

import (
	"github.com/naueramant/go-3d-printer/pkg/printer"
	"github.com/pkg/errors"
)

var (
	ErrStartPrint  = errors.New("Start print failed")
	ErrPausePrint  = errors.New("Pause print failed")
	ErrResumePrint = errors.New("Resume print failed")
)

func (p *Printer) StartPrint(path string) (err error) {
	// if _, err := p.SendGCodes([]string{
	// 	fmt.Sprintf("M23 %s", path),
	// 	"M24",
	// }); err != nil {
	// 	return errors.Wrap(err, ErrStartPrint.Error())
	// }

	// TODO: re-evaluate reading of output from printer....

	return errors.New("Not implemented")
}

func (p *Printer) PausePrint() (err error) {
	if _, err := p.SendGCode("M25"); err != nil {
		return errors.Wrap(err, ErrPausePrint.Error())
	}

	return nil
}

func (p *Printer) ResumePrint() (err error) {
	// TODO: Check if anything is printing

	if _, err := p.SendGCode("M24"); err != nil {
		return errors.Wrap(err, ErrPausePrint.Error())
	}

	return nil
}

func (p *Printer) GetPrintProgress() (stats *printer.PrintProgress, err error) {
	// res, err := p.SendGCode("M27 S0")
	// if err != nil {
	// 	return nil, errors.Wrap(err, ErrPausePrint.Error())
	// }

	// TODO: rework how reading results work...

	// OUTPUT:

	return nil, errors.New("Not implemented")
}
