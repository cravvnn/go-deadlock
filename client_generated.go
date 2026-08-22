package deadlock

import (
	"context"

	"github.com/paralin/go-deadlock/events"
	"github.com/paralin/go-deadlock/protocol"
)

// CancelChallenge cancels a challenge.
// Request ID: k_EMsgClientToGCChallengeCancel
// Response ID: k_EMsgClientToGCChallengeCancelResponse
// Request type: CMsgClientToGCChallengeCancel
// Response type: CMsgClientToGCChallengeCancelResponse
func (d *Deadlock) CancelChallenge(
	ctx context.Context,
	challengeID uint64,
) (*protocol.CMsgClientToGCChallengeCancelResponse, error) {
	req := &protocol.CMsgClientToGCChallengeCancel{
		ChallengeId: &challengeID,
	}
	resp := &protocol.CMsgClientToGCChallengeCancelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeCancel),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeCancelResponse),
		resp,
	)
}

// ClaimChallenge claims a challenge.
// Request ID: k_EMsgClientToGCChallengeClaim
// Response ID: k_EMsgClientToGCChallengeClaimResponse
// Request type: CMsgClientToGCChallengeClaim
// Response type: CMsgClientToGCChallengeClaimResponse
func (d *Deadlock) ClaimChallenge(
	ctx context.Context,
	challengeID uint64,
) (*protocol.CMsgClientToGCChallengeClaimResponse, error) {
	req := &protocol.CMsgClientToGCChallengeClaim{
		ChallengeId: &challengeID,
	}
	resp := &protocol.CMsgClientToGCChallengeClaimResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeClaim),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeClaimResponse),
		resp,
	)
}

// CreateParty creates a party.
// Request ID: k_EMsgClientToGCPartyCreate
// Response ID: k_EMsgClientToGCPartyCreateResponse
// Request type: CMsgClientToGCPartyCreate
// Response type: CMsgClientToGCPartyCreateResponse
func (d *Deadlock) CreateParty(
	ctx context.Context,
	req *protocol.CMsgClientToGCPartyCreate,
) (*protocol.CMsgClientToGCPartyCreateResponse, error) {
	resp := &protocol.CMsgClientToGCPartyCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreate),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreateResponse),
		resp,
	)
}

// FindHeroBuilds finds hero builds.
// Request ID: k_EMsgClientToGCFindHeroBuilds
// Response ID: k_EMsgClientToGCFindHeroBuildsResponse
// Request type: CMsgClientToGCFindHeroBuilds
// Response type: CMsgClientToGCFindHeroBuildsResponse
func (d *Deadlock) FindHeroBuilds(
	ctx context.Context,
	authorAccountID uint32,
	heroID uint32,
	language []int32,
	searchText string,
	heroBuildID uint32,
	tags []uint32,
) (*protocol.CMsgClientToGCFindHeroBuildsResponse, error) {
	req := &protocol.CMsgClientToGCFindHeroBuilds{
		AuthorAccountId: &authorAccountID,
		HeroId:          &heroID,
		Language:        language,
		SearchText:      &searchText,
		HeroBuildId:     &heroBuildID,
		Tags:            tags,
	}
	resp := &protocol.CMsgClientToGCFindHeroBuildsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuilds),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuildsResponse),
		resp,
	)
}

// GetAccountLeaderboards gets account leaderboards.
// Request ID: k_EMsgClientToGCGetAccountLeaderboards
// Response ID: k_EMsgClientToGCGetAccountLeaderboardsResponse
// Request type: CMsgClientToGCGetAccountLeaderboards
// Response type: CMsgClientToGCGetAccountLeaderboardsResponse
func (d *Deadlock) GetAccountLeaderboards(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetAccountLeaderboardsResponse, error) {
	req := &protocol.CMsgClientToGCGetAccountLeaderboards{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetAccountLeaderboardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboards),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboardsResponse),
		resp,
	)
}

// GetAccountMatchReports gets account match reports.
// Request ID: k_EMsgClientToGCGetAccountMatchReports
// Response ID: k_EMsgClientToGCGetAccountMatchReportsResponse
// Request type: CMsgClientToGCGetAccountMatchReports
// Response type: CMsgClientToGCGetAccountMatchReportsResponse
func (d *Deadlock) GetAccountMatchReports(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCGetAccountMatchReportsResponse, error) {
	req := &protocol.CMsgClientToGCGetAccountMatchReports{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCGetAccountMatchReportsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReports),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReportsResponse),
		resp,
	)
}

// GetAccountStats gets account stats.
// Request ID: k_EMsgClientToGCGetAccountStats
// Response ID: k_EMsgClientToGCGetAccountStatsResponse
// Request type: CMsgClientToGCGetAccountStats
// Response type: CMsgClientToGCGetAccountStatsResponse
func (d *Deadlock) GetAccountStats(
	ctx context.Context,
	accountID uint32,
	devAccessHint bool,
	friendAccessHint bool,
) (*protocol.CMsgClientToGCGetAccountStatsResponse, error) {
	req := &protocol.CMsgClientToGCGetAccountStats{
		AccountId:        &accountID,
		DevAccessHint:    &devAccessHint,
		FriendAccessHint: &friendAccessHint,
	}
	resp := &protocol.CMsgClientToGCGetAccountStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStats),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStatsResponse),
		resp,
	)
}

// GetActiveMatches gets active matches.
// Request ID: k_EMsgClientToGCGetActiveMatches
// Response ID: k_EMsgClientToGCGetActiveMatchesResponse
// Request type: CMsgClientToGCGetActiveMatches
// Response type: CMsgClientToGCGetActiveMatchesResponse
func (d *Deadlock) GetActiveMatches(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetActiveMatchesResponse, error) {
	req := &protocol.CMsgClientToGCGetActiveMatches{}
	resp := &protocol.CMsgClientToGCGetActiveMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatches),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatchesResponse),
		resp,
	)
}

