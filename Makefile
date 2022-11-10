.generate-proto:
	protoc --proto_path=testing/schema --go_out=generated --go_opt=paths=source_relative --go-grpc_out=generated --go-grpc_opt=paths=source_relative schema.proto