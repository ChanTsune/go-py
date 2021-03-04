package strings

import "strings"

// RFind Return the highest index in the string where substring sub is found, such that sub is contained within s[start:end]. Optional arguments start and end are interpreted as in slice notation. Return -1 on failure.
func RFind(s, sub string, options ...Option) int {
	length := len(s)
	args := Arg{
		Start: 0,
		End:   length,
	}
	for _, option := range options {
		option(&args)
	}
	start, end := adjustIndex(args.Start, args.End, length)
	if end-start < len(sub) {
		return -1
	}
	i := strings.LastIndex(s[start:end], sub)
	if i == -1 {
		return -1
	}
	return start + i
}