// GetBook gets a book.
// Request ID: k_EMsgClientToGCGetBook
// Response ID: k_EMsgClientToGCGetBookResponse
// Request type: CMsgClientToGCGetBook
// Response type: CMsgClientToGCGetBookResponse
func (d *Deadlock) GetBook(
	ctx context.Context,
	bookID uint32,
) (*protocol.CMsgClientToGCGetBookResponse, error) {
	req := &protocol.CMsgClientToGCGetBook{
		BookId: &bookID,
	}
	resp := &protocol.CMsgClientToGCGetBookResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetBook),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetBookResponse),
		resp,
	)
}

// GetDiscordLink gets a discord link.
// Request ID: k_EMsgClientToGCGetDiscordLink
// Response ID: k_EMsgClientToGCGetDiscordLinkResponse
// Request type: CMsgClientToGCGetDiscordLink
// Response type: CMsgClientToGCGetDiscordLinkResponse
func (d *Deadlock) GetDiscordLink(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetDiscordLinkResponse, error) {
	req := &protocol.CMsgClientToGCGetDiscordLink{}
	resp := &protocol.CMsgClientToGCGetDiscordLinkResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLink),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLinkResponse),
		resp,
	)
}

// GetFriendGameStatus gets friend game status.
// Request ID: k_EMsgClientToGCGetFriendGameStatus
// Response ID: k_EMsgClientToGCGetFriendGameStatusResponse
// Request type: CMsgClientToGCGetFriendGameStatus
// Response type: CMsgClientToGCGetFriendGameStatusResponse
func (d *Deadlock) GetFriendGameStatus(
	ctx context.Context,
	includeInvited bool,
) (*protocol.CMsgClientToGCGetFriendGameStatusResponse, error) {
	req := &protocol.CMsgClientToGCGetFriendGameStatus{
		IncludeInvited: &includeInvited,
	}
	resp := &protocol.CMsgClientToGCGetFriendGameStatusResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatus),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatusResponse),
		resp,
	)
}

// GetHeroChoice gets a hero choice.
// Request ID: k_EMsgClientToGCGetHeroChoice
// Response ID: k_EMsgClientToGCGetHeroChoiceResponse
// Request type: CMsgClientToGCGetHeroChoice
// Response type: CMsgClientToGCGetHeroChoiceResponse
func (d *Deadlock) GetHeroChoice(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetHeroChoiceResponse, error) {
	req := &protocol.CMsgClientToGCGetHeroChoice{}
	resp := &protocol.CMsgClientToGCGetHeroChoiceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoice),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoiceResponse),
		resp,
	)
}

// GetHeroMMRRankings gets hero mmr rankings.
// Request ID: k_EMsgClientToGCGetHeroMMRRankings
// Response ID: k_EMsgClientToGCGetHeroMMRRankingsResponse
// Request type: CMsgClientToGCGetHeroMMRRankings
// Response type: CMsgClientToGCGetHeroMMRRankingsResponse
func (d *Deadlock) GetHeroMMRRankings(
	ctx context.Context,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCGetHeroMMRRankingsResponse, error) {
	req := &protocol.CMsgClientToGCGetHeroMMRRankings{
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCGetHeroMMRRankingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankings),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankingsResponse),
		resp,
	)
}

// GetHideoutSDRTicket gets a hideout sdr ticket.
// Request ID: k_EMsgClientToGCGetHideoutSDRTicket
// Response ID: k_EMsgClientToGCGetHideoutSDRTicketResponse
// Request type: CMsgClientToGCGetHideoutSDRTicket
// Response type: CMsgClientToGCGetHideoutSDRTicketResponse
func (d *Deadlock) GetHideoutSDRTicket(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCGetHideoutSDRTicketResponse, error) {
	req := &protocol.CMsgClientToGCGetHideoutSDRTicket{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCGetHideoutSDRTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHideoutSDRTicket),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHideoutSDRTicketResponse),
		resp,
	)
}

// GetLeaderboard gets a leaderboard.
// Request ID: k_EMsgClientToGCGetLeaderboard
// Response ID: k_EMsgClientToGCGetLeaderboardResponse
// Request type: CMsgClientToGCGetLeaderboard
// Response type: CMsgClientToGCGetLeaderboardResponse
func (d *Deadlock) GetLeaderboard(
	ctx context.Context,
	leaderboardRegion protocol.ECitadelLeaderboardRegion,
	heroID uint32,
) (*protocol.CMsgClientToGCGetLeaderboardResponse, error) {
	req := &protocol.CMsgClientToGCGetLeaderboard{
		LeaderboardRegion: &leaderboardRegion,
		HeroId:            &heroID,
	}
	resp := &protocol.CMsgClientToGCGetLeaderboardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboard),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboardResponse),
		resp,
	)
}

// GetMatchHistory gets a match history.
// Request ID: k_EMsgClientToGCGetMatchHistory
// Response ID: k_EMsgClientToGCGetMatchHistoryResponse
// Request type: CMsgClientToGCGetMatchHistory
// Response type: CMsgClientToGCGetMatchHistoryResponse
func (d *Deadlock) GetMatchHistory(
	ctx context.Context,
	accountID uint32,
	continueCursor uint64,
	gameMode protocol.ECitadelGameMode,
	matchMode protocol.ECitadelMatchMode,
	rankedType protocol.ECitadelRankedType,
	rankInterval uint32,
) (*protocol.CMsgClientToGCGetMatchHistoryResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchHistory{
		AccountId:      &accountID,
		ContinueCursor: &continueCursor,
		GameMode:       &gameMode,
		MatchMode:      &matchMode,
		RankedType:     &rankedType,
		RankInterval:   &rankInterval,
	}
	resp := &protocol.CMsgClientToGCGetMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistory),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistoryResponse),
		resp,
	)
}

// GetMatchMetaData gets a match meta data.
// Request ID: k_EMsgClientToGCGetMatchMetaData
// Response ID: k_EMsgClientToGCGetMatchMetaDataResponse
// Request type: CMsgClientToGCGetMatchMetaData
// Response type: CMsgClientToGCGetMatchMetaDataResponse
func (d *Deadlock) GetMatchMetaData(
	ctx context.Context,
	matchID uint64,
	metadataSalt uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCGetMatchMetaDataResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchMetaData{
		MatchId:         &matchID,
		MetadataSalt:    &metadataSalt,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCGetMatchMetaDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaData),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaDataResponse),
		resp,
	)
}

