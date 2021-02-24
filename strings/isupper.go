package strings

import "unicode"

// IsUpper Return True if all cased characters in the string are uppercase and there is at least one cased character, False otherwise.
func IsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	hasCase := false
	for _, r := range s {
		if isCased(r) {
			if !unicode.IsUpper(r) {
				return false
			}
			hasCase = true
		}
	}
	return hasCase
}
