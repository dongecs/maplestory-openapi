package kms

import "encoding/json"

type CubeHistoryResponse struct {
	Count       int             `json:"count"`
	CubeHistory json.RawMessage `json:"cube_history"`
	NextCursor  *string         `json:"next_cursor"`
}

type PotentialHistoryResponse struct {
	Count            int             `json:"count"`
	PotentialHistory json.RawMessage `json:"potential_history"`
	NextCursor       *string         `json:"next_cursor"`
}

type StarforceHistoryResponse struct {
	Count            int             `json:"count"`
	StarforceHistory json.RawMessage `json:"starforce_history"`
	NextCursor       *string         `json:"next_cursor"`
}
