package errors

import (
	"errors"
	"log"

	"google.golang.org/grpc/codes"
)

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
		case *UnauthorizedError:
			code = codes.PermissionDenied
			msg = e.FrontMessage
			log.Printf("UnauthorizedError: %v", e)
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
