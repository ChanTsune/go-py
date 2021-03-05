package strings

import "strings"

// RJust Return the string right justified in a string of length width. Padding is done using the specified fillchar (default is an ASCII space). The original string is returned if width is less than or equal to len(s).
func RJust(s string, width int, fillChar rune) string {
	if fillLen := width - Length(s); fillLen >= 0 {
		return strings.Repeat(string(fillChar), fillLen) + s
	}
	return s
}
