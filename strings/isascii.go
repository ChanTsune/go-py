package strings

// IsASCII Return True if the string is empty or all characters in the string are ASCII, False otherwise. ASCII characters have code points in the range U+0000-U+007F.
func IsASCII(s string) bool {
	return isX(s, true, isASCII)
}

func isASCII(r rune) bool {
	return 0x00 <= r && r <= 0x7f
}
