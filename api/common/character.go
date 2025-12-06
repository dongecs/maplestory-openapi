package common

// Character identifies a character by ocid.
type Character struct {
	OCID string `json:"ocid"`
}

// CharacterBasic represents basic character information.
type CharacterBasic struct {
	Date                 *string `json:"date"`
	CharacterName        string  `json:"character_name"`
	WorldName            string  `json:"world_name"`
	CharacterGender      string  `json:"character_gender"`
	CharacterClass       string  `json:"character_class"`
	CharacterClassLevel  string  `json:"character_class_level"`
	CharacterLevel       int     `json:"character_level"`
	CharacterExp         int     `json:"character_exp"`
	CharacterExpRate     string  `json:"character_exp_rate"`
	CharacterGuildName   *string `json:"character_guild_name"`
	CharacterImage       string  `json:"character_image"`
	CharacterDateCreate  *string `json:"character_date_create"`
	AccessFlag           string  `json:"access_flag"`
	LiberationQuestClear string  `json:"liberation_quest_clear"`
	LiberationQuestClearFlag string  `json:"liberation_quest_clear_flag"`
}

// CharacterImage describes the rendered character image variants.
type CharacterImage struct {
	Date         *string
	OriginURL    string
	OriginImage  string
	Image        string
	Action       string
	Emotion      string
	Wmotion      string
	ActionFrame  int
	EmotionFrame int
	Width        int
	Height       int
	X            *int
	Y            *int
}

// CharacterPopularity represents popularity information.
type CharacterPopularity struct {
	Date       *string `json:"date"`
	Popularity *int    `json:"popularity"`
}

// CharacterStat represents overall stat information.
type CharacterStat struct {
	Date           *string              `json:"date"`
	CharacterClass *string              `json:"character_class"`
	FinalStat      []CharacterFinalStat `json:"final_stat"`
	RemainAP       *int                 `json:"remain_ap"`
}

type CharacterFinalStat struct {
	StatName  string `json:"stat_name"`
	StatValue string `json:"stat_value"`
}

// CharacterHyperStat represents hyper stat presets.
type CharacterHyperStat struct {
	Date                        *string                    `json:"date"`
	CharacterClass              *string                    `json:"character_class"`
	UsePresetNo                 *string                    `json:"use_preset_no"`
	UseAvailableHyperStat       *int                       `json:"use_available_hyper_stat"`
	HyperStatPreset1            []CharacterHyperStatPreset `json:"hyper_stat_preset_1"`
	HyperStatPreset1RemainPoint *int                       `json:"hyper_stat_preset_1_remain_point"`
	HyperStatPreset2            []CharacterHyperStatPreset `json:"hyper_stat_preset_2"`
	HyperStatPreset2RemainPoint *int                       `json:"hyper_stat_preset_2_remain_point"`
	HyperStatPreset3            []CharacterHyperStatPreset `json:"hyper_stat_preset_3"`
	HyperStatPreset3RemainPoint *int                       `json:"hyper_stat_preset_3_remain_point"`
}

type CharacterHyperStatPreset struct {
	StatType     string  `json:"stat_type"`
	StatPoint    *int    `json:"stat_point"`
	StatLevel    int     `json:"stat_level"`
	StatIncrease *string `json:"stat_increase"`
}

// CharacterPropensity represents personality stats.
type CharacterPropensity struct {
	Date             *string `json:"date"`
	CharismaLevel    *int    `json:"charisma_level"`
	SensibilityLevel *int    `json:"sensibility_level"`
	InsightLevel     *int    `json:"insight_level"`
	WillingnessLevel *int    `json:"willingness_level"`
	HandicraftLevel  *int    `json:"handicraft_level"`
	CharmLevel       *int    `json:"charm_level"`
}

// CharacterAbility represents ability information with presets.
type CharacterAbility struct {
	Date           *string                 `json:"date"`
	AbilityGrade   *string                 `json:"ability_grade"`
	AbilityInfo    []CharacterAbilityInfo  `json:"ability_info"`
	RemainFame     *int                    `json:"remain_fame"`
	PresetNo       *int                    `json:"preset_no"`
	AbilityPreset1 *CharacterAbilityPreset `json:"ability_preset_1"`
	AbilityPreset2 *CharacterAbilityPreset `json:"ability_preset_2"`
	AbilityPreset3 *CharacterAbilityPreset `json:"ability_preset_3"`
}

type CharacterAbilityInfo struct {
	AbilityNo    string `json:"ability_no"`
	AbilityGrade string `json:"ability_grade"`
	AbilityValue string `json:"ability_value"`
}

