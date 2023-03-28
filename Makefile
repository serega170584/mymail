-include .env

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/proto/mail.proto
