package events

import (
	"github.com/golang/protobuf/proto"
	bgcm "github.com/paralin/go-deadlock/protocol"
)

// InvitationCreated confirms that an invitation has been created.
type InvitationCreated struct {
	bgcm.CMsgInvitationCreated
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *InvitationCreated) GetDeadlockEventMsgID() bgcm.EGCCitadelClientMessages {
	return bgcm.EGCCitadelClientMessages(bgcm.EGCBaseMsg_k_EMsgGCInvitationCreated)
}

// GetEventBody returns the event body.
func (e *InvitationCreated) GetEventBody() proto.Message {
	return &e.CMsgInvitationCreated
}

// GetEventName returns the event name.
func (e *InvitationCreated) GetEventName() string {
	return "InvitationCreated"
}
