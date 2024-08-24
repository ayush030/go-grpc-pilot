BUILD_FLAGS ?= CGO_ENABLED=0
PROTO_FILES ?= proto

clean:
	rm -rf build/*
	rm -rf proto/generated/*

vendor:
	go mod tidy -x
	go mod vendor

client:
	$(BUILD_FLAGS) go build -o build/client ./client/cmd

server:
	$(BUILD_FLAGS) go build -o build/server ./server/cmd

proto:
	protoc -I=$(PROTO_FILES) --go_out=generated --go_opt=paths=source_relative --go-grpc_out=generated --go-grpc_opt=paths=source_relative proto/*.proto


.PHONY: clean vendor proto client server
