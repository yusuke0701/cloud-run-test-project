package main

import (
	"context"
	"crypto/x509"
	"flag"
	"fmt"
	"gateway/gen/api"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	sal "github.com/salrashid123/oauth2/google"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	httpPort = "8081"
)

const (
	address  = "server-5dnafyaz7q-ue.a.run.app"
	audience = "https://server-5dnafyaz7q-ue.a.run.app"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := runGW(); err != nil {
		glog.Fatal(err)
	}
}

func runGW() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pool, _ := x509.SystemCertPool()
	creds := credentials.NewClientTLSFromCert(pool, "")
	perRPC, err := getRPCCreds(ctx)
	if err != nil {
		return err
	}

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithBlock(),
	}

	grpcServerEndpoint := flag.String("grpc-server-endpoint", fmt.Sprintf("%s:%s", address, "443"), "gRPC server endpoint")

	if err := api.RegisterPancakeBakerServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = httpPort
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func getRPCCreds(ctx context.Context) (credentials.PerRPCCredentials, error) {
	scopes := "https://www.googleapis.com/auth/userinfo.email"

	creds, err := google.FindDefaultCredentials(ctx, scopes)
	if err != nil {
		return nil, err
	}

	idTokenSource, err := sal.IdTokenSource(
		&sal.IdTokenConfig{
			Credentials: creds,
			Audiences:   []string{audience},
		},
	)
	if err != nil {
		return nil, err
	}

	rpcCreds, err := sal.NewIDTokenRPCCredential(ctx, idTokenSource)
	if err != nil {
		return nil, err
	}

	return rpcCreds, nil
}
