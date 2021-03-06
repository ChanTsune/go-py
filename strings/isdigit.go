package strings

import (
	"unicode"
)

// IsDigit Return True if all characters in the string are digits and there is at least one character, False otherwise. Digits include decimal characters and digits that need special handling, such as the compatibility superscript digits. This covers digits which cannot be used to form numbers in base 10, like the Kharosthi numbers. Formally, a digit is a character that has the property value Numeric_Type=Digit or Numeric_Type=Decimal.
func IsDigit(s string) bool {
	return isX(s, false, isDecimal)
}

func isDigit(r rune) bool {
	return unicode.In(r, unicode.N)
}
