protoc model.proto -I. --go_out=model
protoc query.proto -I. -I${GOPATH}/src/github.com/googleapis/googleapis/ --go_out=plugins=grpc:model