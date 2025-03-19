package errors

// BadRequestError is an error for bad request
type BadRequestError struct {
	DebugMessage string
	FrontMessage string
	Cause        error
}

// NewBadRequestError is a constructor
func NewBadRequestError(
	debugMessage string,
	frontMessage string,
	cause error,
) *BadRequestError {
	return &BadRequestError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
		Cause:        cause,
	}
}

// Error returns an error message
func (e *BadRequestError) Error() string {
	return e.DebugMessage
}
