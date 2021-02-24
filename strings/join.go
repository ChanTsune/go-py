package strings

import "strings"

// Join Return a string which is the concatenation of the strings in iterable.
func Join(s string, iterable []string) string {
	return strings.Join(iterable, s)
}
