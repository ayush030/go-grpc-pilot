BUILD_FLAGS ?= CGO_ENABLED=0

clean:
	rm -rf build/*

client:
	go build $(BUILD_FLAGS) -o build/client ./client/cmd

server:
	go build $(BUILD_FLAGS) -o build/server ./server/cmd

proto:
	protoc --go_out=proto/generated --go_opt=paths=source_relative --go-grpc_out=proto/generated --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: clean proto client server