package proto

//go:generate protoc --proto_path=../../../service-protos  --go_out=plugins=grpc,paths=source_relative:. common.proto
