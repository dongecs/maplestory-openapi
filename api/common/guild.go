package common

// Guild contains oguild id.
type Guild struct {
	OGuildID string `json:"oguild_id"`
}

// GuildBasic represents guild details.
type GuildBasic struct {
	Date               *string      `json:"date"`
	WorldName          string       `json:"world_name"`
	GuildName          string       `json:"guild_name"`
	GuildLevel         int          `json:"guild_level"`
	GuildFame          int          `json:"guild_fame"`
	GuildPoint         int          `json:"guild_point"`
	GuildMasterName    string       `json:"guild_master_name"`
	GuildMemberCount   int          `json:"guild_member_count"`
	GuildMember        []string     `json:"guild_member"`
	GuildSkill         []GuildSkill `json:"guild_skill"`
	GuildNoblesseSkill []GuildSkill `json:"guild_noblesse_skill"`
}

type GuildSkill struct {
	SkillName        string `json:"skill_name"`
	SkillDescription string `json:"skill_description"`
	SkillLevel       int    `json:"skill_level"`
	SkillEffect      string `json:"skill_effect"`
	SkillIcon        string `json:"skill_icon"`
}
