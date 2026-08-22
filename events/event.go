package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-deadlock/protocol"
)

// Event is a Deadlock event.
type Event interface {
	// GetDeadlockEventMsgID returns the deadlock event message ID.
	GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages
	// GetEventBody event body.
	GetEventBody() proto.Message
	// GetEventName returns the event name.
	GetEventName() string
}
