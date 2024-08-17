package printer

import "errors"

type Position struct {
	X float64
	Y float64
	Z float64
}

type MoveMode uint8

const (
	MoveModeRapid  MoveMode = 0
	MoveModeLinear MoveMode = 1
)

func MoveModeToGCode(mode MoveMode) (code string, err error) {
	switch mode {
	case MoveModeRapid:
		return "G0", nil
	case MoveModeLinear:
		return "G1", nil
	default:
		return "", errors.New("Invalid move mode only Rapid (G0) and Linear (G1) is allowed")
	}
}
