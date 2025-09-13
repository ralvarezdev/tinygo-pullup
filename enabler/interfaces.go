//go:build tinygo && (rp2040 || rp2350)

package enabler

import (
	tinygopullup "github.com/ralvarezdev/tinygo-pullup"
)

type (
	// Handler is the interface to handle pull-up resistor operations for enabling features
	Handler interface {
		tinygopullup.Handler
		IsEnabled() bool
		IsDisabled() bool
	}
)
