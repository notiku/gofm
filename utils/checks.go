package utils

import (
	"regexp"
)

// Checks if the given string is a valid MBID (UUID format).
func IsValidMBID(identifier string) bool {
	matched, _ := regexp.MatchString("^[a-f0-9-]{36}$", identifier)
	return matched
}
