package generic

import (
	"github.com/naueramant/go-3d-printer/printer"
	"github.com/pkg/errors"
)

func (p *Printer) GetFirmwareInformation() (info *printer.FirmwareInformation, err error) {
	// res, err := p.SendGCode("M115")
	// if err != nil {
	// 	return nil, err
	// }

	// TODO: rework how reading results work...

	return nil, errors.New("Not implemented")
}

/*
EXAMPLE

FIRMWARE_NAME:Marlin Ver 1.0.1 SOURCE_CODE_URL:https://github.com/MarlinFirmware/Marlin PROTOCOL_VERSION:1.0 MACHINE_TYPE:Ender-3 V2 EXTRUDER_COUNT:1 UUID:cede2a2f-41a2-4748-9b12-c55c62f367ff
Cap:SERIAL_XON_XOFF:0
Cap:BINARY_FILE_TRANSFER:0
Cap:EEPROM:1
Cap:VOLUMETRIC:1
Cap:AUTOREPORT_TEMP:1
Cap:PROGRESS:0
Cap:PRINT_JOB:1
Cap:AUTOLEVEL:0
Cap:Z_PROBE:0
Cap:LEVELING_DATA:0
Cap:BUILD_PERCENT:0
Cap:SOFTWARE_POWER:0
Cap:TOGGLE_LIGHTS:0
Cap:CASE_LIGHT_BRIGHTNESS:0
Cap:EMERGENCY_PARSER:0
Cap:PROMPT_SUPPORT:0
Cap:AUTOREPORT_SD_STATUS:0
Cap:THERMAL_PROTECTION:1
Cap:MOTION_MODES:0
Cap:CHAMBER_TEMPERATURE:0
ok
*/