// GetProfileCard gets a profile card.
// Request ID: k_EMsgClientToGCGetProfileCard
// Request type: CMsgClientToGCGetProfileCard
func (d *Deadlock) GetProfileCard(
	accountID uint32,
	devAccessHint bool,
	friendAccessHint bool,
) {
	req := &protocol.CMsgClientToGCGetProfileCard{
		AccountId:        &accountID,
		DevAccessHint:    &devAccessHint,
		FriendAccessHint: &friendAccessHint,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCard), req)
}

// GetRankData gets a rank data.
// Request ID: k_EMsgClientToGCGetRankData
// Response ID: k_EMsgGCToClientGetRankDataResponse
// Request type: CMsgClientToGCGetRankData
// Response type: CMsgGCToClientGetRankDataResponse
func (d *Deadlock) GetRankData(
	ctx context.Context,
) (*protocol.CMsgGCToClientGetRankDataResponse, error) {
	req := &protocol.CMsgClientToGCGetRankData{}
	resp := &protocol.CMsgGCToClientGetRankDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetRankData),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientGetRankDataResponse),
		resp,
	)
}

// GetSurveyQuestion gets a survey question.
// Request ID: k_EMsgClientToGCGetSurveyQuestion
// Response ID: k_EMsgClientToGCGetSurveyQuestionResponse
// Request type: CMsgClientToGCGetSurveyQuestion
// Response type: CMsgClientToGCGetSurveyQuestionResponse
func (d *Deadlock) GetSurveyQuestion(
	ctx context.Context,
	devForceNewQuestionType protocol.CMsgSurveyQuestion_EQuestionType,
) (*protocol.CMsgClientToGCGetSurveyQuestionResponse, error) {
	req := &protocol.CMsgClientToGCGetSurveyQuestion{
		DevForceNewQuestionType: &devForceNewQuestionType,
	}
	resp := &protocol.CMsgClientToGCGetSurveyQuestionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetSurveyQuestion),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetSurveyQuestionResponse),
		resp,
	)
}

// GrantForumAccess grants forum access.
// Request ID: k_EMsgClientToGCGrantForumAccess
// Response ID: k_EMsgClientToGCGrantForumAccessResponse
// Request type: CMsgClientToGCGrantForumAccess
// Response type: CMsgClientToGCGrantForumAccessResponse
func (d *Deadlock) GrantForumAccess(
	ctx context.Context,
	email string,
) (*protocol.CMsgClientToGCGrantForumAccessResponse, error) {
	req := &protocol.CMsgClientToGCGrantForumAccess{
		Email: &email,
	}
	resp := &protocol.CMsgClientToGCGrantForumAccessResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccess),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccessResponse),
		resp,
	)
}

// JoinParty joins a party.
// Request ID: k_EMsgClientToGCPartyJoin
// Response ID: k_EMsgClientToGCPartyJoinResponse
// Request type: CMsgClientToGCPartyJoin
// Response type: CMsgClientToGCPartyJoinResponse
func (d *Deadlock) JoinParty(
	ctx context.Context,
	partyID uint64,
	isRejoin bool,
	partyMmInfo protocol.CMsgPartyMMInfo,
) (*protocol.CMsgClientToGCPartyJoinResponse, error) {
	req := &protocol.CMsgClientToGCPartyJoin{
		PartyId:     &partyID,
		IsRejoin:    &isRejoin,
		PartyMmInfo: &partyMmInfo,
	}
	resp := &protocol.CMsgClientToGCPartyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoin),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinResponse),
		resp,
	)
}

// JoinPartyViaCode joins a party via code.
// Request ID: k_EMsgClientToGCPartyJoinViaCode
// Response ID: k_EMsgClientToGCPartyJoinViaCodeResponse
// Request type: CMsgClientToGCPartyJoinViaCode
// Response type: CMsgClientToGCPartyJoinViaCodeResponse
func (d *Deadlock) JoinPartyViaCode(
	ctx context.Context,
	joinCode uint64,
	partyMmInfo protocol.CMsgPartyMMInfo,
) (*protocol.CMsgClientToGCPartyJoinViaCodeResponse, error) {
	req := &protocol.CMsgClientToGCPartyJoinViaCode{
		JoinCode:    &joinCode,
		PartyMmInfo: &partyMmInfo,
	}
	resp := &protocol.CMsgClientToGCPartyJoinViaCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCode),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCodeResponse),
		resp,
	)
}

// LeaveLobby leaves a lobby.
// Request ID: k_EMsgClientToGCLeaveLobby
// Response ID: k_EMsgClientToGCLeaveLobbyResponse
// Request type: CMsgClientToGCLeaveLobby
// Response type: CMsgClientToGCLeaveLobbyResponse
func (d *Deadlock) LeaveLobby(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCLeaveLobbyResponse, error) {
	req := &protocol.CMsgClientToGCLeaveLobby{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCLeaveLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobby),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobbyResponse),
		resp,
	)
}

// LeaveParty leaves a party.
// Request ID: k_EMsgClientToGCPartyLeave
// Response ID: k_EMsgClientToGCPartyLeaveResponse
// Request type: CMsgClientToGCPartyLeave
// Response type: CMsgClientToGCPartyLeaveResponse
func (d *Deadlock) LeaveParty(
	ctx context.Context,
	partyID uint64,
) (*protocol.CMsgClientToGCPartyLeaveResponse, error) {
	req := &protocol.CMsgClientToGCPartyLeave{
		PartyId: &partyID,
	}
	resp := &protocol.CMsgClientToGCPartyLeaveResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyLeave),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyLeaveResponse),
		resp,
	)
}

// RecordClientEvents records client events.
// Request ID: k_EMsgClientToGCRecordClientEvents
// Response ID: k_EMsgClientToGCRecordClientEventsResponse
// Request type: CMsgClientToGCRecordClientEvents
// Response type: CMsgClientToGCRecordClientEventsResponse
func (d *Deadlock) RecordClientEvents(
	ctx context.Context,
	events []*protocol.CMsgClientToGCRecordClientEvents_Event,
	clientRunToken uint32,
) (*protocol.CMsgClientToGCRecordClientEventsResponse, error) {
	req := &protocol.CMsgClientToGCRecordClientEvents{
		Events:         events,
		ClientRunToken: &clientRunToken,
	}
	resp := &protocol.CMsgClientToGCRecordClientEventsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEvents),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEventsResponse),
		resp,
	)
}

