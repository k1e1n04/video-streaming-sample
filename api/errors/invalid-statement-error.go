package errors

// InvalidStatementError is an error for invalid statement
type InvalidStatementError struct {
	DebugMessage string
	Cause        error
}

// NewInvalidStatementError is a constructor
func NewInvalidStatementError(
	debugMessage string,
	cause error,
) *InvalidStatementError {
	return &InvalidStatementError{
		DebugMessage: debugMessage,
		Cause:        cause,
	}
}

// Error returns an error message
func (e *InvalidStatementError) Error() string {
	return e.DebugMessage
}
