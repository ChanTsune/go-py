package strings

import (
	"bytes"
	"strings"
	"unicode"
)

// SwapCase Return swap cased string.
func SwapCase(s string) string {
	var buf bytes.Buffer
	for _, x := range s {
		u := unicode.ToUpper(x)
		l := unicode.ToLower(x)
		if u == x {
			buf.WriteRune(l)
		} else if l == x {
			buf.WriteRune(u)
		} else {
			buf.WriteRune(x)
		}
	}
	return buf.String()
}

// RJust Return padded string.
func RJust(s string, width int, fillChar rune) string {
	if fillLen := width - len(s); fillLen >= 0 {
		return strings.Repeat(string(fillChar), fillLen) + s
	}
	return s
}

// LJust Return padded string.
func LJust(s string, width int, fillChar rune) string {
	if fillLen := width - len(s); fillLen >= 0 {
		return s + strings.Repeat(string(fillChar), fillLen)
	}
	return s
}

// Center Return padded string.
func Center(s string, width int, fillChar rune) string {
	if fillLen := width - len(s); fillLen >= 0 {
		r := fillLen / 2
		l := fillLen - r
		f := string(fillChar)
		return strings.Repeat(f, l) + s + strings.Repeat(f, r)
	}
	return s
}
