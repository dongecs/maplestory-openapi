package kms

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
