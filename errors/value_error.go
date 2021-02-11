package errors

// ValueError value error.
type ValueError struct {
	message string
}

func (e *ValueError) Error() string {
	return "ValueError: " + e.message
}

// NewValueError new ValueError
func NewValueError(options ...Option) *ValueError {
	arg := errorConfig{}
	for _, option := range options {
		option(&arg)
	}
	e := new(ValueError)
	e.message = arg.message
	return e
}
