package strings

// IsAlNum Return True if all characters in the string are alphanumeric and there is at least one character, False otherwise. A character c is alphanumeric if one of the following returns True: c.isalpha(), c.isdecimal(), c.isdigit(), or c.isnumeric().
func IsAlNum(s string) bool {
	return isX(s, false, isAlNum)
}
func isAlNum(r rune) bool {
	return isAlpha(r) || isDecimal(r) || isDigit(r) || isNumeric(r)
}
