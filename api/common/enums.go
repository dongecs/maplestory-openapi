package common

// PotentialOptionGrade represents the potential option grade.
type PotentialOptionGrade int

const (
	PotentialOptionGradeRare PotentialOptionGrade = iota
	PotentialOptionGradeEpic
	PotentialOptionGradeUnique
	PotentialOptionGradeLegendary
)

func (p PotentialOptionGrade) String() string {
	switch p {
	case PotentialOptionGradeRare:
		return "RARE"
	case PotentialOptionGradeEpic:
		return "EPIC"
	case PotentialOptionGradeUnique:
		return "UNIQUE"
	case PotentialOptionGradeLegendary:
		return "LEGENDARY"
	default:
		return "UNKNOWN"
	}
}

// CharacterImageAction enumerates possible character actions for image rendering.
type CharacterImageAction string

const (
	CharacterImageActionStand1         CharacterImageAction = "A00"
	CharacterImageActionStand2         CharacterImageAction = "A01"
	CharacterImageActionWalk1          CharacterImageAction = "A02"
	CharacterImageActionWalk2          CharacterImageAction = "A03"
	CharacterImageActionProne          CharacterImageAction = "A04"
	CharacterImageActionFly            CharacterImageAction = "A05"
	CharacterImageActionJump           CharacterImageAction = "A06"
	CharacterImageActionSit            CharacterImageAction = "A07"
	CharacterImageActionLadder         CharacterImageAction = "A08"
	CharacterImageActionRope           CharacterImageAction = "A09"
	CharacterImageActionHeal           CharacterImageAction = "A10"
	CharacterImageActionAlert          CharacterImageAction = "A11"
	CharacterImageActionProneStab      CharacterImageAction = "A12"
	CharacterImageActionSwingO1        CharacterImageAction = "A13"
	CharacterImageActionSwingO2        CharacterImageAction = "A14"
	CharacterImageActionSwingO3        CharacterImageAction = "A15"
	CharacterImageActionSwingOF        CharacterImageAction = "A16"
	CharacterImageActionSwingP1        CharacterImageAction = "A17"
	CharacterImageActionSwingP2        CharacterImageAction = "A18"
	CharacterImageActionSwingPF        CharacterImageAction = "A19"
	CharacterImageActionSwingT1        CharacterImageAction = "A20"
	CharacterImageActionSwingT2        CharacterImageAction = "A21"
	CharacterImageActionSwingT3        CharacterImageAction = "A22"
	CharacterImageActionSwingTF        CharacterImageAction = "A23"
	CharacterImageActionStabO1         CharacterImageAction = "A24"
	CharacterImageActionStabO2         CharacterImageAction = "A25"
	CharacterImageActionStabOF         CharacterImageAction = "A26"
	CharacterImageActionStabT1         CharacterImageAction = "A27"
	CharacterImageActionStabT2         CharacterImageAction = "A28"
	CharacterImageActionStabTF         CharacterImageAction = "A29"
	CharacterImageActionShoot1         CharacterImageAction = "A30"
	CharacterImageActionShoot2         CharacterImageAction = "A31"
	CharacterImageActionShootF         CharacterImageAction = "A32"
	CharacterImageActionDead           CharacterImageAction = "A33"
	CharacterImageActionGhostWalk      CharacterImageAction = "A34"
	CharacterImageActionGhostStand     CharacterImageAction = "A35"
	CharacterImageActionGhostJump      CharacterImageAction = "A36"
	CharacterImageActionGhostProneStab CharacterImageAction = "A37"
	CharacterImageActionGhostLadder    CharacterImageAction = "A38"
	CharacterImageActionGhostRope      CharacterImageAction = "A39"
	CharacterImageActionGhostFly       CharacterImageAction = "A40"
	CharacterImageActionGhostSit       CharacterImageAction = "A41"
)

// CharacterImageEmotion enumerates character emotions for image rendering.
type CharacterImageEmotion string

const (
	CharacterImageEmotionDefault    CharacterImageEmotion = "E00"
	CharacterImageEmotionWink       CharacterImageEmotion = "E01"
	CharacterImageEmotionSmile      CharacterImageEmotion = "E02"
	CharacterImageEmotionCry        CharacterImageEmotion = "E03"
	CharacterImageEmotionAngry      CharacterImageEmotion = "E04"
	CharacterImageEmotionBewildered CharacterImageEmotion = "E05"
	CharacterImageEmotionBlink      CharacterImageEmotion = "E06"
	CharacterImageEmotionBlaze      CharacterImageEmotion = "E07"
	CharacterImageEmotionBowing     CharacterImageEmotion = "E08"
	CharacterImageEmotionCheers     CharacterImageEmotion = "E09"
	CharacterImageEmotionChu        CharacterImageEmotion = "E10"
	CharacterImageEmotionDam        CharacterImageEmotion = "E11"
	CharacterImageEmotionDespair    CharacterImageEmotion = "E12"
	CharacterImageEmotionGlitter    CharacterImageEmotion = "E13"
	CharacterImageEmotionHit        CharacterImageEmotion = "E14"
	CharacterImageEmotionHot        CharacterImageEmotion = "E15"
	CharacterImageEmotionHum        CharacterImageEmotion = "E16"
	CharacterImageEmotionLove       CharacterImageEmotion = "E17"
	CharacterImageEmotionOops       CharacterImageEmotion = "E18"
	CharacterImageEmotionPain       CharacterImageEmotion = "E19"
	CharacterImageEmotionTroubled   CharacterImageEmotion = "E20"
	CharacterImageEmotionQBlue      CharacterImageEmotion = "E21"
	CharacterImageEmotionShine      CharacterImageEmotion = "E22"
	CharacterImageEmotionStunned    CharacterImageEmotion = "E23"
	CharacterImageEmotionVomit      CharacterImageEmotion = "E24"
)

// CharacterImageWeaponMotion enumerates weapon motions for image rendering.
type CharacterImageWeaponMotion string

const (
	CharacterImageWeaponMotionDefault  CharacterImageWeaponMotion = "W00"
	CharacterImageWeaponMotionOneHand  CharacterImageWeaponMotion = "W01"
	CharacterImageWeaponMotionTwoHands CharacterImageWeaponMotion = "W02"
	CharacterImageWeaponMotionGun      CharacterImageWeaponMotion = "W03"
	CharacterImageWeaponMotionNothing  CharacterImageWeaponMotion = "W04"
)
