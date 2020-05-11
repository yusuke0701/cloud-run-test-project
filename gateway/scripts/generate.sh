#!/usr/bin/env bash

cd `dirname $0`

cd ../../

protoc -Iproto \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:./gateway \
  proto/*.proto