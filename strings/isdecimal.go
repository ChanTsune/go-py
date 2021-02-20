package strings

import "unicode"

// IsDecimal Return True if all characters in the string are decimal characters and there is at least one character, False otherwise. Decimal characters are those that can be used to form numbers in base 10, e.g. U+0660, ARABIC-INDIC DIGIT ZERO. Formally a decimal character is a character in the Unicode General Category “Nd”.
func IsDecimal(s string) bool {
	return isX(s, false, isDecimal)
}

func isDecimal(r rune) bool {
	return unicode.Is(unicode.Nd, r)
}
