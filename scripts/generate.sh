#!/usr/bin/env bash

cd $(dirname $0)

cd ../

protoc -Iproto \
    -I$GOPATH/src \
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
    --go_out=plugins=grpc:./server \
    proto/*.proto

protoc -Iproto \
    -I$GOPATH/src \
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:./gateway \
    proto/*.proto
