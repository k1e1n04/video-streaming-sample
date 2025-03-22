package intercepter

import (
	"context"
	"log"

	"google.golang.org/grpc/status"

	"github.com/k1e1n04/video-streaming-sample/api/errors"
	"google.golang.org/grpc"
)

// UnaryErrorInterceptor is a gRPC interceptor that converts errors to gRPC errors
func UnaryErrorInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	resp, err := handler(ctx, req)

	if err != nil {
		log.Printf("Error in %s: %v", info.FullMethod, err)
		msg, code := errors.HandleError(err)
		return nil, status.Errorf(code, msg)
	}

	return resp, err
}
