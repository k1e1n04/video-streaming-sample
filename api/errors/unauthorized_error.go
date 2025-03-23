package errors

// UnauthorizedError is an error for unauthorized request
type UnauthorizedError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewUnauthorizedError is a constructor
func NewUnauthorizedError(
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
func (e *UnauthorizedError) Error() string {
	return e.FrontMessage
}
