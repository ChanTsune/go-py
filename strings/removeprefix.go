package strings

import "strings"

// RemoevPrefix If the string starts with the prefix string, return string[len(prefix):]. Otherwise, return a copy of the original string.
func RemoevPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}