type CharacterAbilityPreset struct {
	AbilityPresetGrade string                 `json:"ability_preset_grade"`
	AbilityInfo        []CharacterAbilityInfo `json:"ability_info"`
}

// CharacterItemEquipment represents equipped items and presets.
type CharacterItemEquipment struct {
	Date                 *string                          `json:"date"`
	CharacterGender      *string                          `json:"character_gender"`
	CharacterClass       *string                          `json:"character_class"`
	PresetNo             *int                             `json:"preset_no"`
	ItemEquipment        []CharacterItemEquipmentInfo     `json:"item_equipment"`
	ItemEquipmentPreset1 []CharacterItemEquipmentInfo     `json:"item_equipment_preset_1"`
	ItemEquipmentPreset2 []CharacterItemEquipmentInfo     `json:"item_equipment_preset_2"`
	ItemEquipmentPreset3 []CharacterItemEquipmentInfo     `json:"item_equipment_preset_3"`
	Title                *CharacterItemEquipmentTitle     `json:"title"`
	DragonEquipment      []CharacterItemEquipmentMechanic `json:"dragon_equipment"`
	MechanicEquipment    []CharacterItemEquipmentMechanic `json:"mechanic_equipment"`
}

type CharacterItemEquipmentInfo struct {
	ItemEquipmentPart              string                                  `json:"item_equipment_part"`
	ItemEquipmentSlot              string                                  `json:"item_equipment_slot"`
	ItemName                       string                                  `json:"item_name"`
	ItemIcon                       string                                  `json:"item_icon"`
	ItemDescription                *string                                 `json:"item_description"`
	ItemShapeName                  string                                  `json:"item_shape_name"`
	ItemShapeIcon                  string                                  `json:"item_shape_icon"`
	ItemGender                     *string                                 `json:"item_gender"`
	ItemTotalOption                CharacterItemEquipmentTotalOption       `json:"item_total_option"`
	ItemBaseOption                 CharacterItemEquipmentBaseOption        `json:"item_base_option"`
	PotentialOptionFlag            *string                                 `json:"potential_option_flag"`
	AdditionalPotentialOptionFlag  *string                                 `json:"additional_potential_option_flag"`
	PotentialOptionGrade           *string                                 `json:"potential_option_grade"`
	AdditionalPotentialOptionGrade *string                                 `json:"additional_potential_option_grade"`
	PotentialOption1               *string                                 `json:"potential_option_1"`
	PotentialOption2               *string                                 `json:"potential_option_2"`
	PotentialOption3               *string                                 `json:"potential_option_3"`
	AdditionalPotentialOption1     *string                                 `json:"additional_potential_option_1"`
	AdditionalPotentialOption2     *string                                 `json:"additional_potential_option_2"`
	AdditionalPotentialOption3     *string                                 `json:"additional_potential_option_3"`
	EquipmentLevelIncrease         int                                     `json:"equipment_level_increase"`
	ItemExceptionalOption          CharacterItemEquipmentExceptionalOption `json:"item_exceptional_option"`
	ItemAddOption                  CharacterItemEquipmentAddOption         `json:"item_add_option"`
	GrowthExp                      int                                     `json:"growth_exp"`
	GrowthLevel                    int                                     `json:"growth_level"`
	ScrollUpgrade                  string                                  `json:"scroll_upgrade"`
	CuttableCount                  string                                  `json:"cuttable_count"`
	GoldenHammerFlag               string                                  `json:"golden_hammer_flag"`
	ScrollResilienceCount          string                                  `json:"scroll_resilience_count"`
	ScrollUpgradeableCount         string                                  `json:"scroll_upgradeable_count"`
	SoulName                       *string                                 `json:"soul_name"`
	SoulOption                     *string                                 `json:"soul_option"`
	ItemEtcOption                  CharacterItemEquipmentEtcOption         `json:"item_etc_option"`
	Starforce                      string                                  `json:"starforce"`
	StarforceScrollFlag            string                                  `json:"starforce_scroll_flag"`
	ItemStarforceOption            CharacterItemEquipmentStarforceOption   `json:"item_starforce_option"`
	SpecialRingLevel               int                                     `json:"special_ring_level"`
	DateExpire                     *string                                 `json:"date_expire"`
}

