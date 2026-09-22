package events

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-deadlock/protocol"
)

// Event is a Deadlock event.
type Event interface {
	// GetDeadlockEventMsgID returns the deadlock event message ID.
	GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages
	// GetEventBody returns the generated lite message.
	GetEventBody() protobuf.Message
	// GetEventName returns the event name.
	GetEventName() string
}
