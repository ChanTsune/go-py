package strings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Capitalize Return a copy of the string with its first character capitalized and the rest lowercased.
func Capitalize(s string) string {
	length := Length(s)
	if length == 0 {
		return s
	} else if length == 1 {
		return strings.ToUpper(s)
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + strings.ToLower(s[size:])
}
