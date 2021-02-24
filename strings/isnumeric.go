package strings

import "unicode"

// IsNumeric Return True if all characters in the string are numeric characters, and there is at least one character, False otherwise. Numeric characters include digit characters, and all characters that have the Unicode numeric value property, e.g. U+2155, VULGAR FRACTION ONE FIFTH. Formally, numeric characters are those with the property value Numeric_Type=Digit, Numeric_Type=Decimal or Numeric_Type=Numeric.
func IsNumeric(s string) bool {
	return isX(s, false, isNumeric)
}

func isNumeric(r rune) bool {
	return unicode.IsOneOf([]*unicode.RangeTable{
		unicode.N,
		unicode.Number,
	}, r)
}