type CharacterItemEquipmentMechanic struct {
	ItemEquipmentPart      string                                  `json:"item_equipment_part"`
	ItemEquipmentSlot      string                                  `json:"item_equipment_slot"`
	ItemName               string                                  `json:"item_name"`
	ItemIcon               string                                  `json:"item_icon"`
	ItemDescription        *string                                 `json:"item_description"`
	ItemShapeName          string                                  `json:"item_shape_name"`
	ItemShapeIcon          string                                  `json:"item_shape_icon"`
	ItemGender             *string                                 `json:"item_gender"`
	ItemTotalOption        CharacterItemEquipmentOption            `json:"item_total_option"`
	ItemBaseOption         CharacterItemEquipmentBaseOption        `json:"item_base_option"`
	EquipmentLevelIncrease int                                     `json:"equipment_level_increase"`
	ItemExceptionalOption  CharacterItemEquipmentExceptionalOption `json:"item_exceptional_option"`
	ItemAddOption          CharacterItemEquipmentAddOption         `json:"item_add_option"`
	GrowthExp              int                                     `json:"growth_exp"`
	GrowthLevel            int                                     `json:"growth_level"`
	ScrollUpgrade          string                                  `json:"scroll_upgrade"`
	CuttableCount          string                                  `json:"cuttable_count"`
	GoldenHammerFlag       string                                  `json:"golden_hammer_flag"`
	ScrollResilienceCount  string                                  `json:"scroll_resilience_count"`
	ScrollUpgradeableCount string                                  `json:"scroll_upgradeable_count"`
	SoulName               *string                                 `json:"soul_name"`
	SoulOption             *string                                 `json:"soul_option"`
	ItemEtcOption          CharacterItemEquipmentOption            `json:"item_etc_option"`
	Starforce              string                                  `json:"starforce"`
	StarforceScrollFlag    string                                  `json:"starforce_scroll_flag"`
	ItemStarforceOption    CharacterItemEquipmentOption            `json:"item_starforce_option"`
	SpecialRingLevel       int                                     `json:"special_ring_level"`
	DateExpire             *string                                 `json:"date_expire"`
}

type CharacterItemEquipmentOption struct {
	Str                    string `json:"str"`
	Dex                    string `json:"dex"`
	Int                    string `json:"int"`
	Luk                    string `json:"luk"`
	MaxHP                  string `json:"max_hp"`
	MaxMP                  string `json:"max_mp"`
	AttackPower            string `json:"attack_power"`
	MagicPower             string `json:"magic_power"`
	Armor                  string `json:"armor"`
	Speed                  string `json:"speed"`
	Jump                   string `json:"jump"`
	BossDamage             string `json:"boss_damage"`
	IgnoreMonsterArmor     string `json:"ignore_monster_armor"`
	AllStat                string `json:"all_stat"`
	Damage                 string `json:"damage"`
	EquipmentLevelDecrease int    `json:"equipment_level_decrease"`
	MaxHPRate              string `json:"max_hp_rate"`
	MaxMPRate              string `json:"max_mp_rate"`
}

type CharacterItemEquipmentExceptionalOption struct {
	Str                string `json:"str"`
	Dex                string `json:"dex"`
	Int                string `json:"int"`
	Luk                string `json:"luk"`
	MaxHP              string `json:"max_hp"`
	MaxMP              string `json:"max_mp"`
	AttackPower        string `json:"attack_power"`
	MagicPower         string `json:"magic_power"`
	ExceptionalUpgrade *int   `json:"exceptional_upgrade"`
}

type CharacterItemEquipmentTotalOption struct {
	Str                    string `json:"str"`
	Dex                    string `json:"dex"`
	Int                    string `json:"int"`
	Luk                    string `json:"luk"`
	MaxHP                  string `json:"max_hp"`
	MaxMP                  string `json:"max_mp"`
	AttackPower            string `json:"attack_power"`
	MagicPower             string `json:"magic_power"`
	Armor                  string `json:"armor"`
	Speed                  string `json:"speed"`
	Jump                   string `json:"jump"`
	BossDamage             string `json:"boss_damage"`
	IgnoreMonsterArmor     string `json:"ignore_monster_armor"`
	AllStat                string `json:"all_stat"`
	Damage                 string `json:"damage"`
	EquipmentLevelDecrease int    `json:"equipment_level_decrease"`
	MaxHPRate              string `json:"max_hp_rate"`
	MaxMPRate              string `json:"max_mp_rate"`
}

type CharacterItemEquipmentBaseOption struct {
	Str                string `json:"str"`
	Dex                string `json:"dex"`
	Int                string `json:"int"`
	Luk                string `json:"luk"`
	MaxHP              string `json:"max_hp"`
	MaxMP              string `json:"max_mp"`
	AttackPower        string `json:"attack_power"`
	MagicPower         string `json:"magic_power"`
	Armor              string `json:"armor"`
	Speed              string `json:"speed"`
	Jump               string `json:"jump"`
	BossDamage         string `json:"boss_damage"`
	IgnoreMonsterArmor string `json:"ignore_monster_armor"`
	AllStat            string `json:"all_stat"`
	MaxHPRate          string `json:"max_hp_rate"`
	MaxMPRate          string `json:"max_mp_rate"`
	BaseEquipmentLevel int    `json:"base_equipment_level"`
}

