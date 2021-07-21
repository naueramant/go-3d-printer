package serial

import (
	goserial "go.bug.st/serial.v1"
)

func GetSerialDevices() ([]string, error) {
	return goserial.GetPortsList()
}
