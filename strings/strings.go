package strings

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// lastExplode splits s into a slice of UTF-8 strings,
// one string per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 sequences become correct encodings of U+FFFD.
func lastExplode(s string, n int) []string {
	r := []rune(s)
	l := len(r)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := n - 1; i > 0; i-- {
		a[i] = string(r[len(r)-1])
		r = r[:len(r)-1]
	}
	if n > 0 {
		a[0] = string(r)
	}
	return a
}

// Length Return rune count.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// SwapCase Return swap cased string.
func SwapCase(s string) string {
	var buf bytes.Buffer
	buf.Grow(len(s))
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
	if fillLen := width - Length(s); fillLen >= 0 {
		return strings.Repeat(string(fillChar), fillLen) + s
	}
	return s
}

// LJust Return padded string.
func LJust(s string, width int, fillChar rune) string {
	if fillLen := width - Length(s); fillLen >= 0 {
		return s + strings.Repeat(string(fillChar), fillLen)
	}
	return s
}

// Center Return padded string.
func Center(s string, width int, fillChar rune) string {
	if fillLen := width - Length(s); fillLen >= 0 {
		l := fillLen/2 + (fillLen & width & 1)
		r := fillLen - l
		f := string(fillChar)
		return strings.Repeat(f, l) + s + strings.Repeat(f, r)
	}
	return s
}

// Partition Return string.
func Partition(s, sep string) (string, string, string) {
	if v := strings.SplitN(s, sep, 2); len(v) == 2 {
		return v[0], sep, v[1]
	}
	return s, "", ""
}

// RPartition Return string.
func RPartition(s, sep string) (string, string, string) {
	if v := LastSplitN(s, sep, 2); len(v) == 2 {
		return v[0], sep, v[1]
	}
	return "", "", s
}

func genLastSplit(s, sep string, sepSave, n int) []string {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return lastExplode(s, n)
	}
	if n < 0 {
		n = strings.Count(s, sep) + 1
	}

	a := make([]string, n)
	n--
	i := 0
	sepLen := len(sep)
	for i < n {
		m := strings.LastIndex(s[:len(s)-sepSave], sep)
		if m < 0 {
			break
		}
		a[n] = s[m+sepLen:]
		s = s[:m+sepSave]
		n--
	}
	a[n] = s
	return a[n:]
}

// LastSplitN Return
func LastSplitN(s, sep string, n int) []string {
	return genLastSplit(s, sep, 0, n)
}

// LastSplit Return
func LastSplit(s, sep string) []string {
	return genLastSplit(s, sep, 0, -1)
}

// LastSplitAfterN Return
func LastSplitAfterN(s, sep string, n int) []string {
	return genLastSplit(s, sep, len(sep), n)
}

// LastSplitAfter Return
func LastSplitAfter(s, sep string) []string {
	return genLastSplit(s, sep, len(sep), -1)
}
