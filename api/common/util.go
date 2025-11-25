package common

import (
	"fmt"
	"reflect"
	"strings"
)

// PotentialOptionGradeFromString converts the Korean grade labels used by the API into the enum.
func PotentialOptionGradeFromString(text string) (PotentialOptionGrade, error) {
	switch strings.TrimSpace(text) {
	case "레어":
		return PotentialOptionGradeRare, nil
	case "에픽":
		return PotentialOptionGradeEpic, nil
	case "유니크":
		return PotentialOptionGradeUnique, nil
	case "레전드리":
		return PotentialOptionGradeLegendary, nil
	default:
		return 0, fmt.Errorf("no enum constant for string: %s", text)
	}
}

// RemoveQuery strips the query string from a URL.
func RemoveQuery(url string) string {
	if url == "" {
		return url
	}

	if idx := strings.IndexRune(url, '?'); idx >= 0 {
		return url[:idx]
	}

	return url
}

// IsEmptyResponse checks whether an API response payload only contains the date field.
// This mirrors the behavior used in other language bindings to detect responses where
// the requested date predates available data.
func IsEmptyResponse(body map[string]any) bool {
	for key, value := range body {
		if key == "date" {
			continue
		}

		if value == nil {
			continue
		}

		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Slice && v.Len() == 0 {
			continue
		}

		return false
	}

	return true
}
