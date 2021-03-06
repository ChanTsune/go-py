package strings

import "strings"

// RemoveSuffix If the string ends with the suffix string and that suffix is not empty, return string[:-len(suffix)]. Otherwise, return a copy of the original string
func RemoveSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
