.PHONY: gen_proto gen_proto_common gen_proto_entity clean_proto

gen_proto: gen_proto_common gen_proto_auth

gen_proto_common:
	@protoc --go_out=. -I=. api/protobuf/common.proto

gen_proto_auth:
	@protoc --go_out=. --go-grpc_out=. -I=. api/protobuf/auth.proto

clean_proto:
	@rm -rf internal/genproto/