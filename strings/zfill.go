package strings

// ZFill Return a copy of the string left filled with ASCII '0' digits to make a string of length width. A leading sign prefix ('+'/'-') is handled by inserting the padding after the sign character rather than before. The original string is returned if width is less than or equal to len(s).
func ZFill(s string, width int) string {
	if Length(s) != 0 {
		if c := s[0]; c == '+' || c == '-' {
			return string(c) + RJust(s[1:], width-1, '0')
		}
	}
	return RJust(s, width, '0')
}
