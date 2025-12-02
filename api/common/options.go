package common

// CharacterImageOptions defines the available options for generating character images.
// Zero values use server defaults; X and Y act as optional offsets.
type CharacterImageOptions struct {
	Action       CharacterImageAction
	Emotion      CharacterImageEmotion
	WeaponMotion CharacterImageWeaponMotion
	ActionFrame  int
	EmotionFrame int
	Width        int
	Height       int
	X            *int
	Y            *int
}
