package errors

// NotFoundError is an error for not found request
type NotFoundError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewNotFoundError is a constructor
func NewNotFoundError(
	debugMessage string,
	frontMessage string,
	cause error,
) *NotFoundError {
	return &NotFoundError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error returns an error message
func (e *NotFoundError) Error() string {
	return e.DebugMessage
}
