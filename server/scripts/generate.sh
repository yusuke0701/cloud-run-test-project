#!/usr/bin/env bash

cd `dirname $0`

cd ../../

protoc -Iproto \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:./server \
  proto/*.proto