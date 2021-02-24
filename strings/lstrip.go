package strings

import (
	"strings"
	"unicode"
)

// LStrip Return a copy of the string with leading characters removed. The chars argument is a string specifying the set of characters to be removed. The chars argument is not a prefix; rather, all combinations of its values are stripped:
func LStrip(s, chars string) string {
	return strings.TrimLeft(s, chars)
}

// LStripWhiteSpace Return a copy of the string with leading whitespaces removed.
func LStripWhiteSpace(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}