type CharacterItemEquipmentAddOption struct {
	Str                    string `json:"str"`
	Dex                    string `json:"dex"`
	Int                    string `json:"int"`
	Luk                    string `json:"luk"`
	MaxHP                  string `json:"max_hp"`
	MaxMP                  string `json:"max_mp"`
	AttackPower            string `json:"attack_power"`
	MagicPower             string `json:"magic_power"`
	Armor                  string `json:"armor"`
	Speed                  string `json:"speed"`
	Jump                   string `json:"jump"`
	BossDamage             string `json:"boss_damage"`
	Damage                 string `json:"damage"`
	AllStat                string `json:"all_stat"`
	EquipmentLevelDecrease int    `json:"equipment_level_decrease"`
}

type CharacterItemEquipmentEtcOption struct {
	Str         string `json:"str"`
	Dex         string `json:"dex"`
	Int         string `json:"int"`
	Luk         string `json:"luk"`
	MaxHP       string `json:"max_hp"`
	MaxMP       string `json:"max_mp"`
	AttackPower string `json:"attack_power"`
	MagicPower  string `json:"magic_power"`
	Armor       string `json:"armor"`
	Speed       string `json:"speed"`
	Jump        string `json:"jump"`
}

type CharacterItemEquipmentStarforceOption = CharacterItemEquipmentEtcOption

type CharacterItemEquipmentTitle struct {
	TitleName             *string `json:"title_name"`
	TitleIcon             *string `json:"title_icon"`
	TitleDescription      *string `json:"title_description"`
	DateExpire            *string `json:"date_expire"`
	DateOptionExpire      *string `json:"date_option_expire"`
	TitleShapeName        *string `json:"title_shape_name"`
	TitleShapeIcon        *string `json:"title_shape_icon"`
	TitleShapeDescription *string `json:"title_shape_description"`
}

// CharacterCashItemEquipment represents cosmetic and cash-equipment data.
type CharacterCashItemEquipment struct {
	Date                               *string                            `json:"date"`
	CharacterGender                    *string                            `json:"character_gender"`
	CharacterClass                     *string                            `json:"character_class"`
	CharacterLookMode                  *string                            `json:"character_look_mode"`
	PresetNo                           *int                               `json:"preset_no"`
	CashItemEquipmentBase              []CharacterCashItemEquipmentPreset `json:"cash_item_equipment_base"`
	CashItemEquipmentPreset1           []CharacterCashItemEquipmentPreset `json:"cash_item_equipment_preset_1"`
	CashItemEquipmentPreset2           []CharacterCashItemEquipmentPreset `json:"cash_item_equipment_preset_2"`
	CashItemEquipmentPreset3           []CharacterCashItemEquipmentPreset `json:"cash_item_equipment_preset_3"`
	AdditionalCashItemEquipmentBase    []CharacterCashItemEquipmentPreset `json:"additional_cash_item_equipment_base"`
	AdditionalCashItemEquipmentPreset1 []CharacterCashItemEquipmentPreset `json:"additional_cash_item_equipment_preset_1"`
	AdditionalCashItemEquipmentPreset2 []CharacterCashItemEquipmentPreset `json:"additional_cash_item_equipment_preset_2"`
	AdditionalCashItemEquipmentPreset3 []CharacterCashItemEquipmentPreset `json:"additional_cash_item_equipment_preset_3"`
}

type CharacterCashItemEquipmentPreset struct {
	CashItemEquipmentPart string                                   `json:"cash_item_equipment_part"`
	CashItemEquipmentSlot string                                   `json:"cash_item_equipment_slot"`
	CashItemName          string                                   `json:"cash_item_name"`
	CashItemIcon          string                                   `json:"cash_item_icon"`
	CashItemDescription   *string                                  `json:"cash_item_description"`
	CashItemOption        []CharacterCashItemEquipmentOption       `json:"cash_item_option"`
	DateExpire            *string                                  `json:"date_expire"`
	DateOptionExpire      *string                                  `json:"date_option_expire"`
	CashItemLabel         *string                                  `json:"cash_item_label"`
	CashItemColoringPrism *CharacterCashItemEquipmentColoringPrism `json:"cash_item_coloring_prism"`
	ItemGender            *string                                  `json:"item_gender"`
	Skills                []string                                 `json:"skills"`
}

