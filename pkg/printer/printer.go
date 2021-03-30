package printer

import "io"

type Printer interface {
	/*
		Fan operations
	*/

	// Set a fan speed
	SetFanSpeed(fanIndex, speed int) (err error) // OK, all

	/*
		Move operations
	*/

	// Enable all steppers
	EnableSteppers() (err error) // OK, all

	// Disable all steppers immediately
	DisableSteppers() (err error) // OK, all

	// Move to a absolute position defined relative
	// to the home position
	MoveAbsolute(x, y, z, rate int, mode MoveMode) (err error) // OK, all

	// Move relative to the current position
	MoveRelative(x, y, z, rate int, mode MoveMode) (err error) // OK, all

	// Extrude filament
	Extrude(amount, rate int) (err error) // OK, all

	// Auto home the printer head
	AutoHome() (err error) // OK, generic

	// All movement will stop
	EmergencyStop() (err error) // OK, all

	/*
		Temperature operations
	*/

	// Set the print bed temperature
	SetBedTemperature(temperature int) (err error) // OK, all

	// Set a hotend temperature
	SetHotendTemperature(hotendIndex, temperature int) (err error) // OK, all

	// Get temperatures of hotend and bed
	GetTemperatures() (temp *Temperature, err error)

	/*
		File operations
	*/

	// List SD Card files
	ListFiles() (files []File, err error) // OK, generic

	// Delete a file from the SD Card
	DeleteFile(path string) (err error) // OK, generic

	// Upload a file to the SD Card
	UploadFile(data io.Reader, path string) (err error)

	/*
		Job operations
	*/

	// TODO

	/*
		Firmware operations
	*/

	// Get firmware information
	GetFirmwareInformation() (info *FirmwareInformation, err error)

	/*
		Serial operations
	*/

	// Disconnect from the printer
	Disconnect() (err error) // OK, all

	// Send a GCode to the printer and get the result back
	SendGCode(gcode string) (result string, err error) // OK, generic

	// Send a Batch of GCodes to the printer and get the results back
	SendGCodes(gcode []string) (results []string, err error) // OK, generic

	/*
		PSU operation
	*/

	// Power on the high voltage PSU
	PowerOn() (err error) // OK, all

	// Power off the high voltage PSU
	PowerOff() (err error) // OK, all
}
