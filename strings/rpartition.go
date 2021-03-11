package strings

// RPartition Split the string at the last occurrence of sep, and return a 3-tuple containing the part before the separator, the separator itself, and the part after the separator. If the separator is not found, return a 3-tuple containing two empty strings, followed by the string itself.
func RPartition(s, sep string) (string, string, string) {
	if v := LastSplitN(s, sep, 2); len(v) == 2 {
		return v[0], sep, v[1]
	}
	return "", "", s
}
