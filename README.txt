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