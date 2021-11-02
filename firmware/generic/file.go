package generic

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

func (p *Printer) ListFiles() (files []printer.File, err error) {
	res, err := p.SendCommand("M20")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to list files")
	}

	readFileLine := false
	for _, l := range strings.Split(res, "\n") {
		if !readFileLine && l == "Begin file list" {
			readFileLine = true

			continue
		}

		if readFileLine && l == "End file list" {
			break
		}

		if readFileLine {
			r := strings.Split(l, " ")

			if len(r) != 2 {
				continue
			}

			sizeInt, err := strconv.Atoi(r[1])
			if err != nil {
				continue
			}

			files = append(files, printer.File{
				Path: r[0],
				Size: sizeInt,
			})
		}
	}

	return files, nil
}

func (p *Printer) GetFileLongPath(path string) (name string, err error) {
	// res, err := p.SendCommand(fmt.Sprintf("M33 %s", path))
	// if err != nil {
	// 	return "", errors.Wrap(err, ErrListFiles.Error())
	// }

	// TODO: rework how reading results work...

	return "", errors.New("Not implemented")
}

func (p *Printer) DeleteFile(path string) (err error) {
	res, err := p.SendCommand(fmt.Sprintf("M30 %s", path))
	if err != nil {
		return errors.Wrap(err, "Failed to delete file")
	}

	if strings.Contains(res, "Deletion failed") {
		return errors.New("Printer reported that file deletion failed")
	}

	return nil
}

func (p *Printer) UploadFile(data io.Reader, path string) (err error) {
	return errors.New("Not implemented")
}
