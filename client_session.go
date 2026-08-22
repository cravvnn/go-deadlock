package deadlock

import (
	"context"

	devents "github.com/paralin/go-deadlock/events"
	gcsdkm "github.com/paralin/go-deadlock/protocol"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-deadlock/state"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
)

// SetPlaying informs Steam we are playing / not playing Deadlock.
func (d *Deadlock) SetPlaying(playing bool) {
	if playing {
		d.client.GC.SetGamesPlayed(AppID)
	} else {
		d.client.GC.SetGamesPlayed()
		_ = d.accessState(func(ns *state.DeadlockState) (changed bool, err error) {
			ns.ClearState()
			return true, nil
		})
	}
}

// SayHello says hello to the deadlock server, in an attempt to get a session.
func (d *Deadlock) SayHello(haveCacheVersions ...*gcsdkm.CMsgSOCacheHaveVersion) {
	d.le.Debug("sending hello to GC")
	partnerAccType := gcsdkm.PartnerAccountType_PARTNER_NONE
	engine := gcsdkm.ESourceEngine_k_ESE_Source2
	var clientSessionNeed uint32 = 104
	d.write(uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientHello), &gcsdkm.CMsgClientHello{
		ClientLauncher:      &partnerAccType,
		Engine:              &engine,
		ClientSessionNeed:   &clientSessionNeed,
		SocacheHaveVersions: haveCacheVersions,
	})
}

// validateConnectionContext checks if the client is ready or not.
func (d *Deadlock) validateConnectionContext() (context.Context, error) {
	d.connectionCtxMtx.Lock()
	defer d.connectionCtxMtx.Unlock()

	cctx := d.connectionCtx
	if cctx == nil {
		return nil, ErrNotReady
	}

	select {
	case <-cctx.Done():
		return nil, ErrNotReady
	default:
		return cctx, nil
	}
}

// handleClientWelcome handles an incoming client welcome event.
func (d *Deadlock) handleClientWelcome(packet *gamecoordinator.GCPacket) error {
	welcome := &gcsdkm.CMsgClientWelcome{}
	if err := d.unmarshalBody(packet, welcome); err != nil {
		return err
	}

	d.le.Debug("received GC welcome")
	for _, cache := range welcome.GetUptodateSubscribedCaches() {
		d.RequestCacheSubscriptionRefresh(cache.GetOwnerSoid())
	}

	for _, cache := range welcome.GetOutofdateSubscribedCaches() {
		if err := d.cache.HandleSubscribed(cache); err != nil {
			d.le.WithError(err).Warn("unable to handle welcome cache")
		}
	}

	d.setConnectionStatus(gcsdkm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION, nil)
	d.emit(&devents.ClientWelcomed{Welcome: welcome})
	return nil
}

// handleConnectionStatus handles the connection status update event.
func (d *Deadlock) handleConnectionStatus(packet *gamecoordinator.GCPacket) error {
	stat := &gcsdkm.CMsgConnectionStatus{}
	if err := d.unmarshalBody(packet, stat); err != nil {
		return err
	}

	if stat.Status == nil {
		return nil
	}

	d.setConnectionStatus(*stat.Status, stat)
	return nil
}

// setConnectionStatus sets the connection status, and emits an event.
// NOTE: do not call from inside accessState.
func (d *Deadlock) setConnectionStatus(
	connStatus gcsdkm.GCConnectionStatus,
	update *gcsdkm.CMsgConnectionStatus,
) {
	_ = d.accessState(func(ns *state.DeadlockState) (changed bool, err error) {
		if ns.ConnectionStatus == connStatus {
			return false, nil
		}

		oldState := ns.ConnectionStatus
		d.le.WithField("old", oldState.String()).
			WithField("new", connStatus.String()).
			Debug("connection status changed")
		d.emit(&devents.GCConnectionStatusChanged{
			OldState: oldState,
			NewState: connStatus,
			Update:   update,
		})

		ns.ClearState() // every time the state changes, we lose the lobbies / etc
		ns.ConnectionStatus = connStatus
		d.connectionCtxMtx.Lock()
		if d.connectionCtxCancel != nil {
			d.connectionCtxCancel()
			d.connectionCtxCancel = nil
			d.connectionCtx = nil
		}
		if connStatus == gcsdkm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION {
			d.connectionCtx, d.connectionCtxCancel = context.WithCancel(context.Background())
		}
		d.connectionCtxMtx.Unlock()
		return true, nil
	})
}
