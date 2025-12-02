package kms

import (
	"github.com/dongecs/maplestory-openapi/api/common"
	"github.com/dongecs/maplestory-openapi/api/tms"
)

// Shared character-related structures (payloads match TMS with additional KMS-only fields elsewhere).
type (
	Character                                      = common.Character
	CharacterBasic                                 = common.CharacterBasic
	CharacterImage                                 = common.CharacterImage
	CharacterPopularity                            = common.CharacterPopularity
	CharacterStat                                  = common.CharacterStat
	CharacterFinalStat                             = common.CharacterFinalStat
	CharacterHyperStat                             = common.CharacterHyperStat
	CharacterHyperStatPreset                       = common.CharacterHyperStatPreset
	CharacterPropensity                            = common.CharacterPropensity
	CharacterAbility                               = common.CharacterAbility
	CharacterAbilityInfo                           = common.CharacterAbilityInfo
	CharacterAbilityPreset                         = common.CharacterAbilityPreset
	CharacterItemEquipment                         = common.CharacterItemEquipment
	CharacterItemEquipmentInfo                     = common.CharacterItemEquipmentInfo
	CharacterItemEquipmentMechanic                 = common.CharacterItemEquipmentMechanic
	CharacterItemEquipmentOption                   = common.CharacterItemEquipmentOption
	CharacterItemEquipmentExceptionalOption        = common.CharacterItemEquipmentExceptionalOption
	CharacterItemEquipmentTotalOption              = common.CharacterItemEquipmentTotalOption
	CharacterItemEquipmentBaseOption               = common.CharacterItemEquipmentBaseOption
	CharacterItemEquipmentAddOption                = common.CharacterItemEquipmentAddOption
	CharacterItemEquipmentEtcOption                = common.CharacterItemEquipmentEtcOption
	CharacterItemEquipmentStarforceOption          = common.CharacterItemEquipmentStarforceOption
	CharacterItemEquipmentTitle                    = common.CharacterItemEquipmentTitle
	CharacterCashItemEquipment                     = common.CharacterCashItemEquipment
	CharacterCashItemEquipmentPreset               = common.CharacterCashItemEquipmentPreset
	CharacterCashItemEquipmentOption               = common.CharacterCashItemEquipmentOption
	CharacterCashItemEquipmentColoringPrism        = common.CharacterCashItemEquipmentColoringPrism
	CharacterSymbolEquipment                       = common.CharacterSymbolEquipment
	CharacterSymbolEquipmentInfo                   = common.CharacterSymbolEquipmentInfo
	CharacterSetEffect                             = common.CharacterSetEffect
	CharacterSetEffectSet                          = common.CharacterSetEffectSet
	CharacterSetEffectInfo                         = common.CharacterSetEffectInfo
	CharacterSetEffectOptionFull                   = common.CharacterSetEffectOptionFull
	CharacterBeautyEquipment                       = common.CharacterBeautyEquipment
	CharacterBeautyEquipmentHair                   = common.CharacterBeautyEquipmentHair
	CharacterBeautyEquipmentFace                   = common.CharacterBeautyEquipmentFace
	CharacterBeautyEquipmentSkin                   = common.CharacterBeautyEquipmentSkin
	CharacterAndroidEquipment                      = common.CharacterAndroidEquipment
	CharacterAndroidEquipmentHair                  = common.CharacterAndroidEquipmentHair
	CharacterAndroidEquipmentFace                  = common.CharacterAndroidEquipmentFace
	CharacterAndroidEquipmentSkin                  = common.CharacterAndroidEquipmentSkin
	CharacterAndroidEquipmentPreset                = common.CharacterAndroidEquipmentPreset
	CharacterAndroidCashItemEquipment              = common.CharacterAndroidCashItemEquipment
	CharacterAndroidCashItemEquipmentOption        = common.CharacterAndroidCashItemEquipmentOption
	CharacterAndroidCashItemEquipmentColoringPrism = common.CharacterAndroidCashItemEquipmentColoringPrism
	CharacterPetEquipment                          = common.CharacterPetEquipment
	CharacterPetEquipmentItem                      = common.CharacterPetEquipmentItem
	CharacterPetEquipmentItemOption                = common.CharacterPetEquipmentItemOption
	CharacterPetEquipmentAutoSkill                 = common.CharacterPetEquipmentAutoSkill
	CharacterSkill                                 = common.CharacterSkill
	CharacterSkillInfo                             = common.CharacterSkillInfo
	CharacterLinkSkill                             = common.CharacterLinkSkill
	CharacterLinkSkillInfo                         = common.CharacterLinkSkillInfo
	CharacterVMatrix                               = common.CharacterVMatrix
	CharacterVMatrixCoreEquipment                  = common.CharacterVMatrixCoreEquipment
	CharacterHexaMatrix                            = common.CharacterHexaMatrix
	CharacterHexaMatrixEquipment                   = common.CharacterHexaMatrixEquipment
	CharacterHexaMatrixEquipmentLinkedSkill        = common.CharacterHexaMatrixEquipmentLinkedSkill
	CharacterHexaMatrixStat                        = common.CharacterHexaMatrixStat
	CharacterHexaMatrixStatCore                    = common.CharacterHexaMatrixStatCore
	CharacterDojang                                = common.CharacterDojang
)

// CharacterImageOptions mirrors rendering options.
type CharacterImageOptions = tms.CharacterImageOptions

// CharacterList holds account-linked characters.
type CharacterList struct {
	AccountList []CharacterAccount `json:"account_list"`
}

type CharacterAccount struct {
	AccountID     string                 `json:"account_id"`
	CharacterList []CharacterAccountInfo `json:"character_list"`
}

type CharacterAccountInfo struct {
	OCID           string `json:"ocid"`
	CharacterName  string `json:"character_name"`
	WorldName      string `json:"world_name"`
	CharacterClass string `json:"character_class"`
	CharacterLevel int    `json:"character_level"`
}

// Achievement groups achievements per account.
type Achievement struct {
	AccountList []AchievementAccount `json:"account_list"`
}

type AchievementAccount struct {
	AccountID          string             `json:"account_id"`
	AchievementAchieve []AchievementEntry `json:"achievement_achieve"`
}

type AchievementEntry struct {
	AchievementName        string `json:"achievement_name"`
	AchievementDescription string `json:"achievement_description"`
}

// CharacterOtherStat represents other stat data (structure kept generic).
type CharacterOtherStat map[string]any

// CharacterRingExchangeSkillEquipment keeps the raw payload.
type CharacterRingExchangeSkillEquipment map[string]any
