package generic

import (
	"errors"
	"io"

	"github.com/naueramant/go-3d-printer/printer"
)

func (p *Printer) ListFiles() (files *[]printer.File, err error) {
	return nil, errors.New("Not implemented")
}

func (p *Printer) DeleteFile(path string) (err error) {
	return errors.New("Not implemented")
}

func (p *Printer) UploadFile(data io.Reader, path string) (err error) {
	return errors.New("Not implemented")
}
