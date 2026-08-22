package state

import (
	gcmm "github.com/paralin/go-deadlock/protocol"
)

// DeadlockState is a snapshot of the client state at a point in time.
type DeadlockState struct {
	// ConnectionStatus is the status of the connection to the GC.
	ConnectionStatus gcmm.GCConnectionStatus
	// Lobby is the current lobby object.
	Lobby *gcmm.CSOCitadelLobby
	// Party is the current party object.
	Party *gcmm.CSOCitadelParty
	// LastConnectionStatusUpdate is the last connection state update we received.
	LastConnectionStatusUpdate *gcmm.CMsgConnectionStatus
}

// ClearState clears everything.
func (s *DeadlockState) ClearState() {
	*s = DeadlockState{ConnectionStatus: gcmm.GCConnectionStatus_GCConnectionStatus_NO_SESSION}
}

// IsReady checks if the client is ready to receive requests.
func (s *DeadlockState) IsReady() bool {
	return s.ConnectionStatus == gcmm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
}
