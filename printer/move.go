package printer

import "errors"

var (
	ErrInvalidMoveMode = errors.New("Invalid move mode only Rapid (G0) and Linear (G1) is allowed")
)

type Position struct {
	X float64
	Y float64
	Z float64
}

type MoveMode uint8

const (
	MoveModeRapid = iota
	MoveModeLinear
)

func MoveModeToGCode(mode MoveMode) (code string, err error) {
	switch mode {
	case MoveModeRapid:
		return "G0", nil
	case MoveModeLinear:
		return "G1", nil
	default:
		return "", ErrInvalidMoveMode
	}
}
