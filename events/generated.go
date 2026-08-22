package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-deadlock/protocol"
)

// BookUpdated event.
// MessageID: k_EMsgGCToClientBookUpdated
type BookUpdated struct {
	protocol.CMsgGCToClientBookUpdated
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *BookUpdated) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientBookUpdated
}

// GetEventBody returns the event body.
func (e *BookUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientBookUpdated
}

// GetEventName returns the event name.
func (e *BookUpdated) GetEventName() string {
	return "BookUpdated"
}

// CanRejoinParty event.
// MessageID: k_EMsgGCToClientCanRejoinParty
type CanRejoinParty struct {
	protocol.CMsgGCToClientCanRejoinParty
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *CanRejoinParty) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientCanRejoinParty
}

// GetEventBody returns the event body.
func (e *CanRejoinParty) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCanRejoinParty
}

// GetEventName returns the event name.
func (e *CanRejoinParty) GetEventName() string {
	return "CanRejoinParty"
}

// CommendNotification event.
// MessageID: k_EMsgGCToClientCommendNotification
type CommendNotification struct {
	protocol.CMsgGCToClientCommendNotification
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *CommendNotification) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientCommendNotification
}

// GetEventBody returns the event body.
func (e *CommendNotification) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCommendNotification
}

// GetEventName returns the event name.
func (e *CommendNotification) GetEventName() string {
	return "CommendNotification"
}

// DevAnnouncements event.
// MessageID: k_EMsgGCToClientDevAnnouncements
type DevAnnouncements struct {
	protocol.CMsgGCToClientDevAnnouncements
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *DevAnnouncements) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientDevAnnouncements
}

// GetEventBody returns the event body.
func (e *DevAnnouncements) GetEventBody() proto.Message {
	return &e.CMsgGCToClientDevAnnouncements
}

// GetEventName returns the event name.
func (e *DevAnnouncements) GetEventName() string {
	return "DevAnnouncements"
}

// DevPlaytestStatus event.
// MessageID: k_EMsgGCToClientDevPlaytestStatus
type DevPlaytestStatus struct {
	protocol.CMsgGCToClientDevPlaytestStatus
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *DevPlaytestStatus) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientDevPlaytestStatus
}

// GetEventBody returns the event body.
func (e *DevPlaytestStatus) GetEventBody() proto.Message {
	return &e.CMsgGCToClientDevPlaytestStatus
}

// GetEventName returns the event name.
func (e *DevPlaytestStatus) GetEventName() string {
	return "DevPlaytestStatus"
}

// MatchmakingStopped event.
// MessageID: k_EMsgGCToClientMatchmakingStopped
type MatchmakingStopped struct {
	protocol.CMsgGCToClientMatchmakingStopped
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *MatchmakingStopped) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientMatchmakingStopped
}

// GetEventBody returns the event body.
func (e *MatchmakingStopped) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMatchmakingStopped
}

// GetEventName returns the event name.
func (e *MatchmakingStopped) GetEventName() string {
	return "MatchmakingStopped"
}

// PartyChatMsg event.
// MessageID: k_EMsgGCToClientPartyChatMsg
type PartyChatMsg struct {
	protocol.CMsgGCToClientPartyChatMsg
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *PartyChatMsg) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyChatMsg
}

// GetEventBody returns the event body.
func (e *PartyChatMsg) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPartyChatMsg
}

// GetEventName returns the event name.
func (e *PartyChatMsg) GetEventName() string {
	return "PartyChatMsg"
}

// PartyEvent event.
// MessageID: k_EMsgGCToClientPartyEvent
type PartyEvent struct {
	protocol.CMsgGCToClientPartyEvent
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *PartyEvent) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyEvent
}

// GetEventBody returns the event body.
func (e *PartyEvent) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPartyEvent
}

// GetEventName returns the event name.
func (e *PartyEvent) GetEventName() string {
	return "PartyEvent"
}

// ProfileCardUpdated event.
// MessageID: k_EMsgGCToClientProfileCardUpdated
type ProfileCardUpdated struct {
	protocol.CMsgCitadelProfileCard
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *ProfileCardUpdated) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated
}

// GetEventBody returns the event body.
func (e *ProfileCardUpdated) GetEventBody() proto.Message {
	return &e.CMsgCitadelProfileCard
}

// GetEventName returns the event name.
func (e *ProfileCardUpdated) GetEventName() string {
	return "ProfileCardUpdated"
}

// RankDataUpdate event.
// MessageID: k_EMsgGCToClientRankDataUpdate
type RankDataUpdate struct {
	protocol.CMsgGCToClientRankDataUpdate
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *RankDataUpdate) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientRankDataUpdate
}

// GetEventBody returns the event body.
func (e *RankDataUpdate) GetEventBody() proto.Message {
	return &e.CMsgGCToClientRankDataUpdate
}

// GetEventName returns the event name.
func (e *RankDataUpdate) GetEventName() string {
	return "RankDataUpdate"
}

// SDRTicket event.
// MessageID: k_EMsgGCToClientSDRTicket
type SDRTicket struct {
	protocol.CMsgGCToClientSDRTicket
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *SDRTicket) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientSDRTicket
}

// GetEventBody returns the event body.
func (e *SDRTicket) GetEventBody() proto.Message {
	return &e.CMsgGCToClientSDRTicket
}

// GetEventName returns the event name.
func (e *SDRTicket) GetEventName() string {
	return "SDRTicket"
}

// UpdateHeroReleaseVoteTally event.
// MessageID: k_EMsgGCToClientUpdateHeroReleaseVoteTally
type UpdateHeroReleaseVoteTally struct {
	protocol.CMsgGCToClientUpdateHeroReleaseVoteTally
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *UpdateHeroReleaseVoteTally) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientUpdateHeroReleaseVoteTally
}

// GetEventBody returns the event body.
func (e *UpdateHeroReleaseVoteTally) GetEventBody() proto.Message {
	return &e.CMsgGCToClientUpdateHeroReleaseVoteTally
}

// GetEventName returns the event name.
func (e *UpdateHeroReleaseVoteTally) GetEventName() string {
	return "UpdateHeroReleaseVoteTally"
}
