package tms

import "maplestory-openapi/api/common"

// Region-wide constants.
const (
	subPath          = "maplestorytw"
	timezoneOffset   = 480 // TST (UTC+8)
	minimumAPIDateY  = 2025
	minimumAPIDateM  = 10
	minimumAPIDateD  = 15
	defaultImageSize = 96
)

var minimumAPIDate = common.Date{Year: minimumAPIDateY, Month: minimumAPIDateM, Day: minimumAPIDateD}
