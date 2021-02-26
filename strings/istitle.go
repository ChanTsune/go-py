package strings

import "unicode"

// IsTitle Return True if the string is a titlecased string and there is at least one character, for example uppercase characters may only follow uncased characters and lowercase characters only cased ones. Return False otherwise.
func IsTitle(s string) bool {
	if len(s) == 0 {
		return false
	}
	var cased = false
	var previousIsCased = false
	for _, ch := range s {
		if unicode.IsUpper(ch) || unicode.IsTitle(ch) {
			if previousIsCased {
				return false
			}
			previousIsCased = true
			cased = true
		} else if unicode.IsLower(ch) {
			if !previousIsCased {
				return false
			}
			previousIsCased = true
			cased = true
		} else {
			previousIsCased = false
		}
	}
	return cased
}
