package main

import (
	"github.com/golang/protobuf/proto"

	dm "github.com/paralin/go-deadlock/protocol"
)

// The deadlock protos deviate from the Valve naming conventions in a few
// places. These maps record each deviation so the apigen heuristics can
// resolve every message ID.

// msgSenderOverrides overrides the heuristic-generated sender parties for each
// message. Messages with no protobuf definition are marked MsgSenderNone.
var msgSenderOverrides = map[dm.EGCCitadelClientMessages]MsgSender{
	dm.EGCCitadelClientMessages_k_EMsgGCToClientAccountStatsUpdated:               MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientHeroLabsSchedule:                  MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevBan:                            MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevBanResponse:                    MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotGetUserData:                  MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotGetUserDataResponse:          MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientTarotUserDataUpdated:              MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotPlaceTile:                    MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotPlaceTileResponse:            MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotApplyRune:                    MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotApplyRuneResponse:            MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotCompleteBoard:                MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotCompleteBoardResponse:        MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateTile:                MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateTileResponse:        MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateBoard:               MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateBoardResponse:       MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevGenerateTestTiles:         MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevGenerateTestTilesResponse: MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevModifyEnergy:              MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevModifyEnergyResponse:      MsgSenderNone,
}

// msgMethodNameOverrides overrides the generated method name for a message.
var msgMethodNameOverrides = map[dm.EGCCitadelClientMessages]string{}

// msgResponseOverrides overrides the response message ID for a request message.
var msgResponseOverrides = map[dm.EGCCitadelClientMessages]dm.EGCCitadelClientMessages{
	// k_EMsgClientToGCGetProfileCardResponse has no protobuf definition in the
	// deadlock protos, so treat GetProfileCard as fire-and-forget.
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCard: 0,
}

// msgProtoTypeOverrides overrides the proto type used for a message ID.
var msgProtoTypeOverrides = map[dm.EGCCitadelClientMessages]proto.Message{
	// The profile card update reuses the profile card message.
	dm.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated: &dm.CMsgCitadelProfileCard{},
	// The cheat report messages drop their Dev prefix in the proto names.
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReports:         &dm.CMsgClientToGCRequestCheatReports{},
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReportsResponse: &dm.CMsgClientToGCRequestCheatReportsResponse{},
}

// msgArgAsParameterOverrides forces a message argument to be passed as a parameter.
var msgArgAsParameterOverrides = map[dm.EGCCitadelClientMessages]bool{}

// msgEventNameOverrides overrides the generated event name for a message.
var msgEventNameOverrides = map[dm.EGCCitadelClientMessages]string{}
