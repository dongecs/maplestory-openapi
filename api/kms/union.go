package kms

import "github.com/dongecs/maplestory-openapi/api/common"

// Union payloads (plus champion info).
type (
	Union                        = common.Union
	UnionRaider                  = common.UnionRaider
	UnionRaiderInnerStat         = common.UnionRaiderInnerStat
	UnionRaiderBlock             = common.UnionRaiderBlock
	UnionRaiderBlockControlPoint = common.UnionRaiderBlockControlPoint
	UnionRaiderBlockPosition     = common.UnionRaiderBlockPosition
	UnionRaiderPreset            = common.UnionRaiderPreset
	UnionArtifact                = common.UnionArtifact
	UnionArtifactEffect          = common.UnionArtifactEffect
	UnionArtifactCrystal         = common.UnionArtifactCrystal
)

type UnionChampion struct {
	Date                   string                   `json:"date"`
	UnionChampion          []UnionChampionInfo      `json:"union_champion"`
	ChampionBadgeTotalInfo []UnionChampionBadgeInfo `json:"champion_badge_total_info"`
}

type UnionChampionInfo struct {
	ChampionName      string                   `json:"champion_name"`
	ChampionSlot      int                      `json:"champion_slot"`
	ChampionGrade     string                   `json:"champion_grade"`
	ChampionClass     string                   `json:"champion_class"`
	ChampionBadgeInfo []UnionChampionBadgeInfo `json:"champion_badge_info"`
}

type UnionChampionBadgeInfo struct {
	Stat string `json:"stat"`
}
