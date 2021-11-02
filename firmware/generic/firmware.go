package generic

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/naueramant/go-3d-printer/printer"
)

func (p *Printer) GetFirmwareInformation() (*printer.FirmwareInformation, error) {
	res, err := p.SendCommand("M115")
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`([A-Z\_]+:)|(Cap:[A-Z\_]+:)`)

	cleaned := r.ReplaceAllString(res, "\n${0}")
	cleaned = strings.Replace(cleaned, "\n\n", "\n", -1)

	spl := strings.Split(cleaned, "\n")
	spl = spl[:len(spl)-1]

	info := printer.FirmwareInformation{}

	for _, v := range spl {
		if v == "" || v == "ok\n" {
			continue
		}

		if strings.HasPrefix(v, "Cap") {
			parseCapability(&info, v)
		} else {
			parseGeneralInformation(&info, v)
		}
	}

	return &info, nil
}

func parseGeneralInformation(info *printer.FirmwareInformation, s string) {
	v := strings.SplitN(s, ":", 2)

	switch v[0] {
	case "FIRMWARE_NAME":
		info.FirmwareName = v[1]
	case "EXTRUDER_COUNT":
		info.ExtruderCount, _ = strconv.Atoi(v[1])
	case "PROTOCOL_VERSION":
		info.ProtocolVersion = v[1]
	case "UUID":
		info.UUID = v[1]
	case "MACHINE_TYPE":
		info.MachineType = v[1]
	}
}

func parseCapability(info *printer.FirmwareInformation, s string) {
	if info.Capabilities == nil {
		info.Capabilities = make(printer.Capabilities)
	}

	s = strings.Replace(s, "Cap:", "", 1)
	v := strings.SplitN(s, ":", 2)
	info.Capabilities[v[0]] = v[1]
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
