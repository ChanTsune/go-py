package strings

import "unicode"

// IsAlpha Return True if all characters in the string are alphabetic and there is at least one character, False otherwise. Alphabetic characters are those characters defined in the Unicode character database as “Letter”, i.e., those with general category property being one of “Lm”, “Lt”, “Lu”, “Ll”, or “Lo”. Note that this is different from the “Alphabetic” property defined in the Unicode Standard.
func IsAlpha(s string) bool {
	return isX(s, false, isAlpha)
}

func isAlpha(r rune) bool {
	return unicode.Is(unicode.Letter, r)
}