type CharacterCashItemEquipmentOption struct {
	OptionType  string `json:"option_type"`
	OptionValue string `json:"option_value"`
}

type CharacterCashItemEquipmentColoringPrism struct {
	ColorRange string `json:"color_range"`
	Hue        int    `json:"hue"`
	Saturation int    `json:"saturation"`
	Value      int    `json:"value"`
}

// CharacterSymbolEquipment represents symbol progression.
type CharacterSymbolEquipment struct {
	Date           *string                        `json:"date"`
	CharacterClass *string                        `json:"character_class"`
	Symbol         []CharacterSymbolEquipmentInfo `json:"symbol"`
}

type CharacterSymbolEquipmentInfo struct {
	SymbolName               string  `json:"symbol_name"`
	SymbolIcon               string  `json:"symbol_icon"`
	SymbolDescription        string  `json:"symbol_description"`
	SymbolForce              string  `json:"symbol_force"`
	SymbolLevel              int     `json:"symbol_level"`
	SymbolStr                string  `json:"symbol_str"`
	SymbolDex                string  `json:"symbol_dex"`
	SymbolInt                string  `json:"symbol_int"`
	SymbolLuk                string  `json:"symbol_luk"`
	SymbolHP                 string  `json:"symbol_hp"`
	SymbolDropRate           *string `json:"symbol_drop_rate"`
	SymbolMesoRate           *string `json:"symbol_meso_rate"`
	SymbolExpRate            *string `json:"symbol_exp_rate"`
	SymbolGrowthCount        int     `json:"symbol_growth_count"`
	SymbolRequireGrowthCount int     `json:"symbol_require_growth_count"`
}

// CharacterSetEffect represents active set effects.
type CharacterSetEffect struct {
	Date      *string                 `json:"date"`
	SetEffect []CharacterSetEffectSet `json:"set_effect"`
}

type CharacterSetEffectSet struct {
	SetName       string                         `json:"set_name"`
	TotalSetCount int                            `json:"total_set_count"`
	SetEffectInfo []CharacterSetEffectInfo       `json:"set_effect_info"`
	SetOptionFull []CharacterSetEffectOptionFull `json:"set_option_full"`
}

type CharacterSetEffectInfo struct {
	SetCount  int    `json:"set_count"`
	SetOption string `json:"set_option"`
}

type CharacterSetEffectOptionFull struct {
	SetCount  int    `json:"set_count"`
	SetOption string `json:"set_option"`
}

// CharacterBeautyEquipment represents hair/face/skin customization.
type CharacterBeautyEquipment struct {
	Date                    *string                       `json:"date"`
	CharacterGender         string                        `json:"character_gender"`
	CharacterClass          string                        `json:"character_class"`
	CharacterHair           *CharacterBeautyEquipmentHair `json:"character_hair"`
	CharacterFace           *CharacterBeautyEquipmentFace `json:"character_face"`
	CharacterSkin           *CharacterBeautyEquipmentSkin `json:"character_skin"`
	AdditionalCharacterHair *CharacterBeautyEquipmentHair `json:"additional_character_hair"`
	AdditionalCharacterFace *CharacterBeautyEquipmentFace `json:"additional_character_face"`
	AdditionalCharacterSkin *CharacterBeautyEquipmentSkin `json:"additional_character_skin"`
}

type CharacterBeautyEquipmentHair struct {
	HairName  string  `json:"hair_name"`
	BaseColor string  `json:"base_color"`
	MixColor  *string `json:"mix_color"`
	MixRate   string  `json:"mix_rate"`
}

type CharacterBeautyEquipmentFace struct {
	FaceName  string  `json:"face_name"`
	BaseColor string  `json:"base_color"`
	MixColor  *string `json:"mix_color"`
	MixRate   string  `json:"mix_rate"`
}

type CharacterBeautyEquipmentSkin struct {
	SkinName   string  `json:"skin_name"`
	ColorStyle *string `json:"color_style"`
	Hue        *int    `json:"hue"`
	Saturation *int    `json:"saturation"`
	Brightness *int    `json:"brightness"`
}

