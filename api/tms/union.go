package tms

import "github.com/dongecs/maplestory-openapi/api/common"

// Union payloads shared across regions; re-exported for the TMS package.
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
	UnionChampion                = common.UnionChampion
	UnionChampionInfo            = common.UnionChampionInfo
	UnionChampionBadgeInfo       = common.UnionChampionBadgeInfo
)
