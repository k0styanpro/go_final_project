package api

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dateStr, repeat string) (string, error) {
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return "", errors.New("invalid date")
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("empty repeat rule")
	}

	switch parts[0] {
	case "d":
		return handleDaily(now, date, parts)
	case "y":
		return handleYearly(now, date)
	case "w":
		return handleWeekly(now, date, parts)
	case "m":
		return handleMonthly(now, date, parts)
	default:
		return "", errors.New("unsupported rule")
	}
}
