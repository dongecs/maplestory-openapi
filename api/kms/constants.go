package kms

import "github.com/dongecs/maplestory-openapi/api/common"

const (
	subPath          = "maplestory"
	timezoneOffset   = 540 // KST (UTC+9)
	minimumAPIDateY  = 2023
	minimumAPIDateM  = 12
	minimumAPIDateD  = 21
	defaultImageSize = 300
	defaultImageX    = 150
	defaultImageY    = 200
)

var minimumAPIDate = common.Date{Year: minimumAPIDateY, Month: minimumAPIDateM, Day: minimumAPIDateD}
