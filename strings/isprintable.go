package strings

import "unicode"

// IsPrintable Return True if all characters in the string are printable or the string is empty, False otherwise. Nonprintable characters are those characters defined in the Unicode character database as “Other” or “Separator”, excepting the ASCII space (0x20) which is considered printable.
func IsPrintable(s string) bool {
	return isX(s, true, unicode.IsPrint)
}
