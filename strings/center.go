package strings

import "strings"

// Center Return centered in a string of length width. Padding is done using the specified fillchar (default is an ASCII space). The original string is returned if width is less than or equal to len(s).
func Center(s string, width int, fillChar rune) string {
	if fillLen := width - Length(s); fillLen >= 0 {
		l := fillLen/2 + (fillLen & width & 1)
		r := fillLen - l
		f := string(fillChar)
		return strings.Repeat(f, l) + s + strings.Repeat(f, r)
	}
	return s
}
