//go:build tinygo && (rp2040 || rp2350)

package tinygo_pullup

import (
	tinygotypes "github.com/ralvarezdev/tinygo-types"
)

const (
	// ErrorCodePullUpResistorStartNumber is the starting number for Pull-Up Resistor-related error codes.
	ErrorCodePullUpResistorStartNumber uint16 = 5200
)

const (
	ErrorCodePullUpResistorNilHandler = tinygotypes.ErrorCode(iota + ErrorCodePullUpResistorStartNumber)
)
