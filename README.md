This repository is a good starting point if one wants to learn about gRPC and how to develop APIs using it.
Pre-requisites for this is some knowledge in Golang, Protobuf and APIs.

![1_P0z3ortvUH4gZtIyj4al_A](https://github.com/user-attachments/assets/10c53999-dddd-4014-9fc5-8833faf6d911)


Steps to run the project:

1.  install "protoc" compiler and protobuf, grpc packages for golang
sudo snap install protobuf --classic OR
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

2. prepare go vendors by running
    make vendor

3. build proto files present in proto directories by running
    make proto

this will generated go code for grpc and protobuf. It runs below command: 
protoc -I=<proto_files_src_dir> --go_out=<path_for_generated_files> --go_opt=paths=source_relative --go-grpc_out=<path_for_generated_grpc_files> --go-grpc_opt=paths=source_relative <list_of_proto_files>


4. build client and server builds
    make client server
this should create two builds(server and client) in build directory
