1.  install "protoc" compiler and protobuf, grpc packages for golang
snap install protobuf --classic
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

2. build proto file
protoc -I=<proto_files_src_dir> --go_out=<path_for_generated_files> --go_opt=paths=source_relative --go-grpc_out=<path_for_generated_grpc_files> --go-grpc_opt=paths=source_relative <list_of_proto_files>

example: protoc --go_out=proto/generated --go_opt=paths=source_relative --go-grpc_out=proto/generated --go-grpc_opt=paths=source_relative proto/*.proto