// ReportPlayerFromMatch reports a player from match.
// Request ID: k_EMsgClientToGCReportPlayerFromMatch
// Response ID: k_EMsgClientToGCReportPlayerFromMatchResponse
// Request type: CMsgClientToGCReportPlayerFromMatch
// Response type: CMsgClientToGCReportPlayerFromMatchResponse
func (d *Deadlock) ReportPlayerFromMatch(
	ctx context.Context,
	matchID uint64,
	targetAccountID uint32,
	reportType protocol.CMsgClientToGCReportPlayerFromMatch_EReportType,
	reportText string,
) (*protocol.CMsgClientToGCReportPlayerFromMatchResponse, error) {
	req := &protocol.CMsgClientToGCReportPlayerFromMatch{
		MatchId:         &matchID,
		TargetAccountId: &targetAccountID,
		ReportType:      &reportType,
		ReportText:      &reportText,
	}
	resp := &protocol.CMsgClientToGCReportPlayerFromMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatch),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatchResponse),
		resp,
	)
}

// RequestDevCheatReports requests dev cheat reports.
// Request ID: k_EMsgClientToGCDevRequestCheatReports
// Response ID: k_EMsgClientToGCDevRequestCheatReportsResponse
// Request type: CMsgClientToGCRequestCheatReports
// Response type: CMsgClientToGCRequestCheatReportsResponse
func (d *Deadlock) RequestDevCheatReports(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestCheatReportsResponse, error) {
	req := &protocol.CMsgClientToGCRequestCheatReports{}
	resp := &protocol.CMsgClientToGCRequestCheatReportsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReports),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReportsResponse),
		resp,
	)
}

// RequestHeroReleaseVoteTally requests a hero release vote tally.
// Request ID: k_EMsgClientToGCRequestHeroReleaseVoteTally
// Request type: CMsgClientToGCRequestHeroReleaseVoteTally
func (d *Deadlock) RequestHeroReleaseVoteTally(
	voteRounds []uint32,
) {
	req := &protocol.CMsgClientToGCRequestHeroReleaseVoteTally{
		VoteRounds: voteRounds,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRequestHeroReleaseVoteTally), req)
}

// RequestModerator requests a moderator.
// Request ID: k_EMsgClientToGCModeratorRequest
// Response ID: k_EMsgClientToGCModeratorRequestResponse
// Request type: CMsgClientToGCModeratorRequest
// Response type: CMsgClientToGCModeratorRequestResponse
func (d *Deadlock) RequestModerator(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCModeratorRequestResponse, error) {
	req := &protocol.CMsgClientToGCModeratorRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCModeratorRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequest),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequestResponse),
		resp,
	)
}

// SendBookUnlock sends a book unlock.
// Request ID: k_EMsgClientToGCBookUnlock
// Response ID: k_EMsgClientToGCBookUnlockResponse
// Request type: CMsgClientToGCBookUnlock
// Response type: CMsgClientToGCBookUnlockResponse
func (d *Deadlock) SendBookUnlock(
	ctx context.Context,
	bookID uint32,
	unlockID uint32,
	expectedCost uint32,
) (*protocol.CMsgClientToGCBookUnlockResponse, error) {
	req := &protocol.CMsgClientToGCBookUnlock{
		BookId:       &bookID,
		UnlockId:     &unlockID,
		ExpectedCost: &expectedCost,
	}
	resp := &protocol.CMsgClientToGCBookUnlockResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlock),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlockResponse),
		resp,
	)
}

// SendChallengeActivate sends a challenge activate.
// Request ID: k_EMsgClientToGCChallengeActivate
// Response ID: k_EMsgClientToGCChallengeActivateResponse
// Request type: CMsgClientToGCChallengeActivate
// Response type: CMsgClientToGCChallengeActivateResponse
func (d *Deadlock) SendChallengeActivate(
	ctx context.Context,
	challengeID uint64,
) (*protocol.CMsgClientToGCChallengeActivateResponse, error) {
	req := &protocol.CMsgClientToGCChallengeActivate{
		ChallengeId: &challengeID,
	}
	resp := &protocol.CMsgClientToGCChallengeActivateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeActivate),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeActivateResponse),
		resp,
	)
}

// SendCommendPlayerFromMatch sends a commend player from match.
// Request ID: k_EMsgClientToGCCommendPlayerFromMatch
// Response ID: k_EMsgClientToGCCommendPlayerFromMatchResponse
// Request type: CMsgClientToGCCommendPlayerFromMatch
// Response type: CMsgClientToGCCommendPlayerFromMatchResponse
func (d *Deadlock) SendCommendPlayerFromMatch(
	ctx context.Context,
	matchID uint64,
	targetAccountID uint32,
	commendType protocol.ECommendType,
	fakeCommendHeroID uint32,
) (*protocol.CMsgClientToGCCommendPlayerFromMatchResponse, error) {
	req := &protocol.CMsgClientToGCCommendPlayerFromMatch{
		MatchId:           &matchID,
		TargetAccountId:   &targetAccountID,
		CommendType:       &commendType,
		FakeCommendHeroId: &fakeCommendHeroID,
	}
	resp := &protocol.CMsgClientToGCCommendPlayerFromMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatch),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatchResponse),
		resp,
	)
}

// SendDeleteHeroBuild sends a delete hero build.
// Request ID: k_EMsgClientToGCDeleteHeroBuild
// Response ID: k_EMsgClientToGCDeleteHeroBuildResponse
// Request type: CMsgClientToGCDeleteHeroBuild
// Response type: CMsgClientToGCDeleteHeroBuildResponse
func (d *Deadlock) SendDeleteHeroBuild(
	ctx context.Context,
	authorAccountID uint32,
	heroBuildID uint32,
) (*protocol.CMsgClientToGCDeleteHeroBuildResponse, error) {
	req := &protocol.CMsgClientToGCDeleteHeroBuild{
		AuthorAccountId: &authorAccountID,
		HeroBuildId:     &heroBuildID,
	}
	resp := &protocol.CMsgClientToGCDeleteHeroBuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuild),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuildResponse),
		resp,
	)
}

