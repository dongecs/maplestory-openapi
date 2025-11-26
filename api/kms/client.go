package kms

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

// Client is the MapleStory OpenAPI client for the KMS region.
type Client struct {
	*common.Client
}

// NewClient constructs a KMS client with the given API key.
func NewClient(apiKey string) *Client {
	loc := common.FixedOffsetLocation(timezoneOffset)
	return &Client{
		Client: common.NewClient(apiKey, subPath, loc),
	}
}

// Account

func (c *Client) GetCharacterList(ctx context.Context) (*CharacterList, error) {
	var resp CharacterList
	if err := c.Fetch(ctx, "v1/character/list", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetAchievement(ctx context.Context) (*Achievement, error) {
	var resp Achievement
	if err := c.Fetch(ctx, "v1/user/achievement", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Character

func (c *Client) GetCharacter(ctx context.Context, characterName string) (*Character, error) {
	var resp Character
	if err := c.Fetch(ctx, "v1/id", map[string]string{"character_name": characterName}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetCharacterBasic(ctx context.Context, ocid string, date any) (*CharacterBasic, error) {
	dateStr, err := c.toDateString(date)
	if err != nil {
		return nil, err
	}
	raw, err := c.fetchRaw(ctx, "v1/character/basic", map[string]string{"ocid": ocid, "date": dateStr})
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

func (c *Client) GetCharacterImage(ctx context.Context, ocid string, opts *CharacterImageOptions, date any) (*CharacterImage, error) {
	basic, err := c.GetCharacterBasic(ctx, ocid, date)
	if err != nil || basic == nil {
		return nil, err
	}
	action := common.CharacterImageActionStand1
	emotion := common.CharacterImageEmotionDefault
	wmotion := common.CharacterImageWeaponMotionDefault
	actionFrame := 0
	emotionFrame := 0
	width := defaultImageSize
	height := defaultImageSize
	x := defaultImageX
	y := defaultImageY
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
		if opts.X != nil {
			x = *opts.X
		}
		if opts.Y != nil {
			y = *opts.Y
		}
	}
	baseURL := common.RemoveQuery(basic.CharacterImage)
	query := map[string]string{
		"action":  fmt.Sprintf("%s.%d", action, actionFrame),
		"emotion": fmt.Sprintf("%s.%d", emotion, emotionFrame),
		"wmotion": string(wmotion),
		"width":   strconv.Itoa(width),
		"height":  strconv.Itoa(height),
		"x":       strconv.Itoa(x),
		"y":       strconv.Itoa(y),
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
		X:            &x,
		Y:            &y,
	}, nil
}

func (c *Client) GetCharacterPopularity(ctx context.Context, ocid string, date any) (*CharacterPopularity, error) {
	return fetchWithEmpty[CharacterPopularity](ctx, c, "v1/character/popularity", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterStat(ctx context.Context, ocid string, date any) (*CharacterStat, error) {
	return fetchWithEmpty[CharacterStat](ctx, c, "v1/character/stat", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterHyperStat(ctx context.Context, ocid string, date any) (*CharacterHyperStat, error) {
	return fetchWithEmpty[CharacterHyperStat](ctx, c, "v1/character/hyper-stat", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterPropensity(ctx context.Context, ocid string, date any) (*CharacterPropensity, error) {
	return fetchWithEmpty[CharacterPropensity](ctx, c, "v1/character/propensity", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterAbility(ctx context.Context, ocid string, date any) (*CharacterAbility, error) {
	return fetchWithEmpty[CharacterAbility](ctx, c, "v1/character/ability", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterItemEquipment(ctx context.Context, ocid string, date any) (*CharacterItemEquipment, error) {
	return fetchWithEmpty[CharacterItemEquipment](ctx, c, "v1/character/item-equipment", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterCashItemEquipment(ctx context.Context, ocid string, date any) (*CharacterCashItemEquipment, error) {
	return fetchWithEmpty[CharacterCashItemEquipment](ctx, c, "v1/character/cashitem-equipment", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterSymbolEquipment(ctx context.Context, ocid string, date any) (*CharacterSymbolEquipment, error) {
	return fetchWithEmpty[CharacterSymbolEquipment](ctx, c, "v1/character/symbol-equipment", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterSetEffect(ctx context.Context, ocid string, date any) (*CharacterSetEffect, error) {
	return fetchWithEmpty[CharacterSetEffect](ctx, c, "v1/character/set-effect", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterBeautyEquipment(ctx context.Context, ocid string, date any) (*CharacterBeautyEquipment, error) {
	return fetchWithEmpty[CharacterBeautyEquipment](ctx, c, "v1/character/beauty-equipment", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterAndroidEquipment(ctx context.Context, ocid string, date any) (*CharacterAndroidEquipment, error) {
	return fetchWithEmpty[CharacterAndroidEquipment](ctx, c, "v1/character/android-equipment", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterPetEquipment(ctx context.Context, ocid string, date any) (*CharacterPetEquipment, error) {
	return fetchWithEmpty[CharacterPetEquipment](ctx, c, "v1/character/pet-equipment", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterSkill(ctx context.Context, ocid, characterSkillGrade string, date any) (*CharacterSkill, error) {
	return fetchWithEmpty[CharacterSkill](ctx, c, "v1/character/skill", map[string]string{"ocid": ocid, "character_skill_grade": characterSkillGrade, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterLinkSkill(ctx context.Context, ocid string, date any) (*CharacterLinkSkill, error) {
	return fetchWithEmpty[CharacterLinkSkill](ctx, c, "v1/character/link-skill", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterVMatrix(ctx context.Context, ocid string, date any) (*CharacterVMatrix, error) {
	return fetchWithEmpty[CharacterVMatrix](ctx, c, "v1/character/vmatrix", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterHexaMatrix(ctx context.Context, ocid string, date any) (*CharacterHexaMatrix, error) {
	return fetchWithEmpty[CharacterHexaMatrix](ctx, c, "v1/character/hexamatrix", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterHexaMatrixStat(ctx context.Context, ocid string, date any) (*CharacterHexaMatrixStat, error) {
	return fetchWithEmpty[CharacterHexaMatrixStat](ctx, c, "v1/character/hexamatrix-stat", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterDojang(ctx context.Context, ocid string, date any) (*CharacterDojang, error) {
	return fetchWithEmpty[CharacterDojang](ctx, c, "v1/character/dojang", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetCharacterOtherStat(ctx context.Context, ocid string, date any) (CharacterOtherStat, error) {
	raw, err := c.fetchRaw(ctx, "v1/character/another-name-stat", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
	if err != nil {
		return nil, err
	}
	if empty, err := isEmpty(raw); err != nil {
		return nil, err
	} else if empty {
		return nil, nil
	}
	var resp CharacterOtherStat
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetCharacterRingExchangeSkillEquipment(ctx context.Context, ocid string, date any) (CharacterRingExchangeSkillEquipment, error) {
	raw, err := c.fetchRaw(ctx, "v1/character/ring-power", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
	if err != nil {
		return nil, err
	}
	if empty, err := isEmpty(raw); err != nil {
		return nil, err
	} else if empty {
		return nil, nil
	}
	var resp CharacterRingExchangeSkillEquipment
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Union

func (c *Client) GetUnion(ctx context.Context, ocid string, date any) (*Union, error) {
	return fetchWithEmpty[Union](ctx, c, "v1/user/union", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetUnionRaider(ctx context.Context, ocid string, date any) (*UnionRaider, error) {
	return fetchWithEmpty[UnionRaider](ctx, c, "v1/user/union-raider", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetUnionArtifact(ctx context.Context, ocid string, date any) (*UnionArtifact, error) {
	return fetchWithEmpty[UnionArtifact](ctx, c, "v1/user/union-artifact", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

func (c *Client) GetUnionChampion(ctx context.Context, ocid string, date any) (*UnionChampion, error) {
	return fetchWithEmpty[UnionChampion](ctx, c, "v1/user/union-champion", map[string]string{"ocid": ocid, "date": c.dateOrEmpty(date)})
}

// Guild

func (c *Client) GetGuild(ctx context.Context, guildName, worldName string) (*Guild, error) {
	raw, err := c.fetchRaw(ctx, "v1/guild/id", map[string]string{"guild_name": guildName, "world_name": worldName})
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

func (c *Client) GetGuildBasic(ctx context.Context, guildID string, date any) (*GuildBasic, error) {
	return fetchWithEmpty[GuildBasic](ctx, c, "v1/guild/basic", map[string]string{"oguild_id": guildID, "date": c.dateOrEmpty(date)})
}

// History

type HistoryOptions struct {
	Date   any
	Cursor string
}

func (c *Client) GetCubeHistory(ctx context.Context, count int, opts *HistoryOptions) (*CubeHistoryResponse, error) {
	query := map[string]string{"count": strconv.Itoa(count)}
	if opts != nil {
		if opts.Cursor != "" {
			query["cursor"] = opts.Cursor
		}
		if opts.Date != nil {
			query["date"] = c.dateOrEmpty(opts.Date)
		}
	}
	var resp CubeHistoryResponse
	if err := c.Fetch(ctx, "v1/history/cube", query, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPotentialHistory(ctx context.Context, count int, opts *HistoryOptions) (*PotentialHistoryResponse, error) {
	query := map[string]string{"count": strconv.Itoa(count)}
	if opts != nil {
		if opts.Cursor != "" {
			query["cursor"] = opts.Cursor
		}
		if opts.Date != nil {
			query["date"] = c.dateOrEmpty(opts.Date)
		}
	}
	var resp PotentialHistoryResponse
	if err := c.Fetch(ctx, "v1/history/potential", query, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetStarforceHistory(ctx context.Context, count int, opts *HistoryOptions) (*StarforceHistoryResponse, error) {
	query := map[string]string{"count": strconv.Itoa(count)}
	if opts != nil {
		if opts.Cursor != "" {
			query["cursor"] = opts.Cursor
		}
		if opts.Date != nil {
			query["date"] = c.dateOrEmpty(opts.Date)
		}
	}
	var resp StarforceHistoryResponse
	if err := c.Fetch(ctx, "v1/history/starforce", query, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Ranking

type RankingFilter struct {
	WorldName      string
	WorldType      string
	CharacterClass string
	OCID           string
	Page           string
	Region         string
}

func (c *Client) GetOverallRanking(ctx context.Context, filter *RankingFilter, date any) (*OverallRankingResponse, error) {
	query := map[string]string{"date": c.dateOrEmpty(date)}
	applyFilter(query, filter)
	return fetchRanking(ctx, c, "v1/ranking/overall", query)
}

func (c *Client) GetUnionRanking(ctx context.Context, filter *RankingFilter, date any) (*UnionRankingResponse, error) {
	query := map[string]string{"date": c.dateOrEmpty(date)}
	applyFilter(query, filter)
	return fetchRanking(ctx, c, "v1/ranking/union", query)
}

func (c *Client) GetGuildRanking(ctx context.Context, filter *RankingFilter, date any) (*GuildRankingResponse, error) {
	query := map[string]string{"date": c.dateOrEmpty(date)}
	applyFilter(query, filter)
	return fetchRanking(ctx, c, "v1/ranking/guild", query)
}

func (c *Client) GetDojangRanking(ctx context.Context, filter *RankingFilter, date any) (*DojangRankingResponse, error) {
	query := map[string]string{"date": c.dateOrEmpty(date)}
	applyFilter(query, filter)
	return fetchRanking(ctx, c, "v1/ranking/dojang", query)
}

func (c *Client) GetSeedRanking(ctx context.Context, filter *RankingFilter, date any) (*SeedRankingResponse, error) {
	query := map[string]string{"date": c.dateOrEmpty(date)}
	applyFilter(query, filter)
	return fetchRanking(ctx, c, "v1/ranking/the-seed", query)
}

func (c *Client) GetAchievementRanking(ctx context.Context, filter *RankingFilter, date any) (*AchievementRankingResponse, error) {
	query := map[string]string{"date": c.dateOrEmpty(date)}
	applyFilter(query, filter)
	return fetchRanking(ctx, c, "v1/ranking/achievement", query)
}

// Notices

func (c *Client) GetNoticeList(ctx context.Context) (*NoticeList, error) {
	var resp NoticeList
	if err := c.Fetch(ctx, "v1/notice", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetNoticeDetail(ctx context.Context, noticeID int) (*NoticeDetail, error) {
	var resp NoticeDetail
	if err := c.Fetch(ctx, "v1/notice/"+strconv.Itoa(noticeID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUpdateNoticeList(ctx context.Context) (*UpdateNoticeList, error) {
	var resp UpdateNoticeList
	if err := c.Fetch(ctx, "v1/notice/update", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUpdateNoticeDetail(ctx context.Context, noticeID int) (*UpdateNoticeDetail, error) {
	var resp UpdateNoticeDetail
	if err := c.Fetch(ctx, "v1/notice/update/"+strconv.Itoa(noticeID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetEventNoticeList(ctx context.Context) (*EventNoticeList, error) {
	var resp EventNoticeList
	if err := c.Fetch(ctx, "v1/notice/event", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetEventNoticeDetail(ctx context.Context, noticeID int) (*EventNoticeDetail, error) {
	var resp EventNoticeDetail
	if err := c.Fetch(ctx, "v1/notice/event/"+strconv.Itoa(noticeID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetCashshopNoticeList(ctx context.Context) (*CashshopNoticeList, error) {
	var resp CashshopNoticeList
	if err := c.Fetch(ctx, "v1/notice/cashshop", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetCashshopNoticeDetail(ctx context.Context, noticeID int) (*CashshopNoticeDetail, error) {
	var resp CashshopNoticeDetail
	if err := c.Fetch(ctx, "v1/notice/cashshop/"+strconv.Itoa(noticeID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Inspection

func (c *Client) GetInspectionInfo(ctx context.Context) (*InspectionInfo, error) {
	var resp json.RawMessage
	if err := c.Fetch(ctx, "v1/inspection", nil, &resp); err != nil {
		return nil, err
	}
	return &InspectionInfo{Raw: resp}, nil
}

// helpers

func applyFilter(query map[string]string, filter *RankingFilter) {
	if filter == nil {
		return
	}
	if filter.WorldName != "" {
		query["world_name"] = filter.WorldName
	}
	if filter.WorldType != "" {
		query["world_type"] = filter.WorldType
	}
	if filter.CharacterClass != "" {
		query["class"] = filter.CharacterClass
	}
	if filter.OCID != "" {
		query["ocid"] = filter.OCID
	}
	if filter.Page != "" {
		query["page"] = filter.Page
	}
	if filter.Region != "" {
		query["region"] = filter.Region
	}
}

func fetchRanking(ctx context.Context, c *Client, path string, query map[string]string) (*RankingResponse, error) {
	var resp RankingResponse
	if err := c.Fetch(ctx, path, query, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

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

func (c *Client) dateOrEmpty(value any) string {
	if value == nil {
		return ""
	}
	str, err := c.toDateString(value)
	if err != nil {
		return ""
	}
	return str
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
