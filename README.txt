1.  install "protoc" compiler and protobuf, grpc packages for golang
snap install protobuf
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

2. build proto file
protoc -I<path to proto files> --go.out=<path to output directory for generated code> --go-grpc_out=<path to output dir for generated grpc code>
--go_opt=module=<module name in proto file>  --go-grpc_opt=module=<path to proto file>

example: protoc --go_out=proto/generated --go_opt=paths=source_relative --go-grpc_out=proto/generated --go-grpc_opt=paths=source_relative proto/*.proto