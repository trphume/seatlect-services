.PHONY: gen_proto gen_proto_entity gen_proto_entity

gen_proto: gen_proto_entity gen_proto_auth

gen_proto_entity:
	@protoc --go_out=. -I=. api/protobuf/entity.proto

gen_proto_auth:
	@protoc --go_out=. --go-grpc_out=. -I=. api/protobuf/auth.proto