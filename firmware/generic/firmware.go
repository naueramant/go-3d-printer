package generic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/naueramant/go-3d-printer/printer"
)

func (p *Printer) GetFirmwareInformation() (*printer.FirmwareInformation, error) {
	res, err := p.SendCommand("M115")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(res, "\n")

	firmwareLine := lines[0]

	info, err := parseFirmwareInfo(firmwareLine)
	if err != nil {
		return nil, fmt.Errorf("error parsing firmware information: %w", err)
	}

	caps, err := parseCapabilities(lines)
	if err != nil {
		return nil, fmt.Errorf("error parsing capabilities: %w", err)
	}

	info.Capabilities = caps

	return info, nil
}

func parseFirmwareInfo(line string) (*printer.FirmwareInformation, error) {
	// The firmware information line is always in the format:
	// FIRMWARE_NAME:XXX SOURCE_CODE_URL:XXX PROTOCOL_VERSION:XXX MACHINE_TYPE:XXX EXTRUDER_COUNT:XXX UUID:XXX
	// We can therefore assume that values are always in between the field name and the next field name.

	firmwareName := strings.Split(strings.Split(line, "FIRMWARE_NAME:")[1], " SOURCE_CODE_URL:")[0]

	sourceCodeURL := strings.Split(strings.Split(line, "SOURCE_CODE_URL:")[1], " PROTOCOL_VERSION:")[0]

	protocolVersion := strings.Split(strings.Split(line, "PROTOCOL_VERSION:")[1], " MACHINE_TYPE:")[0]

	machineType := strings.Split(strings.Split(line, "MACHINE_TYPE:")[1], " EXTRUDER_COUNT:")[0]

	extruderCountStr := strings.Split(strings.Split(line, "EXTRUDER_COUNT:")[1], " UUID:")[0]
	extruderCountInt, err := strconv.Atoi(extruderCountStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing extruder count: %w", err)
	}

	// get text after UUID:
	uuid := strings.Split(line, " UUID:")[1]

	return &printer.FirmwareInformation{
		FirmwareName:    firmwareName,
		SourceCodeURL:   sourceCodeURL,
		ProtocolVersion: protocolVersion,
		MachineType:     machineType,
		ExtruderCount:   extruderCountInt,
		UUID:            uuid,
	}, nil
}

func parseCapabilities(lines []string) (printer.Capabilities, error) {
	capabilities := make(printer.Capabilities)

	for _, line := range lines {
		if !strings.HasPrefix(line, "Cap:") {
			continue
		}

		line = strings.TrimPrefix(line, "Cap:")

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]

		value, err := strconv.ParseBool(parts[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing capability value: %w", err)
		}

		capabilities[key] = value
	}

	return capabilities, nil
}