// SendDevAction sends a dev action.
// Request ID: k_EMsgClientToGCDevAction
// Response ID: k_EMsgClientToGCDevActionResponse
// Request type: CMsgClientToGCDevAction
// Response type: CMsgClientToGCDevActionResponse
func (d *Deadlock) SendDevAction(
	ctx context.Context,
	action protocol.CMsgClientToGCDevAction_EAction,
	accountID uint32,
	uintValue uint32,
	intValue int32,
	boolValue bool,
	strValue string,
	matchID uint64,
) (*protocol.CMsgClientToGCDevActionResponse, error) {
	req := &protocol.CMsgClientToGCDevAction{
		Action:    &action,
		AccountId: &accountID,
		UintValue: &uintValue,
		IntValue:  &intValue,
		BoolValue: &boolValue,
		StrValue:  &strValue,
		MatchId:   &matchID,
	}
	resp := &protocol.CMsgClientToGCDevActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevAction),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevActionResponse),
		resp,
	)
}

// SendIsInMatchmaking sends a is in matchmaking.
// Request ID: k_EMsgClientToGCIsInMatchmaking
// Response ID: k_EMsgClientToGCIsInMatchmakingResponse
// Request type: CMsgClientToGCIsInMatchmaking
// Response type: CMsgClientToGCIsInMatchmakingResponse
func (d *Deadlock) SendIsInMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCIsInMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCIsInMatchmaking{}
	resp := &protocol.CMsgClientToGCIsInMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmaking),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmakingResponse),
		resp,
	)
}

// SendModifyDevAnnouncements sends modify dev announcements.
// Request ID: k_EMsgClientToGCModifyDevAnnouncements
// Response ID: k_EMsgClientToGCModifyDevAnnouncementsResponse
// Request type: CMsgClientToGCModifyDevAnnouncements
// Response type: CMsgClientToGCModifyDevAnnouncementsResponse
func (d *Deadlock) SendModifyDevAnnouncements(
	ctx context.Context,
	operation protocol.CMsgClientToGCModifyDevAnnouncements_EOperation,
	targetID uint32,
	priority uint32,
	title string,
	message string,
	uRL string,
	patchVersion string,
) (*protocol.CMsgClientToGCModifyDevAnnouncementsResponse, error) {
	req := &protocol.CMsgClientToGCModifyDevAnnouncements{
		Operation:    &operation,
		TargetId:     &targetID,
		Priority:     &priority,
		Title:        &title,
		Message:      &message,
		Url:          &uRL,
		PatchVersion: &patchVersion,
	}
	resp := &protocol.CMsgClientToGCModifyDevAnnouncementsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCModifyDevAnnouncements),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCModifyDevAnnouncementsResponse),
		resp,
	)
}

// SendPartyAction sends a party action.
// Request ID: k_EMsgClientToGCPartyAction
// Response ID: k_EMsgClientToGCPartyActionResponse
// Request type: CMsgClientToGCPartyAction
// Response type: CMsgClientToGCPartyActionResponse
func (d *Deadlock) SendPartyAction(
	ctx context.Context,
	partyID uint64,
	targetAccountID uint32,
	actionID protocol.CMsgClientToGCPartyAction_EAction,
	uintValue uint64,
	boolValue bool,
	strValue string,
) (*protocol.CMsgClientToGCPartyActionResponse, error) {
	req := &protocol.CMsgClientToGCPartyAction{
		PartyId:         &partyID,
		TargetAccountId: &targetAccountID,
		ActionId:        &actionID,
		UintValue:       &uintValue,
		BoolValue:       &boolValue,
		StrValue:        &strValue,
	}
	resp := &protocol.CMsgClientToGCPartyActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyAction),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyActionResponse),
		resp,
	)
}

// SendPartyChatMsg sends a party chat msg.
// Request ID: k_EMsgClientToGCPartySendChatMsg
// Response ID: k_EMsgClientToGCPartySendChatMsgResponse
// Request type: CMsgClientToGCPartySendChatMsg
// Response type: CMsgClientToGCPartySendChatMsgResponse
func (d *Deadlock) SendPartyChatMsg(
	ctx context.Context,
	partyID uint64,
	chatMsg string,
) (*protocol.CMsgClientToGCPartySendChatMsgResponse, error) {
	req := &protocol.CMsgClientToGCPartySendChatMsg{
		PartyId: &partyID,
		ChatMsg: &chatMsg,
	}
	resp := &protocol.CMsgClientToGCPartySendChatMsgResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySendChatMsg),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySendChatMsgResponse),
		resp,
	)
}

// SendPartyInviteUser sends a party invite user.
// Request ID: k_EMsgClientToGCPartyInviteUser
// Response ID: k_EMsgClientToGCPartyInviteUserResponse
// Request type: CMsgClientToGCPartyInviteUser
// Response type: CMsgClientToGCPartyInviteUserResponse
func (d *Deadlock) SendPartyInviteUser(
	ctx context.Context,
	partyID uint64,
	inviteAccountID uint32,
) (*protocol.CMsgClientToGCPartyInviteUserResponse, error) {
	req := &protocol.CMsgClientToGCPartyInviteUser{
		PartyId:         &partyID,
		InviteAccountId: &inviteAccountID,
	}
	resp := &protocol.CMsgClientToGCPartyInviteUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUser),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUserResponse),
		resp,
	)
}

// SendReplacementSDRTicket sends a replacement sdr ticket.
// Request ID: k_EMsgClientToGCReplacementSDRTicket
// Response ID: k_EMsgClientToGCReplacementSDRTicketResponse
// Request type: CMsgClientToGCReplacementSDRTicket
// Response type: CMsgClientToGCReplacementSDRTicketResponse
func (d *Deadlock) SendReplacementSDRTicket(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCReplacementSDRTicketResponse, error) {
	req := &protocol.CMsgClientToGCReplacementSDRTicket{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCReplacementSDRTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicket),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicketResponse),
		resp,
	)
}

