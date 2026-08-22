package deadlock

import (
	bgcm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-steam/steamid"
)

// InviteLobbyMember attempts to invite a player to the current lobby.
func (d *Deadlock) InviteLobbyMember(playerID steamid.SteamId) {
	steamID := playerID.ToUint64()
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToLobby), &bgcm.CMsgInviteToLobby{
		SteamId: &steamID,
	})
}

// RespondLobbyInvite responds to a lobby invite.
func (d *Deadlock) RespondLobbyInvite(lobbyId uint64, accept bool) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCLobbyInviteResponse), &bgcm.CMsgLobbyInviteResponse{
		LobbyId: &lobbyId,
		Accept:  &accept,
	})
}
