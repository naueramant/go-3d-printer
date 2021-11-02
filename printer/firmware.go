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

type FirmwareInformation struct {
	FirmwareName    string
	ProtocolVersion string
	MachineType     string
	ExtruderCount   int
	UUID            string
	Capabilities    Capabilities
}

type Capabilities map[string]string