// SendUnlockHero sends a unlock hero.
// Request ID: k_EMsgClientToGCUnlockHero
// Response ID: k_EMsgClientToGCUnlockHeroResponse
// Request type: CMsgClientToGCUnlockHero
// Response type: CMsgClientToGCUnlockHeroResponse
func (d *Deadlock) SendUnlockHero(
	ctx context.Context,
	heroIDs []uint32,
	heroChoiceID uint32,
) (*protocol.CMsgClientToGCUnlockHeroResponse, error) {
	req := &protocol.CMsgClientToGCUnlockHero{
		HeroIds:      heroIDs,
		HeroChoiceId: &heroChoiceID,
	}
	resp := &protocol.CMsgClientToGCUnlockHeroResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHero),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHeroResponse),
		resp,
	)
}

// SendUpdateAccountSync sends a update account sync.
// Request ID: k_EMsgClientToGCUpdateAccountSync
// Response ID: k_EMsgClientToGCUpdateAccountSyncResponse
// Request type: CMsgClientToGCUpdateAccountSync
// Response type: CMsgClientToGCUpdateAccountSyncResponse
func (d *Deadlock) SendUpdateAccountSync(
	ctx context.Context,
	ids []uint32,
	values []uint32,
) (*protocol.CMsgClientToGCUpdateAccountSyncResponse, error) {
	req := &protocol.CMsgClientToGCUpdateAccountSync{
		Ids:    ids,
		Values: values,
	}
	resp := &protocol.CMsgClientToGCUpdateAccountSyncResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSync),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSyncResponse),
		resp,
	)
}

// SendUpdateHeroBuild sends a update hero build.
// Request ID: k_EMsgClientToGCUpdateHeroBuild
// Response ID: k_EMsgClientToGCUpdateHeroBuildResponse
// Request type: CMsgClientToGCUpdateHeroBuild
// Response type: CMsgClientToGCUpdateHeroBuildResponse
func (d *Deadlock) SendUpdateHeroBuild(
	ctx context.Context,
	heroBuild protocol.CMsgHeroBuild,
) (*protocol.CMsgClientToGCUpdateHeroBuildResponse, error) {
	req := &protocol.CMsgClientToGCUpdateHeroBuild{
		HeroBuild: &heroBuild,
	}
	resp := &protocol.CMsgClientToGCUpdateHeroBuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuild),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildResponse),
		resp,
	)
}

// SendUpdateHeroBuildPreference sends a update hero build preference.
// Request ID: k_EMsgClientToGCUpdateHeroBuildPreference
// Response ID: k_EMsgClientToGCUpdateHeroBuildPreferenceResponse
// Request type: CMsgClientToGCUpdateHeroBuildPreference
// Response type: CMsgClientToGCUpdateHeroBuildPreferenceResponse
func (d *Deadlock) SendUpdateHeroBuildPreference(
	ctx context.Context,
	heroBuildID uint32,
	preference protocol.CMsgHeroBuildPreference,
) (*protocol.CMsgClientToGCUpdateHeroBuildPreferenceResponse, error) {
	req := &protocol.CMsgClientToGCUpdateHeroBuildPreference{
		HeroBuildId: &heroBuildID,
		Preference:  &preference,
	}
	resp := &protocol.CMsgClientToGCUpdateHeroBuildPreferenceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreference),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreferenceResponse),
		resp,
	)
}

// SendUpdateRoster sends a update roster.
// Request ID: k_EMsgClientToGCUpdateRoster
// Response ID: k_EMsgClientToGCUpdateRosterResponse
// Request type: CMsgClientToGCUpdateRoster
// Response type: CMsgClientToGCUpdateRosterResponse
func (d *Deadlock) SendUpdateRoster(
	ctx context.Context,
	heroes protocol.CMsgHeroSelectionMatchInfo,
	gameMode protocol.ECitadelGameMode,
	matchMode protocol.ECitadelMatchMode,
) (*protocol.CMsgClientToGCUpdateRosterResponse, error) {
	req := &protocol.CMsgClientToGCUpdateRoster{
		Heroes:    &heroes,
		GameMode:  &gameMode,
		MatchMode: &matchMode,
	}
	resp := &protocol.CMsgClientToGCUpdateRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRoster),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRosterResponse),
		resp,
	)
}

// SendUpdateSpectatorStatus sends update spectator status.
// Request ID: k_EMsgClientToGCUpdateSpectatorStatus
// Request type: CMsgClientToGCUpdateSpectatorStatus
func (d *Deadlock) SendUpdateSpectatorStatus(
	spectatingLobbyID uint64,
	stoppedSpectating bool,
) {
	req := &protocol.CMsgClientToGCUpdateSpectatorStatus{
		SpectatingLobbyId: &spectatingLobbyID,
		StoppedSpectating: &stoppedSpectating,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateSpectatorStatus), req)
}

// SetAccountPrivacySetting sets a account privacy setting.
// Request ID: k_EMsgClientToGCSetAccountPrivacySetting
// Response ID: k_EMsgClientToGCSetAccountPrivacySettingResponse
// Request type: CMsgClientToGCSetAccountPrivacySetting
// Response type: CMsgClientToGCSetAccountPrivacySettingResponse
func (d *Deadlock) SetAccountPrivacySetting(
	ctx context.Context,
	privacySetting protocol.CSOGameAccountClient_EPrivacySetting,
) (*protocol.CMsgClientToGCSetAccountPrivacySettingResponse, error) {
	req := &protocol.CMsgClientToGCSetAccountPrivacySetting{
		PrivacySetting: &privacySetting,
	}
	resp := &protocol.CMsgClientToGCSetAccountPrivacySettingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetAccountPrivacySetting),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetAccountPrivacySettingResponse),
		resp,
	)
}

// SetChallengeGenerate sets a challenge generate.
// Request ID: k_EMsgClientToGCGenerateChallengeSet
// Response ID: k_EMsgClientToGCGenerateChallengeSetResponse
// Request type: CMsgClientToGCGenerateChallengeSet
// Response type: CMsgClientToGCGenerateChallengeSetResponse
func (d *Deadlock) SetChallengeGenerate(
	ctx context.Context,
) (*protocol.CMsgClientToGCGenerateChallengeSetResponse, error) {
	req := &protocol.CMsgClientToGCGenerateChallengeSet{}
	resp := &protocol.CMsgClientToGCGenerateChallengeSetResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGenerateChallengeSet),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGenerateChallengeSetResponse),
		resp,
	)
}

