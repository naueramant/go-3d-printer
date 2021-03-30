package printer

type FirmwareType uint8

const (
	FirmwareTypeGeneric = iota
	FirmwareTypeMarlin
	FirmwareTypeRepRap
	FirmwareTypeRepetier
	FirmwareTypeSmoothie
	FirmwareTypePrusa
)

type FirmwareCapabilities map[string]bool

type FirmwareInformation struct {
	FirmwareVersion      string
	ProtocolVersion      string
	MachineType          string
	ExtruderCount        int
	UUID                 string
	FirmwareCapabilities FirmwareCapabilities
}
