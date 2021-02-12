package strings

import "strings"

// Find Return the lowest index in the string where substring sub is found within the slice s[start:end]. Optional arguments start and end are interpreted as in slice notation. Return -1 if sub is not found.
func Find(s, sub string, options ...Option) int {
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
	i := strings.Index(s[start:end], sub)
	if i == -1 {
		return -1
	}
	return start + i
}
