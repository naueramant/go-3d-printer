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

type Capabilities map[string]bool

type FirmwareInformation struct {
	FirmwareName    string
	SourceCodeURL   string
	ProtocolVersion string
	MachineType     string
	ExtruderCount   int
	UUID            string
	Capabilities    Capabilities
}
