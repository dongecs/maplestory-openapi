package common

import (
	"fmt"
	"time"
)

// Date represents a year-month-day tuple used by the API.
type Date struct {
	Year  int
	Month int
	Day   int
}

// LatestAPIUpdateTimeOptions expresses the server update window used to compute the latest retrievable date.
type LatestAPIUpdateTimeOptions struct {
	Hour       int
	Minute     int
	DateOffset int
}

// ToTime converts the Date into a time.Time in the provided timezone.
func (d Date) ToTime(loc *time.Location) time.Time {
	if loc == nil {
		loc = time.UTC
	}

	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, loc)
}

// FixedOffsetLocation returns a timezone with a fixed offset in minutes (e.g., 540 -> UTC+09:00).
func FixedOffsetLocation(offsetMinutes int) *time.Location {
	abs := offsetMinutes
	if abs < 0 {
		abs = -abs
	}
	hours := abs / 60
	minutes := abs % 60

	sign := "+"
	if offsetMinutes < 0 {
		sign = "-"
	}

	return time.FixedZone(fmt.Sprintf("UTC%s%02d:%02d", sign, hours, minutes), offsetMinutes*60)
}

// ToDateString converts either a time.Time or a Date into the YYYY-MM-DD string the API expects.
// When min is provided, it enforces that the target date is not earlier than the allowed minimum.
func ToDateString(value any, loc *time.Location, min *Date) (string, error) {
	target, err := normalizeDate(value, loc)
	if err != nil {
		return "", err
	}

	if min != nil {
		minDate, err := normalizeDate(*min, loc)
		if err != nil {
			return "", err
		}

		if target.Before(minDate) {
			return "", fmt.Errorf("you can only retrieve data after %s", minDate.Format("2006-01-02"))
		}
	}

	return target.Format("2006-01-02"), nil
}

// GetProperDefaultDate calculates the most recent available date based on the API's refresh window.
// It mirrors the logic used by the other language bindings: if the current time is earlier than the
// daily update time, the previous day is used, then the optional DateOffset is applied.
func GetProperDefaultDate(now time.Time, loc *time.Location, opts LatestAPIUpdateTimeOptions) Date {
	if loc == nil {
		loc = time.UTC
	}

	localNow := now.In(loc)
	updateTime := time.Date(localNow.Year(), localNow.Month(), localNow.Day(), opts.Hour, opts.Minute, 0, 0, loc)

	adjusted := localNow
	if !localNow.After(updateTime) {
		adjusted = adjusted.AddDate(0, 0, -1)
	}

	if opts.DateOffset > 0 {
		adjusted = adjusted.AddDate(0, 0, -opts.DateOffset)
	}

	return Date{
		Year:  adjusted.Year(),
		Month: int(adjusted.Month()),
		Day:   adjusted.Day(),
	}
}

func normalizeDate(value any, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.UTC
	}

	switch v := value.(type) {
	case time.Time:
		// Use only the date component; normalize to the provided timezone.
		return time.Date(v.In(loc).Year(), v.In(loc).Month(), v.In(loc).Day(), 0, 0, 0, 0, loc), nil
	case Date:
		return v.ToTime(loc), nil
	case *Date:
		if v == nil {
			return time.Time{}, fmt.Errorf("nil *Date provided")
		}
		return v.ToTime(loc), nil
	default:
		return time.Time{}, fmt.Errorf("unsupported date type %T", value)
	}
}
