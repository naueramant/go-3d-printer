package printer

import "io"

type Printer interface {
	/*
		Fan operations
	*/
	SetFanSpeed(fanIndex, speed int) (err error)

	/*
		Move operations
	*/
	EnableSteppers() (err error)
	DisableSteppers() (err error)

	MoveAbsolute(x, y, z int) (err error)
	MoveRelative(x, y, z int) (err error)

	Extrude(extruderIndex, n int) (err error)

	AutoHome() (err error)

	/*
		Temperature operations
	*/
	SetBedTemperature(temperature int) (err error)
	SetHotendTemperature(temperature int) (err error)

	/*
		File operations
	*/
	ListFiles() (files *[]File, err error)
	DeleteFile(path string) (err error)
	UploadFile(data io.Reader, path string) (err error)

	/*
		Job operations
	*/

	// TODO

	/*
		Firmware operations
	*/
	GetFirmwareInformation() (info FirmwareInformation, err error)

	/*
		Printer operations
	*/
	SendGCode(gcode string) (result string, err error)
	SendGCodes(gcode []string) (results []string, err error)
}
