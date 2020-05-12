package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"server/gen/api"
	"server/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
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
