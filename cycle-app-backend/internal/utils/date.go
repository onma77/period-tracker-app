package utils

import (
	"errors"
	"time"
)

const dateLayout = "2006-01-02"

func ParseDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, errors.New("date is required")
	}

	parsedDate, err := time.Parse(dateLayout, date)
	if err != nil {
		return time.Time{}, errors.New("date must be in YYYY-MM-DD format")
	}

	return parsedDate, nil
}

func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}
