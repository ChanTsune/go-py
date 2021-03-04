package strings

import "go-py/errors"

// RIndex Return the highest index in the string where substring sub is found, such that sub is contained within s[start:end]. Optional arguments start and end are interpreted as in slice notation. Return error if sub is not found.
func RIndex(s, sub string, options ...Option) (int, error) {
	i := RFind(s, sub, options...)
	return i, func(i int) error {
		if i == -1 {
			return errors.NewValueError(errors.Message("substring not found"))
		}
		return nil
	}(i)
}
