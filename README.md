# video-streaming-sample

## Overview

## TecStack

- [gRPC](https://grpc.io/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Go](https://golang.org/): 1.23.2

## Preparation

### Install the required packages

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

cd api
go mod tidy
```

### Generate the gRPC code

```bash
protoc --proto_path=./proto \
  --go_out=. --go-grpc_out=. video.proto
```