// CharacterAndroidEquipment represents android cosmetic information.
type CharacterAndroidEquipment struct {
	Date                     *string                             `json:"date"`
	AndroidName              *string                             `json:"android_name"`
	AndroidNickname          *string                             `json:"android_nickname"`
	AndroidIcon              *string                             `json:"android_icon"`
	AndroidDescription       *string                             `json:"android_description"`
	AndroidHair              *CharacterAndroidEquipmentHair      `json:"android_hair"`
	AndroidFace              *CharacterAndroidEquipmentFace      `json:"android_face"`
	AndroidSkin              *CharacterAndroidEquipmentSkin      `json:"android_skin"`
	AndroidCashItemEquipment []CharacterAndroidCashItemEquipment `json:"android_cash_item_equipment"`
	AndroidEarSensorClipFlag *string                             `json:"android_ear_sensor_clip_flag"`
	AndroidGender            *string                             `json:"android_gender"`
	AndroidGrade             *string                             `json:"android_grade"`
	AndroidNonHumanoidFlag   *string                             `json:"android_non_humanoid_flag"`
	AndroidShopUsableFlag    *string                             `json:"android_shop_usable_flag"`
	PresetNo                 *int                                `json:"preset_no"`
	AndroidPreset1           *CharacterAndroidEquipmentPreset    `json:"android_preset_1"`
	AndroidPreset2           *CharacterAndroidEquipmentPreset    `json:"android_preset_2"`
	AndroidPreset3           *CharacterAndroidEquipmentPreset    `json:"android_preset_3"`
}

type CharacterAndroidEquipmentHair struct {
	HairName  *string `json:"hair_name"`
	BaseColor *string `json:"base_color"`
	MixColor  *string `json:"mix_color"`
	MixRate   string  `json:"mix_rate"`
}

type CharacterAndroidEquipmentFace struct {
	FaceName  *string `json:"face_name"`
	BaseColor *string `json:"base_color"`
	MixColor  *string `json:"mix_color"`
	MixRate   string  `json:"mix_rate"`
}

type CharacterAndroidEquipmentSkin struct {
	SkinName   string  `json:"skin_name"`
	ColorStyle *string `json:"color_style"`
	Hue        *int    `json:"hue"`
	Saturation *int    `json:"saturation"`
	Brightness *int    `json:"brightness"`
}

type CharacterAndroidEquipmentPreset struct {
	AndroidName              string                         `json:"android_name"`
	AndroidNickname          string                         `json:"android_nickname"`
	AndroidIcon              string                         `json:"android_icon"`
	AndroidDescription       string                         `json:"android_description"`
	AndroidGender            *string                        `json:"android_gender"`
	AndroidGrade             string                         `json:"android_grade"`
	AndroidHair              CharacterAndroidEquipmentHair  `json:"android_hair"`
	AndroidFace              CharacterAndroidEquipmentFace  `json:"android_face"`
	AndroidSkin              *CharacterAndroidEquipmentSkin `json:"android_skin"`
	AndroidEarSensorClipFlag string                         `json:"android_ear_sensor_clip_flag"`
	AndroidNonHumanoidFlag   string                         `json:"android_non_humanoid_flag"`
	AndroidShopUsableFlag    string                         `json:"android_shop_usable_flag"`
}

type CharacterAndroidCashItemEquipment struct {
	CashItemEquipmentPart string                                          `json:"cash_item_equipment_part"`
	CashItemEquipmentSlot string                                          `json:"cash_item_equipment_slot"`
	CashItemName          string                                          `json:"cash_item_name"`
	CashItemIcon          string                                          `json:"cash_item_icon"`
	CashItemDescription   *string                                         `json:"cash_item_description"`
	CashItemOption        []CharacterAndroidCashItemEquipmentOption       `json:"cash_item_option"`
	DateExpire            *string                                         `json:"date_expire"`
	DateOptionExpire      *string                                         `json:"date_option_expire"`
	CashItemLabel         *string                                         `json:"cash_item_label"`
	CashItemColoringPrism *CharacterAndroidCashItemEquipmentColoringPrism `json:"cash_item_coloring_prism"`
	AndroidItemGender     *string                                         `json:"android_item_gender"`
}

type CharacterAndroidCashItemEquipmentOption struct {
	OptionType  string `json:"option_type"`
	OptionValue string `json:"option_value"`
}

type CharacterAndroidCashItemEquipmentColoringPrism struct {
	ColorRange string `json:"color_range"`
	Hue        int    `json:"hue"`
	Saturation int    `json:"saturation"`
	Value      int    `json:"value"`
}

