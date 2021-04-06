package printer

import "io"

type Printer interface {
	/*
		Fan operations
	*/

	// Set fan speed
	SetFanSpeed(fanIndex, speed int) (err error) // OK, generic

	/*
		Move operations
	*/

	// Enable all steppers
	EnableSteppers() (err error) // OK, generic

	// Disable all steppers immediately
	DisableSteppers() (err error) // OK, generic

	// Move to a absolute position defined relative
	// to the home position
	MoveAbsolute(x, y, z, rate int, mode MoveMode) (err error) // OK, generic

	// Move relative to the current position
	MoveRelative(x, y, z, rate int, mode MoveMode) (err error) // OK, generic

	// Extrude filament
	Extrude(amount, rate int) (err error)

	// Retract filament
	Retract(amount, rate int) (err error)

	// Auto home the printer head
	AutoHome() (err error) // OK, generic

	// All movement will stop
	EmergencyStop() (err error) // OK, generic

	// Get current print head position
	GetCurrentPosition() (pos *Position, err error)

	/*
		Temperature operations
	*/

	// Set the print bed temperature
	SetBedTemperature(temperature int) (err error) // OK, generic

	// Set a hotend temperature
	SetHotendTemperature(hotendIndex, temperature int) (err error) // OK, generic

	// Get temperatures of hotend and bed
	GetTemperatures() (temp *Temperature, err error)

	/*
		File operations
	*/

	// List SD Card files
	ListFiles() (files []File, err error) // OK, generic

	// Get the long name for a file based on the DOS 8.3 path
	GetLongPath(path string) (name string, err error)

	// Delete a file from the SD Card
	DeleteFile(path string) (err error) // OK, generic

	// Upload a file to the SD Card
	UploadFile(data io.Reader, path string) (err error)

	/*
		Job operations
	*/

	// Print a file on the SD Card
	StartPrint(path string) (err error)

	// Pause current print job
	PausePrint() (err error) // OK, generic

	// Resume current print job
	ResumePrint() (err error)

	// Get current print job status
	GetPrintProgress() (stats *PrintProgress, err error)

	/*
		Firmware operations
	*/

	// Get firmware information
	GetFirmwareInformation() (info *FirmwareInformation, err error)

	/*
		Serial operations
	*/

	// Disconnect from the printer
	Disconnect() (err error) // OK, generic

	// Send a GCode to the printer and get the result back
	SendGCode(gcode string) (result string, err error) // OK, generic

	// TODO: consider removing
	// Send a Batch of GCodes to the printer and get the results back
	SendGCodes(gcode []string) (results []string, err error) // OK, generic

	/*
		PSU operation
	*/

	// Power on the high voltage PSU
	PowerOn() (err error) // OK, generic

	// Power off the high voltage PSU
	PowerOff() (err error) // OK, generic
}
