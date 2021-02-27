package strings

import "strings"

// Partition Return string.
func Partition(s, sep string) (string, string, string) {
	if v := strings.SplitN(s, sep, 2); len(v) == 2 {
		return v[0], sep, v[1]
	}
	return s, "", ""
}
