package printer

type FirmwareCapabilities map[string]bool

type FirmwareInformation struct {
	Props map[string]string
	FirmwareVersion      string
	ProtocolVersion      string
	MachineType          string
	ExtruderCount        int
	UUID                 string
	FirmwareCapabilities FirmwareCapabilities
}

func (I FirmwareInformation)GetFirmwareVersion()string{
	if ver, ok := I.Props["FIRMWARE_VERSION"]; ok{
		return ver
	}
	return I.FirmwareVersion
}

func (I FirmwareInformation)GetFirmwareDate()string{
	if ver, ok := I.Props["FIRMWARE_DATE"]; ok{
		return ver
	}
	return "unknown"
}

func (I FirmwareInformation)GetFirmwareName()string{
	if ver, ok := I.Props["FIRMWARE_NAME"]; ok{
		return ver
	}
	return "unknown"
}
func (I FirmwareInformation)GetMachineType()string{
	if ver, ok := I.Props["ELECTRONICS"]; ok{
		return ver
	}
	return "unknown"
}
