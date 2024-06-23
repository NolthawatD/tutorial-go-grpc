# tutorial-go-grpc

## Ref : https://github.com/codebangkok/golang
> Install grpc begin Go gRPC for macOS


## Build protoc to server go

> protoc calculator.proto --go_out=../server  
get request and response only

> --go-grpc_out=../server
file server


## new file in folder proto
paste command for can use excute file
> chmod +x gen.sh


## install package grpc in server
go get google.golang.org/grpc

## install Evan for gRPC like Postman
> brew tap ktr0731/evans
> brew install evans

## Open terminal in area
> cmd + shift + p
> >terminal: create new terminal in editor area


## Ref to file Proto
> evans --proto=./calculator.proto (-p 50051 if change port)
help
show package
show service
desc HelloRequest
desc HelloResponse
cal Hello (method rpc)
exit

> evans -r