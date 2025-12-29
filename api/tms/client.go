package tms

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dongecs/maplestory-openapi/api/common"
)

// Client is the MapleStory OpenAPI client for the TMS region.
type Client struct {
	*common.Client
}

var _ common.MapleClient = (*Client)(nil)

// NewClient constructs a TMS client with the given API key.
func NewClient(apiKey string) *Client {
	loc := common.FixedOffsetLocation(timezoneOffset)
	return &Client{
		Client: common.NewClient(apiKey, subPath, loc),
	}
}

// CharacterImageOptions defines rendering options for character images.
type CharacterImageOptions = common.CharacterImageOptions

// GetCharacter retrieves the ocid for a character name.
func (c *Client) GetCharacter(ctx context.Context, characterName string) (*Character, error) {
	query := map[string]string{
		"character_name": characterName,
	}

	var resp Character
	if err := c.Fetch(ctx, "v1/id", query, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetCharacterBasic fetches basic character information.
func (c *Client) GetCharacterBasic(ctx context.Context, ocid string, date any) (*CharacterBasic, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	raw, err := c.fetchRaw(ctx, "v1/character/basic", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
	if err != nil {
		return nil, err
	}

	if empty, err := isEmpty(raw); err != nil {
		return nil, err
	} else if empty {
		return nil, nil
	}

	var resp CharacterBasic
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	resp.CharacterImage = common.RemoveQuery(resp.CharacterImage)

	return &resp, nil
}

// GetCharacterImage renders the character image with the requested options.
func (c *Client) GetCharacterImage(ctx context.Context, ocid string, opts *CharacterImageOptions, date any) (*CharacterImage, error) {
	basic, err := c.GetCharacterBasic(ctx, ocid, date)
	if err != nil {
		return nil, err
	}
	if basic == nil {
		return nil, nil
	}

	action := common.CharacterImageActionStand1
	emotion := common.CharacterImageEmotionDefault
	wmotion := common.CharacterImageWeaponMotionDefault
	actionFrame := 0
	emotionFrame := 0
	width := defaultImageSize
	height := defaultImageSize
	var x, y *int

	if opts != nil {
		if opts.Action != "" {
			action = opts.Action
		}
		if opts.Emotion != "" {
			emotion = opts.Emotion
		}
		if opts.WeaponMotion != "" {
			wmotion = opts.WeaponMotion
		}
		if opts.ActionFrame != 0 {
			actionFrame = opts.ActionFrame
		}
		if opts.EmotionFrame != 0 {
			emotionFrame = opts.EmotionFrame
		}
		if opts.Width != 0 {
			width = opts.Width
		}
		if opts.Height != 0 {
			height = opts.Height
		}
		x = opts.X
		y = opts.Y
	}

	baseURL := common.RemoveQuery(basic.CharacterImage)
	query := map[string]string{
		"action":  fmt.Sprintf("%s.%d", action, actionFrame),
		"emotion": fmt.Sprintf("%s.%d", emotion, emotionFrame),
		"wmotion": string(wmotion),
		"width":   strconv.Itoa(width),
		"height":  strconv.Itoa(height),
	}
	if x != nil {
		query["x"] = strconv.Itoa(*x)
	}
	if y != nil {
		query["y"] = strconv.Itoa(*y)
	}

	originImage, err := c.fetchImageBase64(ctx, baseURL, nil)
	if err != nil {
		return nil, err
	}
	renderedImage, err := c.fetchImageBase64(ctx, baseURL, query)
	if err != nil {
		return nil, err
	}

	return &CharacterImage{
		Date:         basic.Date,
		OriginURL:    baseURL,
		OriginImage:  originImage,
		Image:        renderedImage,
		Action:       string(action),
		Emotion:      string(emotion),
		Wmotion:      string(wmotion),
		ActionFrame:  actionFrame,
		EmotionFrame: emotionFrame,
		Width:        width,
		Height:       height,
		X:            x,
		Y:            y,
	}, nil
}

// GetCharacterPopularity fetches popularity for a character.
func (c *Client) GetCharacterPopularity(ctx context.Context, ocid string, date any) (*CharacterPopularity, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterPopularity](ctx, c, "v1/character/popularity", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterStat fetches aggregate stat data.
func (c *Client) GetCharacterStat(ctx context.Context, ocid string, date any) (*CharacterStat, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterStat](ctx, c, "v1/character/stat", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterHyperStat fetches hyper stat presets.
func (c *Client) GetCharacterHyperStat(ctx context.Context, ocid string, date any) (*CharacterHyperStat, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterHyperStat](ctx, c, "v1/character/hyper-stat", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterPropensity fetches propensity information.
func (c *Client) GetCharacterPropensity(ctx context.Context, ocid string, date any) (*CharacterPropensity, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterPropensity](ctx, c, "v1/character/propensity", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterAbility fetches ability information.
func (c *Client) GetCharacterAbility(ctx context.Context, ocid string, date any) (*CharacterAbility, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterAbility](ctx, c, "v1/character/ability", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterItemEquipment fetches equipped items.
func (c *Client) GetCharacterItemEquipment(ctx context.Context, ocid string, date any) (*CharacterItemEquipment, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterItemEquipment](ctx, c, "v1/character/item-equipment", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterCashItemEquipment fetches cosmetic equipment.
func (c *Client) GetCharacterCashItemEquipment(ctx context.Context, ocid string, date any) (*CharacterCashItemEquipment, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterCashItemEquipment](ctx, c, "v1/character/cashitem-equipment", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterSymbolEquipment fetches symbol equipment.
func (c *Client) GetCharacterSymbolEquipment(ctx context.Context, ocid string, date any) (*CharacterSymbolEquipment, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterSymbolEquipment](ctx, c, "v1/character/symbol-equipment", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterSetEffect fetches set effect data.
func (c *Client) GetCharacterSetEffect(ctx context.Context, ocid string, date any) (*CharacterSetEffect, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterSetEffect](ctx, c, "v1/character/set-effect", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterBeautyEquipment fetches hair/face/skin info.
func (c *Client) GetCharacterBeautyEquipment(ctx context.Context, ocid string, date any) (*CharacterBeautyEquipment, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterBeautyEquipment](ctx, c, "v1/character/beauty-equipment", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterAndroidEquipment fetches android cosmetics.
func (c *Client) GetCharacterAndroidEquipment(ctx context.Context, ocid string, date any) (*CharacterAndroidEquipment, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterAndroidEquipment](ctx, c, "v1/character/android-equipment", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterPetEquipment fetches pet equipment/skills.
func (c *Client) GetCharacterPetEquipment(ctx context.Context, ocid string, date any) (*CharacterPetEquipment, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterPetEquipment](ctx, c, "v1/character/pet-equipment", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterSkill fetches skill information for a skill grade.
func (c *Client) GetCharacterSkill(ctx context.Context, ocid, characterSkillGrade string, date any) (*CharacterSkill, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterSkill](ctx, c, "v1/character/skill", map[string]string{
		"ocid":                  ocid,
		"character_skill_grade": characterSkillGrade,
		"date":                  dateStr,
	})
}

// GetCharacterLinkSkill fetches link skill info.
func (c *Client) GetCharacterLinkSkill(ctx context.Context, ocid string, date any) (*CharacterLinkSkill, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterLinkSkill](ctx, c, "v1/character/link-skill", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterVMatrix fetches V matrix info.
func (c *Client) GetCharacterVMatrix(ctx context.Context, ocid string, date any) (*CharacterVMatrix, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterVMatrix](ctx, c, "v1/character/vmatrix", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterHexaMatrix fetches HEXA matrix info.
func (c *Client) GetCharacterHexaMatrix(ctx context.Context, ocid string, date any) (*CharacterHexaMatrix, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterHexaMatrix](ctx, c, "v1/character/hexamatrix", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterHexaMatrixStat fetches HEXA stat info.
func (c *Client) GetCharacterHexaMatrixStat(ctx context.Context, ocid string, date any) (*CharacterHexaMatrixStat, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterHexaMatrixStat](ctx, c, "v1/character/hexamatrix-stat", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetCharacterDojang fetches Dojang records.
func (c *Client) GetCharacterDojang(ctx context.Context, ocid string, date any) (*CharacterDojang, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[CharacterDojang](ctx, c, "v1/character/dojang", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetUnion fetches union information.
func (c *Client) GetUnion(ctx context.Context, ocid string, date any) (*Union, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[Union](ctx, c, "v1/user/union", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetUnionRaider fetches union raider details.
func (c *Client) GetUnionRaider(ctx context.Context, ocid string, date any) (*UnionRaider, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[UnionRaider](ctx, c, "v1/user/union-raider", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetUnionArtifact fetches union artifact information.
func (c *Client) GetUnionArtifact(ctx context.Context, ocid string, date any) (*UnionArtifact, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[UnionArtifact](ctx, c, "v1/user/union-artifact", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetUnionChampion fetches union champion information.
func (c *Client) GetUnionChampion(ctx context.Context, ocid string, date any) (*UnionChampion, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[UnionChampion](ctx, c, "v1/user/union-champion", map[string]string{
		"ocid": ocid,
		"date": dateStr,
	})
}

// GetGuild fetches guild id by name/world.
func (c *Client) GetGuild(ctx context.Context, guildName, worldName string) (*Guild, error) {
	raw, err := c.fetchRaw(ctx, "v1/guild/id", map[string]string{
		"guild_name": guildName,
		"world_name": worldName,
	})
	if err != nil {
		return nil, err
	}

	if empty, err := isEmpty(raw); err != nil {
		return nil, err
	} else if empty {
		return nil, nil
	}

	var resp Guild
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetGuildBasic fetches guild basic info.
func (c *Client) GetGuildBasic(ctx context.Context, guildID string, date any) (*GuildBasic, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	return fetchWithEmpty[GuildBasic](ctx, c, "v1/guild/basic", map[string]string{
		"oguild_id": guildID,
		"date":      dateStr,
	})
}

// Ranking (unsupported)

func (c *Client) GetOverallRanking(ctx context.Context, filter *common.RankingFilter, date any) (*common.OverallRankingResponse, error) {
	return nil, fmt.Errorf("ranking endpoints are not supported for TMS")
}

func (c *Client) GetUnionRanking(ctx context.Context, filter *common.RankingFilter, date any) (*common.UnionRankingResponse, error) {
	return nil, fmt.Errorf("ranking endpoints are not supported for TMS")
}

func (c *Client) GetGuildRanking(ctx context.Context, filter *common.RankingFilter, date any) (*common.GuildRankingResponse, error) {
	return nil, fmt.Errorf("ranking endpoints are not supported for TMS")
}

func (c *Client) GetDojangRanking(ctx context.Context, filter *common.RankingFilter, date any) (*common.DojangRankingResponse, error) {
	return nil, fmt.Errorf("ranking endpoints are not supported for TMS")
}

func (c *Client) GetSeedRanking(ctx context.Context, filter *common.RankingFilter, date any) (*common.SeedRankingResponse, error) {
	return nil, fmt.Errorf("ranking endpoints are not supported for TMS")
}

func (c *Client) GetAchievementRanking(ctx context.Context, filter *common.RankingFilter, date any) (*common.AchievementRankingResponse, error) {
	return nil, fmt.Errorf("ranking endpoints are not supported for TMS")
}

// fetchWithEmpty fetches into T and returns nil when the response is effectively empty.
func fetchWithEmpty[T any](ctx context.Context, c *Client, path string, query map[string]string) (*T, error) {
	raw, err := c.fetchRaw(ctx, path, query)
	if err != nil {
		return nil, err
	}

	if empty, err := isEmpty(raw); err != nil {
		return nil, err
	} else if empty {
		return nil, nil
	}

	var resp T
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) fetchRaw(ctx context.Context, path string, query map[string]string) (json.RawMessage, error) {
	var raw json.RawMessage
	if err := c.Fetch(ctx, path, query, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

func (c *Client) toDateString(value any) (string, error) {
	if value == nil {
		return "", nil
	}
	return c.ToDateString(value, &minimumAPIDate)
}

func isEmpty(raw json.RawMessage) (bool, error) {
	if len(raw) == 0 {
		return true, nil
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return false, err
	}
	return common.IsEmptyResponse(m), nil
}

func (c *Client) fetchImageBase64(ctx context.Context, rawURL string, query map[string]string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if len(query) > 0 {
		q := u.Query()
		for k, v := range query {
			if v == "" {
				continue
			}
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: c.Timeout(),
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	mimeType := resp.Header.Get("Content-Type")
	encoded := base64.StdEncoding.EncodeToString(body)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, encoded), nil
}
