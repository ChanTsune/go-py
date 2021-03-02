package strings

import "strings"

// Replace Return a copy of the string with all occurrences of substring old replaced by new.
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// ReplaceN Return a copy of the string with all occurrences of substring old replaced by new. The argument count is given, only the first count occurrences are replaced.
func ReplaceN(s, old, new string, count int) string {
	return strings.Replace(s, old, new, count)
}
