package interceptor

import (
	"context"
	"fmt"
	"strings"

	context2 "github.com/k1e1n04/video-streaming-sample/api/adapter/context"
	"github.com/k1e1n04/video-streaming-sample/api/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// allowedMethodsWithoutAuth is a map of gRPC methods that are allowed without authentication
var allowedMethodsWithoutAuth = []string{
	"/user.AuthService/Login",
}

// AuthInterceptor is a gRPC interceptor that authenticates the user
type AuthInterceptor struct {
	verifyToken func(string) (string, error)
}

// NewAuthInterceptor is a constructor
func NewAuthInterceptor(verifyToken func(string) (string, error)) *AuthInterceptor {
	return &AuthInterceptor{verifyToken: verifyToken}
}

// Unary  is a gRPC interceptor that authenticates the user
func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod

		// check if the method is allowed without authentication
		for _, allowedMethod := range allowedMethodsWithoutAuth {
			if method == allowedMethod {
				return handler(ctx, req)
			}
		}

		// call the handler
		return handler(ctx, req)
	}
}

// applyAuth is a function that applies authentication to the gRPC method
func (a *AuthInterceptor) applyAuth(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.NewUnauthorizedError(
			"failed to parse metadata",
			"unauthorized",
			nil,
		)
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, errors.NewUnauthorizedError(
			"authorization header not found",
			"unauthorized",
			nil,
		)
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")
	if token == "" {
		return nil, errors.NewUnauthorizedError(
			"token must not be empty",
			"unauthorized",
			nil,
		)
	}

	userID, err := a.verifyToken(token)
	if err != nil {
		return nil, errors.NewUnauthorizedError(
			fmt.Sprintf("failed to verify token: %v", err),
			"unauthorized",
			err,
		)
	}

	ctx = context2.SetUserID(ctx, userID)

	return ctx, nil
}
