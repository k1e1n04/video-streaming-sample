package errors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// ErrorHandler is a function to handle errors
func ErrorHandler[T any](f func() (*T, error)) (*T, error) {
	res, err := f()
	if err != nil {
		msg, code := HandleError(err)
		return nil, status.Error(code, msg)
	}
	return res, nil
}

// ErrorHandlerWithoutResponse is a function to handle errors without response
func ErrorHandlerWithoutResponse(f func() error) error {
	err := f()
	if err != nil {
		msg, code := HandleError(err)
		return status.Error(code, msg)
	}
	return nil
}

// HandleError is a function to handle error
func HandleError(err error) (string, codes.Code) {
	code := codes.Internal
	msg := "internal server error"
	var e interface{}

	if errors.As(err, &e) {
		switch e := e.(type) {
		case *BadRequestError:
			code = codes.InvalidArgument
			msg = e.FrontMessage
			log.Printf("BadRequestError: %v", e)
		case *NotFoundError:
			code = codes.NotFound
			msg = e.FrontMessage
			log.Printf("NotFoundError: %v", e)
		case *InvalidStatementError:
			log.Printf("InvalidStatementError: %v", e)
			if e.Cause != nil {
				log.Printf("Cause: %v", e.Cause)
			}
		default:
			log.Printf("Unknown error: %v", e)
		}
	} else {
		log.Printf("Unknown error: %v", err)
	}

	return msg, code
}
