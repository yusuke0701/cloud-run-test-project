#!/usr/bin/env bash

cd `dirname $0`

cd ../../

# TODO: googleapisのバージョン指定はダサい

protoc -Iproto \
  -I$GOPATH/src \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis \
  --go_out=plugins=grpc:./server \
  proto/*.proto