package common

import "context"

// MapleClient defines the shared region client surface for character/union/guild APIs.
type MapleClient interface {
	GetCharacter(ctx context.Context, characterName string) (*Character, error)
	GetCharacterBasic(ctx context.Context, ocid string, date any) (*CharacterBasic, error)
	GetCharacterImage(ctx context.Context, ocid string, opts *CharacterImageOptions, date any) (*CharacterImage, error)
	GetCharacterPopularity(ctx context.Context, ocid string, date any) (*CharacterPopularity, error)
	GetCharacterStat(ctx context.Context, ocid string, date any) (*CharacterStat, error)
	GetCharacterHyperStat(ctx context.Context, ocid string, date any) (*CharacterHyperStat, error)
	GetCharacterPropensity(ctx context.Context, ocid string, date any) (*CharacterPropensity, error)
	GetCharacterAbility(ctx context.Context, ocid string, date any) (*CharacterAbility, error)
	GetCharacterItemEquipment(ctx context.Context, ocid string, date any) (*CharacterItemEquipment, error)
	GetCharacterCashItemEquipment(ctx context.Context, ocid string, date any) (*CharacterCashItemEquipment, error)
	GetCharacterSymbolEquipment(ctx context.Context, ocid string, date any) (*CharacterSymbolEquipment, error)
	GetCharacterSetEffect(ctx context.Context, ocid string, date any) (*CharacterSetEffect, error)
	GetCharacterBeautyEquipment(ctx context.Context, ocid string, date any) (*CharacterBeautyEquipment, error)
	GetCharacterAndroidEquipment(ctx context.Context, ocid string, date any) (*CharacterAndroidEquipment, error)
	GetCharacterPetEquipment(ctx context.Context, ocid string, date any) (*CharacterPetEquipment, error)
	GetCharacterSkill(ctx context.Context, ocid, characterSkillGrade string, date any) (*CharacterSkill, error)
	GetCharacterLinkSkill(ctx context.Context, ocid string, date any) (*CharacterLinkSkill, error)
	GetCharacterVMatrix(ctx context.Context, ocid string, date any) (*CharacterVMatrix, error)
	GetCharacterHexaMatrix(ctx context.Context, ocid string, date any) (*CharacterHexaMatrix, error)
	GetCharacterHexaMatrixStat(ctx context.Context, ocid string, date any) (*CharacterHexaMatrixStat, error)
	GetCharacterDojang(ctx context.Context, ocid string, date any) (*CharacterDojang, error)
	GetUnion(ctx context.Context, ocid string, date any) (*Union, error)
	GetUnionRaider(ctx context.Context, ocid string, date any) (*UnionRaider, error)
	GetUnionArtifact(ctx context.Context, ocid string, date any) (*UnionArtifact, error)
	GetGuild(ctx context.Context, guildName, worldName string) (*Guild, error)
	GetGuildBasic(ctx context.Context, guildID string, date any) (*GuildBasic, error)
}