// CharacterPetEquipment represents pet equipment and skills.
type CharacterPetEquipment struct {
	Date               *string                         `json:"date"`
	Pet1Name           *string                         `json:"pet_1_name"`
	Pet1Nickname       *string                         `json:"pet_1_nickname"`
	Pet1Icon           *string                         `json:"pet_1_icon"`
	Pet1Description    *string                         `json:"pet_1_description"`
	Pet1Equipment      *CharacterPetEquipmentItem      `json:"pet_1_equipment"`
	Pet1AutoSkill      *CharacterPetEquipmentAutoSkill `json:"pet_1_auto_skill"`
	Pet1PetType        *string                         `json:"pet_1_pet_type"`
	Pet1Skill          []string                        `json:"pet_1_skill"`
	Pet1DateExpire     *string                         `json:"pet_1_date_expire"`
	Pet1Appearance     *string                         `json:"pet_1_appearance"`
	Pet1AppearanceIcon *string                         `json:"pet_1_appearance_icon"`
	Pet2Name           *string                         `json:"pet_2_name"`
	Pet2Nickname       *string                         `json:"pet_2_nickname"`
	Pet2Icon           *string                         `json:"pet_2_icon"`
	Pet2Description    *string                         `json:"pet_2_description"`
	Pet2Equipment      *CharacterPetEquipmentItem      `json:"pet_2_equipment"`
	Pet2AutoSkill      *CharacterPetEquipmentAutoSkill `json:"pet_2_auto_skill"`
	Pet2PetType        *string                         `json:"pet_2_pet_type"`
	Pet2Skill          []string                        `json:"pet_2_skill"`
	Pet2DateExpire     *string                         `json:"pet_2_date_expire"`
	Pet2Appearance     *string                         `json:"pet_2_appearance"`
	Pet2AppearanceIcon *string                         `json:"pet_2_appearance_icon"`
	Pet3Name           *string                         `json:"pet_3_name"`
	Pet3Nickname       *string                         `json:"pet_3_nickname"`
	Pet3Icon           *string                         `json:"pet_3_icon"`
	Pet3Description    *string                         `json:"pet_3_description"`
	Pet3Equipment      *CharacterPetEquipmentItem      `json:"pet_3_equipment"`
	Pet3AutoSkill      *CharacterPetEquipmentAutoSkill `json:"pet_3_auto_skill"`
	Pet3PetType        *string                         `json:"pet_3_pet_type"`
	Pet3Skill          []string                        `json:"pet_3_skill"`
	Pet3DateExpire     *string                         `json:"pet_3_date_expire"`
	Pet3Appearance     *string                         `json:"pet_3_appearance"`
	Pet3AppearanceIcon *string                         `json:"pet_3_appearance_icon"`
}

type CharacterPetEquipmentItem struct {
	ItemName         *string                           `json:"item_name"`
	ItemIcon         *string                           `json:"item_icon"`
	ItemDescription  *string                           `json:"item_description"`
	ItemOption       []CharacterPetEquipmentItemOption `json:"item_option"`
	ScrollUpgrade    int                               `json:"scroll_upgrade"`
	ScrollUpgradable int                               `json:"scroll_upgradable"`
	ItemShape        *string                           `json:"item_shape"`
	ItemShapeIcon    *string                           `json:"item_shape_icon"`
}

type CharacterPetEquipmentItemOption struct {
	OptionType  string `json:"option_type"`
	OptionValue string `json:"option_value"`
}

type CharacterPetEquipmentAutoSkill struct {
	Skill1     *string `json:"skill_1"`
	Skill1Icon *string `json:"skill_1_icon"`
	Skill2     *string `json:"skill_2"`
	Skill2Icon *string `json:"skill_2_icon"`
}

// CharacterSkill represents skill info.
type CharacterSkill struct {
	Date                *string              `json:"date"`
	CharacterClass      *string              `json:"character_class"`
	CharacterSkillGrade *string              `json:"character_skill_grade"`
	CharacterSkill      []CharacterSkillInfo `json:"character_skill"`
}

type CharacterSkillInfo struct {
	SkillName        string  `json:"skill_name"`
	SkillDescription string  `json:"skill_description"`
	SkillLevel       int     `json:"skill_level"`
	SkillEffect      *string `json:"skill_effect"`
	SkillEffectNext  *string `json:"skill_effect_next"`
	SkillIcon        string  `json:"skill_icon"`
}

// CharacterLinkSkill represents link skills and presets.
type CharacterLinkSkill struct {
	Date                           *string                  `json:"date"`
	CharacterClass                 *string                  `json:"character_class"`
	CharacterLinkSkill             []CharacterLinkSkillInfo `json:"character_link_skill"`
	CharacterLinkSkillPreset1      []CharacterLinkSkillInfo `json:"character_link_skill_preset_1"`
	CharacterLinkSkillPreset2      []CharacterLinkSkillInfo `json:"character_link_skill_preset_2"`
	CharacterLinkSkillPreset3      []CharacterLinkSkillInfo `json:"character_link_skill_preset_3"`
	CharacterOwnedLinkSkill        *CharacterLinkSkillInfo  `json:"character_owned_link_skill"`
	CharacterOwnedLinkSkillPreset1 *CharacterLinkSkillInfo  `json:"character_owned_link_skill_preset_1"`
	CharacterOwnedLinkSkillPreset2 *CharacterLinkSkillInfo  `json:"character_owned_link_skill_preset_2"`
	CharacterOwnedLinkSkillPreset3 *CharacterLinkSkillInfo  `json:"character_owned_link_skill_preset_3"`
}

