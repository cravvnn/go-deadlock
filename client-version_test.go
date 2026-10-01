package deadlock

import (
	"io"
	"testing"

	devents "github.com/paralin/go-deadlock/events"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/sirupsen/logrus"
)

// TestClientVersionUpdated checks that the GC's new version notice reaches
// event consumers with its client version.
func TestClientVersionUpdated(t *testing.T) {
	var emitted []any
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	client := NewWithCoordinator(&requestCoordinator{}, func(event any) { emitted = append(emitted, event) }, logger)
	t.Cleanup(client.Close)

	body, err := (&gcsm.CMsgGCClientVersionUpdated{ClientVersion: new(uint32(6727))}).MarshalVT()
	if err != nil {
		t.Fatal(err)
	}
	client.HandleGCPacket(&gamecoordinator.GCPacket{AppId: AppID, MsgType: uint32(gcsm.EGCItemMsg_k_EMsgGCClientVersionUpdated), Body: body})

	if len(emitted) != 1 {
		t.Fatalf("emitted %d events, want 1", len(emitted))
	}
	updated, ok := emitted[0].(*devents.ClientVersionUpdated)
	if !ok || updated.GetClientVersion() != 6727 {
		t.Fatalf("emitted %#v, want client version 6727", emitted[0])
	}
}
