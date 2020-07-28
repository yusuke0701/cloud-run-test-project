#!/usr/bin/env bash

cd $(dirname $0)

cd ../

protoc -Iproto \
    -I$GOPATH/src \
    -Igrpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:./server \
    proto/*.proto

protoc -Iproto \
    -I$GOPATH/src \
    -Igrpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:./gateway \
    proto/*.proto
