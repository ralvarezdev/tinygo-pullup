//go:build tinygo && (rp2040 || rp2350)

package tinygo_pullup

type (
	// Handler is the interface to handle pull-up resistor operations
	Handler interface {
		Setup()
		IsHigh() bool
		IsLow() bool
		IsOpen() bool
		IsShorted() bool
	}
)
