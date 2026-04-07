package utils

import(
	"strings"
)
var validDays = map[string]bool{
	"Monday":    true,
	"Tuesday":   true,
	"Wednesday": true,
	"Thursday":  true,
	"Friday":    true,
	"Saturday":  true,
	"Sunday":    true,
}

func IsValidDay(day string) (string, bool) {
	// normalize input → "monday" → "Monday"
	normalized := strings.Title(strings.ToLower(day))

	if validDays[normalized] {
		return normalized, true
	}
	return "", false
}