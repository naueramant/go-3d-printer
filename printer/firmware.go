package printer

type FirmwareCapabilities map[string]bool

type FirmwareInformation struct {
	FirmwareVersion      string
	ProtocolVersion      string
	MachineType          string
	ExtruderCount        int
	UUID                 string
	FirmwareCapabilities FirmwareCapabilities
}
