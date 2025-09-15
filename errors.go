package tinygo_pullup

import (
	tinygoerrors "github.com/ralvarezdev/tinygo-errors"
)

const (
	// ErrorCodePullUpResistorStartNumber is the starting number for Pull-Up Resistor-related error codes.
	ErrorCodePullUpResistorStartNumber uint16 = 5200
)

const (
	ErrorCodePullUpResistorNilHandler = tinygoerrors.ErrorCode(iota + ErrorCodePullUpResistorStartNumber)
)
