package kms

import "github.com/dongecs/maplestory-openapi/api/tms"

// Union payloads (plus champion info).
type (
	Union                        = tms.Union
	UnionRaider                  = tms.UnionRaider
	UnionRaiderInnerStat         = tms.UnionRaiderInnerStat
	UnionRaiderBlock             = tms.UnionRaiderBlock
	UnionRaiderBlockControlPoint = tms.UnionRaiderBlockControlPoint
	UnionRaiderBlockPosition     = tms.UnionRaiderBlockPosition
	UnionRaiderPreset            = tms.UnionRaiderPreset
	UnionArtifact                = tms.UnionArtifact
	UnionArtifactEffect          = tms.UnionArtifactEffect
	UnionArtifactCrystal         = tms.UnionArtifactCrystal
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
