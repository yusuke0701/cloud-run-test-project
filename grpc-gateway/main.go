package main

import (
	"grpc-gateway/gen/api"
	"grpc-gateway/handler"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	portOfServer = "5000"
	endpoint     = "localhost:" + portOfServer
)

func main() {
	go runServer()
	go runGateway()
}

func runServer() {
	port := portOfServer

	lis, err := net.Listen("tcp", ":"+port)
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

func runGateway() error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterPancakeBakerServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("start gRPC gateway port: %v", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("failed to serve gateway: %v", err)
	}
}
