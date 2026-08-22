package main

import (
	"fmt"

	gcm "github.com/paralin/go-deadlock/protocol"
)

type generatedEventHandler struct {
	msgID     gcm.EGCCitadelClientMessages
	eventName string
	eventType *ProtoType
}

// buildGeneratedEventHandler builds a generated event handler.
func buildGeneratedEventHandler(
	msgID gcm.EGCCitadelClientMessages,
	protoMap map[string]*ProtoType,
	eventImports map[string]struct{},
) (*generatedEventHandler, error) {
	eventProtoType, err := LookupMessageProtoType(protoMap, msgID)
	if err != nil {
		return nil, err
	}
	eventImports[eventProtoType.Pak.Path()] = struct{}{}

	return &generatedEventHandler{
		msgID:     msgID,
		eventName: GetMessageEventName(msgID),
		eventType: eventProtoType,
	}, nil
}

func (g *generatedEventHandler) generateComment() string {
	return fmt.Sprintf("// %s event.\n// MessageID: %s\n", g.eventName, g.msgID.String())
}
