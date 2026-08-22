package main

import (
	"slices"

	gcm "github.com/paralin/go-deadlock/protocol"
)

// IsValidMsg checks if the message is valid.
func IsValidMsg(msg gcm.EGCCitadelClientMessages) bool {
	_, ok := gcm.EGCCitadelClientMessages_name[int32(msg)]
	return ok && msg > gcm.EGCCitadelClientMessages(0)
}

func getSortedMsgIDs() []gcm.EGCCitadelClientMessages {
	var msgIds []gcm.EGCCitadelClientMessages
	for msgIDNum := range gcm.EGCCitadelClientMessages_name {
		msgID := gcm.EGCCitadelClientMessages(msgIDNum)
		msgIds = append(msgIds, msgID)
	}

	slices.Sort(msgIds)
	return msgIds
}
