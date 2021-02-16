package strings

import "go-py/errors"

// Index Return the lowest index in the string where substring sub is found within the slice s[start:end]. Optional arguments start and end are interpreted as in slice notation. Return error if sub is not found.
func Index(s, sub string, options ...Option) (int, error) {
	i := Find(s, sub, options...)
	return i, func(i int) error {
		if i == -1 {
			return errors.NewValueError(errors.Message("substring not found"))
		}
		return nil
	}(i)
}
