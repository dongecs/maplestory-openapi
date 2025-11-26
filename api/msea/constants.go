package msea

import "github.com/dongecs/maplestory-openapi/api/common"

const (
	subPath          = "maplestorysea"
	timezoneOffset   = 480 // SGT (UTC+8)
	minimumAPIDateY  = 2025
	minimumAPIDateM  = 4
	minimumAPIDateD  = 20
	defaultImageSize = 96
)

var minimumAPIDate = common.Date{Year: minimumAPIDateY, Month: minimumAPIDateM, Day: minimumAPIDateD}
