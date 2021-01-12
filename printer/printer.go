package printer

import "io"

type Printer interface {
	/*
		Fan operations
	*/
	SetFanSpeed(fanIndex, speed int) (err error) // OK, generic

	/*
		Move operations
	*/
	EnableSteppers() (err error)  // OK, generic
	DisableSteppers() (err error) // OK, generic

	MoveAbsolute(x, y, z, rate int) (err error) // OK, generic
	MoveRelative(x, y, z, rate int) (err error) // OK, generic

	Extrude(amount, rate int) (err error)

	AutoHome() (err error) // OK, generic

	/*
		Temperature operations
	*/
	SetBedTemperature(temperature int) (err error)
	SetHotendTemperature(temperature int) (err error)
	GetTemperatures() (temp *Temperature, err error)

	// GetTemperaturesContinuesly returns a channel with temperatures from the
	// printer every "seconds"
	GetTemperaturesContinuesly(seconds int) (ch chan *Temperature, stop func(), err error)

	/*
		File operations
	*/
	ListFiles() (files []*File, err error)
	DeleteFile(path string) (err error)
	UploadFile(data io.Reader, path string) (err error)

	/*
		Job operations
	*/

	// TODO

	/*
		Firmware operations
	*/
	GetFirmwareInformation() (info *FirmwareInformation, err error)

	/*
		Printer operations
	*/
	Disconnect() (err error)

	SendGCode(gcode string) (result string, err error)
	SendGCodes(gcode []string) (results []string, err error)
}
