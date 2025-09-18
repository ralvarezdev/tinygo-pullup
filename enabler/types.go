package enabler

import (
	"machine"

	tinygoerrors "github.com/ralvarezdev/tinygo-errors"
	tinygopullup "github.com/ralvarezdev/tinygo-pullup"
)

type (
	// DefaultHandler is the default implementation of the Handler interface.
	DefaultHandler struct {
		tinygopullup.Handler
	}
)

// NewDefaultHandler creates a new default debug handler with the given pull-up handler.
//
// Parameters:
//
//	pullUpHandler: The pull-up handler to use
//
// Returns:
//
//	An instance of DefaultHandler, or an error if the pull-up handler is nil
func NewDefaultHandler(
	pullUpHandler tinygopullup.Handler,
) (*DefaultHandler, tinygoerrors.ErrorCode) {
	if pullUpHandler == nil {
		return nil, tinygopullup.ErrorCodePullUpResistorNilHandler
	}

	return &DefaultHandler{
		pullUpHandler,
	}, tinygoerrors.ErrorCodeNil
}

// NewDefaultHandlerFromPin creates a new default debug handler with a pull-up handler created from the given pin.
//
// Parameters:
//
//	pin: The pin to use for the pull-up handler
//
// Returns:
//
//	An instance of DefaultHandler
func NewDefaultHandlerFromPin(pin machine.Pin) *DefaultHandler {
	pullUpHandler := tinygopullup.NewDefaultHandler(pin)

	return &DefaultHandler{
		pullUpHandler,
	}
}

// IsEnabled checks if the debug mode is enabled.
//
// Returns:
//
// True if the debug mode is enabled, otherwise False
func (d *DefaultHandler) IsEnabled() bool {
	return d.IsShorted()
}

// IsDisabled checks if the debug mode is disabled.
//
// Returns:
//
// True if the debug mode is disabled, otherwise False
func (d *DefaultHandler) IsDisabled() bool {
	return d.IsOpen()
}
