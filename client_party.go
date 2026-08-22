package deadlock

import (
	bgcm "github.com/paralin/go-deadlock/protocol"
)

// RespondPartyInvite attempts to respond to a party invite.
func (d *Deadlock) RespondPartyInvite(partyId uint64, accept bool) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCPartyInviteResponse), &bgcm.CMsgPartyInviteResponse{
		PartyId: &partyId,
		Accept:  &accept,
	})
}