type CharacterLinkSkillInfo struct {
	SkillName        string  `json:"skill_name"`
	SkillDescription string  `json:"skill_description"`
	SkillLevel       int     `json:"skill_level"`
	SkillEffect      string  `json:"skill_effect"`
	SkillEffectNext  *string `json:"skill_effect_next"`
	SkillIcon        string  `json:"skill_icon"`
}

// CharacterVMatrix represents V matrix data.
type CharacterVMatrix struct {
	Date                                   *string                         `json:"date"`
	CharacterClass                         *string                         `json:"character_class"`
	CharacterVCoreEquipment                []CharacterVMatrixCoreEquipment `json:"character_v_core_equipment"`
	CharacterVMatrixRemainSlotUpgradePoint *int                            `json:"character_v_matrix_remain_slot_upgrade_point"`
}

type CharacterVMatrixCoreEquipment struct {
	SlotID      string  `json:"slot_id"`
	SlotLevel   int     `json:"slot_level"`
	VCoreName   *string `json:"v_core_name"`
	VCoreType   *string `json:"v_core_type"`
	VCoreLevel  int     `json:"v_core_level"`
	VCoreSkill1 string  `json:"v_core_skill_1"`
	VCoreSkill2 *string `json:"v_core_skill_2"`
	VCoreSkill3 *string `json:"v_core_skill_3"`
}

// CharacterHexaMatrix represents HEXA matrix data.
type CharacterHexaMatrix struct {
	Date                       *string                        `json:"date"`
	CharacterHexaCoreEquipment []CharacterHexaMatrixEquipment `json:"character_hexa_core_equipment"`
}

type CharacterHexaMatrixEquipment struct {
	HexaCoreName  string                                    `json:"hexa_core_name"`
	HexaCoreLevel int                                       `json:"hexa_core_level"`
	HexaCoreType  string                                    `json:"hexa_core_type"`
	LinkedSkill   []CharacterHexaMatrixEquipmentLinkedSkill `json:"linked_skill"`
}

type CharacterHexaMatrixEquipmentLinkedSkill struct {
	HexaSkillID string `json:"hexa_skill_id"`
}

// CharacterHexaMatrixStat represents HEXA stat configurations.
type CharacterHexaMatrixStat struct {
	Date                   *string                       `json:"date"`
	CharacterClass         *string                       `json:"character_class"`
	CharacterHexaStatCore  []CharacterHexaMatrixStatCore `json:"character_hexa_stat_core"`
	CharacterHexaStatCore2 []CharacterHexaMatrixStatCore `json:"character_hexa_stat_core_2"`
	CharacterHexaStatCore3 []CharacterHexaMatrixStatCore `json:"character_hexa_stat_core_3"`
	PresetHexaStatCore     []CharacterHexaMatrixStatCore `json:"preset_hexa_stat_core"`
	PresetHexaStatCore2    []CharacterHexaMatrixStatCore `json:"preset_hexa_stat_core_2"`
	PresetHexaStatCore3    []CharacterHexaMatrixStatCore `json:"preset_hexa_stat_core_3"`
}

type CharacterHexaMatrixStatCore struct {
	SlotID        string `json:"slot_id"`
	MainStatName  string `json:"main_stat_name"`
	SubStatName1  string `json:"sub_stat_name_1"`
	SubStatName2  string `json:"sub_stat_name_2"`
	MainStatLevel int    `json:"main_stat_level"`
	SubStatLevel1 int    `json:"sub_stat_level_1"`
	SubStatLevel2 int    `json:"sub_stat_level_2"`
	StatGrade     int    `json:"stat_grade"`
}

// CharacterDojang represents Dojang record.
type CharacterDojang struct {
	Date             *string `json:"date"`
	CharacterClass   *string `json:"character_class"`
	WorldName        *string `json:"world_name"`
	DojangBestFloor  *int    `json:"dojang_best_floor"`
	DateDojangRecord *string `json:"date_dojang_record"`
	DojangBestTime   *int    `json:"dojang_best_time"`
}
