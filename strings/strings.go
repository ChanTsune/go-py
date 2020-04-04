package strings

import (
	"bytes"
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
