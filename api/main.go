package main

import (
	"github.com/k1e1n04/video-streaming-sample/api/adapter/controllers"
	"github.com/k1e1n04/video-streaming-sample/api/adapter/grpc/video"
	"github.com/k1e1n04/video-streaming-sample/api/di"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	container := di.Init()

	var vc controllers.VideoController
	err = container.Invoke(func(c controllers.VideoController) {
		vc = c
	})
	if err != nil {
		panic(err)
	}
	video.RegisterVideoServiceServer(grpcServer, &vc)

	log.Println("gRPC server running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
