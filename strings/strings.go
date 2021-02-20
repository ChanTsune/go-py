package strings

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
)

// Arg slice args
type Arg struct {
	Start int
	End   int
}

// Option option arg
type Option func(*Arg)

// Start option argument named start
func Start(start int) Option {
	return func(args *Arg) {
		args.Start = start
	}
}

// End option argument named end
func End(end int) Option {
	return func(args *Arg) {
		args.End = end
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func adjustIndex(start, end, length int) (int, int) {
	if end > length {
		end = length
	} else if end < 0 {
		end += length
		end = max(end, 0)
	}
	if start < 0 {
		start += length
		start = max(start, 0)
	}
	return start, end
}

func allForStringFunc(s string, condition func(rune) bool) bool {
	for _, r := range s {
		if !condition(r) {
			return false
		}
	}
	return true
}

func isX(s string, whenEmpty bool, condition func(rune) bool) bool {
	if len(s) == 0 {
		return whenEmpty
	}
	return allForStringFunc(s, condition)
}

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

// Capitalize Return a copy of the string with its first character capitalized and the rest lowercased.
func Capitalize(s string) string {
	length := Length(s)
	if length == 0 {
		return s
	} else if length == 1 {
		return strings.ToUpper(s)
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + strings.ToLower(s[size:])
}

// CaseFold Return a casefolded copy of the string. Casefolded strings may be used for caseless matching. Casefolding is similar to lowercasing but more aggressive because it is intended to remove all case distinctions in a string. For example, the German lowercase letter 'ß' is equivalent to "ss". Since it is already lowercase, lower() would do nothing to 'ß'; casefold() converts it to "ss".
func CaseFold(s string) string {
	c := cases.Fold()
	return c.String(s)
}

// Count Return the number of non-overlapping occurrences of substring sub in the range [start, end]. Optional arguments start and end are interpreted as in slice notation.
func Count(s, sub string, options ...Option) int {
	arg := Arg{
		Start: 0,
		End:   len(s),
	}
	for _, option := range options {
		option(&arg)
	}
	start, end := adjustIndex(arg.Start, arg.End, len(s))
	if (end - start) < 0 {
		return 0
	}
	return strings.Count(s[start:end], sub)
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
