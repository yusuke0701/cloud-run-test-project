package main

import (
	"flag"
	"fmt"
	"gateway/gen/api"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const endpoint = "https://server-5dnafyaz7q-ue.a.run.app"

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterPancakeBakerServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "50050"
	}
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
