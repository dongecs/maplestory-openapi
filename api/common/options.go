package common

// CharacterImageOptions defines the available options for generating character images.
type CharacterImageOptions struct {
	Action       CharacterImageAction       `json:"action,omitempty"`
	Emotion      CharacterImageEmotion      `json:"emotion,omitempty"`
	WeaponMotion CharacterImageWeaponMotion `json:"wmotion,omitempty"`
	ActionFrame  *int                       `json:"action_frame,omitempty"`
	EmotionFrame *int                       `json:"emotion_frame,omitempty"`
}
