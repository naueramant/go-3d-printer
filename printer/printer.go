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

	Extrude(amount, rate int) (err error) // OK, generic

	AutoHome() (err error) // OK, generic

	/*
		Temperature operations
	*/
	SetBedTemperature(temperature int) (err error)                 // OK, generic
	SetHotendTemperature(hotendIndex, temperature int) (err error) // OK, generic
	GetTemperatures() (temp *Temperature, err error)
	GetTemperaturesContinuesly(seconds int) (ch chan *Temperature, stop func(), err error)

	/*
		File operations
	*/
	ListFiles() (files []*File, err error) // OK, generic
	DeleteFile(path string) (err error)    // OK, generic
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
	Disconnect() (err error) // OK, generic

	SendGCode(gcode string) (result string, err error)       // OK, generic
	SendGCodes(gcode []string) (results []string, err error) // OK, generic

	/*
		PSU operation
	*/

	PowerOff() (err error) // OK, generic
	PowerOn() (err error)  // OK, generic
}
