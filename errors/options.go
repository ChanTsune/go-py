package errors

type errorConfig struct {
	message string
}

// Option Error constractor option
type Option func(*errorConfig)

// Message parameter message
func Message(message string) Option {
	return func(a *errorConfig) {
		a.message = message
	}
}
