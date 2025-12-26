package common

import "encoding/json"

// RankingResponse keeps ranking list raw to avoid duplicating every DTO.
type RankingResponse struct {
	Ranking json.RawMessage `json:"ranking"`
}

type OverallRankingResponse = RankingResponse
type UnionRankingResponse = RankingResponse
type GuildRankingResponse = RankingResponse
type DojangRankingResponse = RankingResponse
type SeedRankingResponse = RankingResponse
type AchievementRankingResponse = RankingResponse

// RankingFilter controls the ranking query filters used by KMS.
type RankingFilter struct {
	WorldName      string
	WorldType      string
	CharacterClass string
	OCID           string
	Page           string
	Region         string
}
