package msea

import "github.com/dongecs/maplestory-openapi/api/tms"

// Re-export character-related structures from TMS since payloads match.
type (
	Character                                      = tms.Character
	CharacterBasic                                 = tms.CharacterBasic
	CharacterImage                                 = tms.CharacterImage
	CharacterPopularity                            = tms.CharacterPopularity
	CharacterStat                                  = tms.CharacterStat
	CharacterFinalStat                             = tms.CharacterFinalStat
	CharacterHyperStat                             = tms.CharacterHyperStat
	CharacterHyperStatPreset                       = tms.CharacterHyperStatPreset
	CharacterPropensity                            = tms.CharacterPropensity
	CharacterAbility                               = tms.CharacterAbility
	CharacterAbilityInfo                           = tms.CharacterAbilityInfo
	CharacterAbilityPreset                         = tms.CharacterAbilityPreset
	CharacterItemEquipment                         = tms.CharacterItemEquipment
	CharacterItemEquipmentInfo                     = tms.CharacterItemEquipmentInfo
	CharacterItemEquipmentMechanic                 = tms.CharacterItemEquipmentMechanic
	CharacterItemEquipmentOption                   = tms.CharacterItemEquipmentOption
	CharacterItemEquipmentExceptionalOption        = tms.CharacterItemEquipmentExceptionalOption
	CharacterItemEquipmentTotalOption              = tms.CharacterItemEquipmentTotalOption
	CharacterItemEquipmentBaseOption               = tms.CharacterItemEquipmentBaseOption
	CharacterItemEquipmentAddOption                = tms.CharacterItemEquipmentAddOption
	CharacterItemEquipmentEtcOption                = tms.CharacterItemEquipmentEtcOption
	CharacterItemEquipmentStarforceOption          = tms.CharacterItemEquipmentStarforceOption
	CharacterItemEquipmentTitle                    = tms.CharacterItemEquipmentTitle
	CharacterCashItemEquipment                     = tms.CharacterCashItemEquipment
	CharacterCashItemEquipmentPreset               = tms.CharacterCashItemEquipmentPreset
	CharacterCashItemEquipmentOption               = tms.CharacterCashItemEquipmentOption
	CharacterCashItemEquipmentColoringPrism        = tms.CharacterCashItemEquipmentColoringPrism
	CharacterSymbolEquipment                       = tms.CharacterSymbolEquipment
	CharacterSymbolEquipmentInfo                   = tms.CharacterSymbolEquipmentInfo
	CharacterSetEffect                             = tms.CharacterSetEffect
	CharacterSetEffectSet                          = tms.CharacterSetEffectSet
	CharacterSetEffectInfo                         = tms.CharacterSetEffectInfo
	CharacterSetEffectOptionFull                   = tms.CharacterSetEffectOptionFull
	CharacterBeautyEquipment                       = tms.CharacterBeautyEquipment
	CharacterBeautyEquipmentHair                   = tms.CharacterBeautyEquipmentHair
	CharacterBeautyEquipmentFace                   = tms.CharacterBeautyEquipmentFace
	CharacterBeautyEquipmentSkin                   = tms.CharacterBeautyEquipmentSkin
	CharacterAndroidEquipment                      = tms.CharacterAndroidEquipment
	CharacterAndroidEquipmentHair                  = tms.CharacterAndroidEquipmentHair
	CharacterAndroidEquipmentFace                  = tms.CharacterAndroidEquipmentFace
	CharacterAndroidEquipmentSkin                  = tms.CharacterAndroidEquipmentSkin
	CharacterAndroidEquipmentPreset                = tms.CharacterAndroidEquipmentPreset
	CharacterAndroidCashItemEquipment              = tms.CharacterAndroidCashItemEquipment
	CharacterAndroidCashItemEquipmentOption        = tms.CharacterAndroidCashItemEquipmentOption
	CharacterAndroidCashItemEquipmentColoringPrism = tms.CharacterAndroidCashItemEquipmentColoringPrism
	CharacterPetEquipment                          = tms.CharacterPetEquipment
	CharacterPetEquipmentItem                      = tms.CharacterPetEquipmentItem
	CharacterPetEquipmentItemOption                = tms.CharacterPetEquipmentItemOption
	CharacterPetEquipmentAutoSkill                 = tms.CharacterPetEquipmentAutoSkill
	CharacterSkill                                 = tms.CharacterSkill
	CharacterSkillInfo                             = tms.CharacterSkillInfo
	CharacterLinkSkill                             = tms.CharacterLinkSkill
	CharacterLinkSkillInfo                         = tms.CharacterLinkSkillInfo
	CharacterVMatrix                               = tms.CharacterVMatrix
	CharacterVMatrixCoreEquipment                  = tms.CharacterVMatrixCoreEquipment
	CharacterHexaMatrix                            = tms.CharacterHexaMatrix
	CharacterHexaMatrixEquipment                   = tms.CharacterHexaMatrixEquipment
	CharacterHexaMatrixEquipmentLinkedSkill        = tms.CharacterHexaMatrixEquipmentLinkedSkill
	CharacterHexaMatrixStat                        = tms.CharacterHexaMatrixStat
	CharacterHexaMatrixStatCore                    = tms.CharacterHexaMatrixStatCore
	CharacterDojang                                = tms.CharacterDojang
)

// CharacterImageOptions mirrors tms rendering options.
type CharacterImageOptions = tms.CharacterImageOptions
