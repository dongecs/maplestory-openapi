package common

// Union represents base union info.
type Union struct {
	Date               *string `json:"date"`
	UnionLevel         *int    `json:"union_level"`
	UnionGrade         *string `json:"union_grade"`
	UnionArtifactLevel *int    `json:"union_artifact_level"`
	UnionArtifactExp   *int    `json:"union_artifact_exp"`
	UnionArtifactPoint *int    `json:"union_artifact_point"`
}

// UnionRaider represents union raider configurations.
type UnionRaider struct {
	Date               *string                `json:"date"`
	UnionRaiderStat    []string               `json:"union_raider_stat"`
	UnionOccupiedStat  []string               `json:"union_occupied_stat"`
	UnionInnerStat     []UnionRaiderInnerStat `json:"union_inner_stat"`
	UnionBlock         []UnionRaiderBlock     `json:"union_block"`
	UsePresetNo        int                    `json:"use_preset_no"`
	UnionRaiderPreset1 *UnionRaiderPreset     `json:"union_raider_preset_1"`
	UnionRaiderPreset2 *UnionRaiderPreset     `json:"union_raider_preset_2"`
	UnionRaiderPreset3 *UnionRaiderPreset     `json:"union_raider_preset_3"`
	UnionRaiderPreset4 *UnionRaiderPreset     `json:"union_raider_preset_4"`
	UnionRaiderPreset5 *UnionRaiderPreset     `json:"union_raider_preset_5"`
}

type UnionRaiderInnerStat struct {
	StatFieldID     string `json:"stat_field_id"`
	StatFieldEffect string `json:"stat_field_effect"`
}

type UnionRaiderBlock struct {
	BlockType         string                       `json:"block_type"`
	BlockClass        *string                      `json:"block_class"`
	BlockLevel        string                       `json:"block_level"`
	BlockControlPoint UnionRaiderBlockControlPoint `json:"block_control_point"`
	BlockPosition     []UnionRaiderBlockPosition   `json:"block_position"`
}

type UnionRaiderBlockControlPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type UnionRaiderBlockPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type UnionRaiderPreset struct {
	UnionRaiderStat   []string               `json:"union_raider_stat"`
	UnionOccupiedStat []string               `json:"union_occupied_stat"`
	UnionInnerStat    []UnionRaiderInnerStat `json:"union_inner_stat"`
	UnionBlock        []UnionRaiderBlock     `json:"union_block"`
}

// UnionArtifact represents artifact info.
type UnionArtifact struct {
	Date                  *string                `json:"date"`
	UnionArtifactEffect   []UnionArtifactEffect  `json:"union_artifact_effect"`
	UnionArtifactCrystal  []UnionArtifactCrystal `json:"union_artifact_crystal"`
	UnionArtifactRemainAP *int                   `json:"union_artifact_remain_ap"`
}

type UnionArtifactEffect struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type UnionArtifactCrystal struct {
	Name               string `json:"name"`
	ValidityFlag       string `json:"validity_flag"`
	DateExpire         string `json:"date_expire"`
	Level              int    `json:"level"`
	CrystalOptionName1 string `json:"crystal_option_name_1"`
	CrystalOptionName2 string `json:"crystal_option_name_2"`
	CrystalOptionName3 string `json:"crystal_option_name_3"`
}

// UnionChampion represents union champion info.
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
