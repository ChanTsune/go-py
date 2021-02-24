package strings

import "unicode"

// IsLower Return True if all cased characters in the string are lowercase and there is at least one cased character, False otherwise.
func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	hasCase := false
	for _, r := range s {
		if isCased(r) {
			if !unicode.IsLower(r) {
				return false
			}
			hasCase = true
		}
	}
	return hasCase
}
