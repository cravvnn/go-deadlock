package events

import (
	"github.com/paralin/go-deadlock/state"
)

// ClientStateChanged is emitted whenever anything about the client state changes.
type ClientStateChanged struct {
	// OldState is the old state.
	OldState state.DeadlockState
	// NewState is the new state.
	NewState state.DeadlockState
}
