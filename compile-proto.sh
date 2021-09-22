# Using protoc directly:
# protoc model.proto -I. --go_out=model
# protoc query.proto -I. -I${GOPATH}/src/github.com/googleapis/googleapis/ --go_out=plugins=grpc:model

# Using all protoc binaries and dependencies from a docker image:
docker run --rm -u $(id -u) -v${PWD}:${PWD} -w${PWD} jaegertracing/protobuf:latest \
  --proto_path=${PWD} --go_out=${PWD}/model ${PWD}/model.proto
docker run --rm -u $(id -u) -v${PWD}:${PWD} -w${PWD} jaegertracing/protobuf:latest \
  --proto_path=${PWD} --go_out=plugins=grpc:${PWD}/model ${PWD}/query.proto
