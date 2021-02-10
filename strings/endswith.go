package strings

import "strings"

// EndsWith Return True if the string ends with the specified suffix, otherwise return False. With optional start, test beginning at that position. With optional end, stop comparing at that position.
func EndsWith(s, suffix string, options ...Option) bool {
	length := len(s)
	args := Arg{
		Start: 0,
		End:   length,
	}
	for _, option := range options {
		option(&args)
	}
	start, end := adjustIndex(args.Start, args.End, length)
	if end-start < len(suffix) {
		return false
	}
	return strings.HasSuffix(s[start:end], suffix)
}

// EndsWithAny Return True if the string ends with the specified suffixes, otherwise return False. With optional start, test beginning at that position. With optional end, stop comparing at that position.
func EndsWithAny(s string, suffixes []string, options ...Option) bool {
	length := len(s)
	args := Arg{
		Start: 0,
		End:   length,
	}
	for _, option := range options {
		option(&args)
	}
	start, end := adjustIndex(args.Start, args.End, length)
	str := s[start:end]
	for _, suffix := range suffixes {
		if end-start >= len(suffix) && strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}
