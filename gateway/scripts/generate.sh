#!/usr/bin/env bash

cd `dirname $0`

cd ../../

# TODO: googleapisのバージョン指定はダサい & server/scripts/generate.sh を使いたい

protoc -Iproto \
  -I$GOPATH/src \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis \
  --go_out=plugins=grpc:./gateway \
  proto/*.proto

protoc -Iproto \
  -I$GOPATH/src \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:./gateway \
  proto/*.proto