// SetDevMMBias sets dev mm bias.
// Request ID: k_EMsgClientToGCDevSetMMBias
// Request type: CMsgClientToGCDevSetMMBias
func (d *Deadlock) SetDevMMBias(
	accountID uint32,
	value uint32,
) {
	req := &protocol.CMsgClientToGCDevSetMMBias{
		AccountId: &accountID,
		Value:     &value,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevSetMMBias), req)
}

// SetNewPlayerProgress sets new player progress.
// Request ID: k_EMsgClientToGCSetNewPlayerProgress
// Response ID: k_EMsgClientToGCSetNewPlayerProgressResponse
// Request type: CMsgClientToGCSetNewPlayerProgress
// Response type: CMsgClientToGCSetNewPlayerProgressResponse
func (d *Deadlock) SetNewPlayerProgress(
	ctx context.Context,
	flag protocol.ECitadelNewPlayerProgressFlag,
) (*protocol.CMsgClientToGCSetNewPlayerProgressResponse, error) {
	req := &protocol.CMsgClientToGCSetNewPlayerProgress{
		Flag: &flag,
	}
	resp := &protocol.CMsgClientToGCSetNewPlayerProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgress),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgressResponse),
		resp,
	)
}

// SetPartyMode sets a party mode.
// Request ID: k_EMsgClientToGCPartySetMode
// Response ID: k_EMsgClientToGCPartySetModeResponse
// Request type: CMsgClientToGCPartySetMode
// Response type: CMsgClientToGCPartySetModeResponse
func (d *Deadlock) SetPartyMode(
	ctx context.Context,
	partyID uint64,
	matchMode protocol.ECitadelMatchMode,
	gameMode protocol.ECitadelGameMode,
	botDifficulty protocol.ECitadelBotDifficulty,
	devServerCommand string,
	regionMode protocol.ECitadelRegionMode,
) (*protocol.CMsgClientToGCPartySetModeResponse, error) {
	req := &protocol.CMsgClientToGCPartySetMode{
		PartyId:          &partyID,
		MatchMode:        &matchMode,
		GameMode:         &gameMode,
		BotDifficulty:    &botDifficulty,
		DevServerCommand: &devServerCommand,
		RegionMode:       &regionMode,
	}
	resp := &protocol.CMsgClientToGCPartySetModeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetMode),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetModeResponse),
		resp,
	)
}

// SetPartyReadyState sets a party ready state.
// Request ID: k_EMsgClientToGCPartySetReadyState
// Response ID: k_EMsgClientToGCPartySetReadyStateResponse
// Request type: CMsgClientToGCPartySetReadyState
// Response type: CMsgClientToGCPartySetReadyStateResponse
func (d *Deadlock) SetPartyReadyState(
	ctx context.Context,
	partyID uint64,
	readyState bool,
	heroRoster protocol.CMsgHeroSelectionMatchInfo,
) (*protocol.CMsgClientToGCPartySetReadyStateResponse, error) {
	req := &protocol.CMsgClientToGCPartySetReadyState{
		PartyId:    &partyID,
		ReadyState: &readyState,
		HeroRoster: &heroRoster,
	}
	resp := &protocol.CMsgClientToGCPartySetReadyStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyState),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyStateResponse),
		resp,
	)
}

// SetServerConVar sets a server con var.
// Request ID: k_EMsgClientToGCSetServerConVar
// Response ID: k_EMsgClientToGCSetServerConVarResponse
// Request type: CMsgClientToGCSetServerConVar
// Response type: CMsgClientToGCSetServerConVarResponse
func (d *Deadlock) SetServerConVar(
	ctx context.Context,
	convarName string,
	convarValue string,
	lobbyID uint64,
) (*protocol.CMsgClientToGCSetServerConVarResponse, error) {
	req := &protocol.CMsgClientToGCSetServerConVar{
		ConvarName:  &convarName,
		ConvarValue: &convarValue,
		LobbyId:     &lobbyID,
	}
	resp := &protocol.CMsgClientToGCSetServerConVarResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVar),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVarResponse),
		resp,
	)
}

// SpectateLobby spectates a lobby.
// Request ID: k_EMsgClientToGCSpectateLobby
// Response ID: k_EMsgClientToGCSpectateLobbyResponse
// Request type: CMsgClientToGCSpectateLobby
// Response type: CMsgClientToGCSpectateLobbyResponse
func (d *Deadlock) SpectateLobby(
	ctx context.Context,
	lobbyID uint64,
	clientPlatform protocol.EGCPlatform,
	matchID uint64,
) (*protocol.CMsgClientToGCSpectateLobbyResponse, error) {
	req := &protocol.CMsgClientToGCSpectateLobby{
		LobbyId:        &lobbyID,
		ClientPlatform: &clientPlatform,
		MatchId:        &matchID,
	}
	resp := &protocol.CMsgClientToGCSpectateLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobby),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobbyResponse),
		resp,
	)
}

// SpectateUser spectates a user.
// Request ID: k_EMsgClientToGCSpectateUser
// Response ID: k_EMsgClientToGCSpectateUserResponse
// Request type: CMsgClientToGCSpectateUser
// Response type: CMsgClientToGCSpectateUserResponse
func (d *Deadlock) SpectateUser(
	ctx context.Context,
	spectateAccountID uint32,
	clientPlatform protocol.EGCPlatform,
) (*protocol.CMsgClientToGCSpectateUserResponse, error) {
	req := &protocol.CMsgClientToGCSpectateUser{
		SpectateAccountId: &spectateAccountID,
		ClientPlatform:    &clientPlatform,
	}
	resp := &protocol.CMsgClientToGCSpectateUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUser),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUserResponse),
		resp,
	)
}

