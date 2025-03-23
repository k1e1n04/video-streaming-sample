package errors

// BadRequestError is an error for bad request
type BadRequestError struct {
	DebugMessage string
	FrontMessage string
}

// NewBadRequestError is a constructor
func NewBadRequestError(
	debugMessage string,
	frontMessage string,
) *BadRequestError {
	return &BadRequestError{
		DebugMessage: debugMessage,
		FrontMessage: frontMessage,
	}
}

// Error returns an error message
func (e *BadRequestError) Error() string {
	return e.FrontMessage
}
