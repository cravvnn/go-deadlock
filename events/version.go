package events

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	bgcm "github.com/paralin/go-deadlock/protocol"
)

// ClientVersionUpdated is emitted when the GC starts serving a new game client
// version. ClientVersion matches ClientVersion in the game's steam.inf.
type ClientVersionUpdated struct {
	bgcm.CMsgGCClientVersionUpdated
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *ClientVersionUpdated) GetDeadlockEventMsgID() bgcm.EGCCitadelClientMessages {
	return bgcm.EGCCitadelClientMessages(bgcm.EGCItemMsg_k_EMsgGCClientVersionUpdated)
}

// GetEventBody returns the event body.
func (e *ClientVersionUpdated) GetEventBody() protobuf.Message {
	return &e.CMsgGCClientVersionUpdated
}

// GetEventName returns the event name.
func (e *ClientVersionUpdated) GetEventName() string {
	return "ClientVersionUpdated"
}