// StartMatchmaking starts a matchmaking.
// Request ID: k_EMsgClientToGCStartMatchmaking
// Response ID: k_EMsgClientToGCStartMatchmakingResponse
// Request type: CMsgClientToGCStartMatchmaking
// Response type: CMsgClientToGCStartMatchmakingResponse
func (d *Deadlock) StartMatchmaking(
	ctx context.Context,
	clientPlatform protocol.EGCPlatform,
	matchInfo protocol.CMsgStartFindingMatchInfo,
	pingTimes protocol.CMsgRegionPingTimesClient,
	heroes protocol.CMsgHeroSelectionMatchInfo,
	pgiVerified bool,
) (*protocol.CMsgClientToGCStartMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStartMatchmaking{
		ClientPlatform: &clientPlatform,
		MatchInfo:      &matchInfo,
		PingTimes:      &pingTimes,
		Heroes:         &heroes,
		PgiVerified:    &pgiVerified,
	}
	resp := &protocol.CMsgClientToGCStartMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmaking),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmakingResponse),
		resp,
	)
}

// StartPartyMatch starts a party match.
// Request ID: k_EMsgClientToGCPartyStartMatch
// Response ID: k_EMsgClientToGCPartyStartMatchResponse
// Request type: CMsgClientToGCPartyStartMatch
// Response type: CMsgClientToGCPartyStartMatchResponse
func (d *Deadlock) StartPartyMatch(
	ctx context.Context,
	partyID uint64,
) (*protocol.CMsgClientToGCPartyStartMatchResponse, error) {
	req := &protocol.CMsgClientToGCPartyStartMatch{
		PartyId: &partyID,
	}
	resp := &protocol.CMsgClientToGCPartyStartMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatch),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatchResponse),
		resp,
	)
}

// StartRankedInterval starts a ranked interval.
// Request ID: k_EMsgClientToGCStartRankedInterval
// Response ID: k_EMsgClientToGCStartRankedIntervalResponse
// Request type: CMsgClientToGCStartRankedInterval
// Response type: CMsgClientToGCStartRankedIntervalResponse
func (d *Deadlock) StartRankedInterval(
	ctx context.Context,
	rankType protocol.ECitadelRankedType,
	interval uint32,
) (*protocol.CMsgClientToGCStartRankedIntervalResponse, error) {
	req := &protocol.CMsgClientToGCStartRankedInterval{
		RankType: &rankType,
		Interval: &interval,
	}
	resp := &protocol.CMsgClientToGCStartRankedIntervalResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartRankedInterval),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartRankedIntervalResponse),
		resp,
	)
}

// StopMatchmaking stops a matchmaking.
// Request ID: k_EMsgClientToGCStopMatchmaking
// Response ID: k_EMsgClientToGCStopMatchmakingResponse
// Request type: CMsgClientToGCStopMatchmaking
// Response type: CMsgClientToGCStopMatchmakingResponse
func (d *Deadlock) StopMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCStopMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStopMatchmaking{}
	resp := &protocol.CMsgClientToGCStopMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmaking),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmakingResponse),
		resp,
	)
}

// SubmitHeroReleaseVote submits a hero release vote.
// Request ID: k_EMsgClientToGCSubmitHeroReleaseVote
// Request type: CMsgClientToGCSubmitHeroReleaseVote
func (d *Deadlock) SubmitHeroReleaseVote(
	voteRound uint32,
	voteOption uint32,
) {
	req := &protocol.CMsgClientToGCSubmitHeroReleaseVote{
		VoteRound:  &voteRound,
		VoteOption: &voteOption,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitHeroReleaseVote), req)
}

// SubmitPlaytestUser submits a playtest user.
// Request ID: k_EMsgClientToGCSubmitPlaytestUser
// Response ID: k_EMsgClientToGCSubmitPlaytestUserResponse
// Request type: CMsgClientToGCSubmitPlaytestUser
// Response type: CMsgClientToGCSubmitPlaytestUserResponse
func (d *Deadlock) SubmitPlaytestUser(
	ctx context.Context,
	location string,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCSubmitPlaytestUserResponse, error) {
	req := &protocol.CMsgClientToGCSubmitPlaytestUser{
		Location:        &location,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCSubmitPlaytestUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUser),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUserResponse),
		resp,
	)
}

// SubmitSurvey submits a survey.
// Request ID: k_EMsgClientToGCSubmitSurvey
// Response ID: k_EMsgClientToGCSubmitSurveyResponse
// Request type: CMsgClientToGCSubmitSurvey
// Response type: CMsgClientToGCSubmitSurveyResponse
func (d *Deadlock) SubmitSurvey(
	ctx context.Context,
	questionID uint32,
	isSkip bool,
	responseValue uint32,
	devForceNewQuestionType protocol.CMsgSurveyQuestion_EQuestionType,
) (*protocol.CMsgClientToGCSubmitSurveyResponse, error) {
	req := &protocol.CMsgClientToGCSubmitSurvey{
		QuestionId:              &questionID,
		IsSkip:                  &isSkip,
		ResponseValue:           &responseValue,
		DevForceNewQuestionType: &devForceNewQuestionType,
	}
	resp := &protocol.CMsgClientToGCSubmitSurveyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitSurvey),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitSurveyResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *Deadlock) registerGeneratedHandlers() {
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientBookUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.BookUpdated{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientCanRejoinParty)] = d.getEventEmitter(func() events.Event {
		return &events.CanRejoinParty{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientCommendNotification)] = d.getEventEmitter(func() events.Event {
		return &events.CommendNotification{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientDevAnnouncements)] = d.getEventEmitter(func() events.Event {
		return &events.DevAnnouncements{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientDevPlaytestStatus)] = d.getEventEmitter(func() events.Event {
		return &events.DevPlaytestStatus{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientMatchmakingStopped)] = d.getEventEmitter(func() events.Event {
		return &events.MatchmakingStopped{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyChatMsg)] = d.getEventEmitter(func() events.Event {
		return &events.PartyChatMsg{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyEvent)] = d.getEventEmitter(func() events.Event {
		return &events.PartyEvent{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ProfileCardUpdated{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientRankDataUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.RankDataUpdate{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientSDRTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SDRTicket{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientUpdateHeroReleaseVoteTally)] = d.getEventEmitter(func() events.Event {
		return &events.UpdateHeroReleaseVoteTally{}
	})
}
