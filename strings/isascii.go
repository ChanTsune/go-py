package strings

// IsASCII Return True if the string is empty or all characters in the string are ASCII, False otherwise. ASCII characters have code points in the range U+0000-U+007F.
func IsASCII(s string) bool {
	return isX(s, true, isASCII)
}

func isASCII(r rune) bool {
	i := int64(r)
	return 0x00 <= i && i <= 0x7f
}
