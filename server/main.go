package main

import (
	"fmt"
	"log"
	"net"
	"server/gen/api"
	"server/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterPancakeBakerServiceServer(
		server,
		handler.NewBakerHandler(),
	)
	reflection.Register(server)

	log.Printf("start gRPC server port: %v", port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
