.PHONY: gen_proto gen_proto_common gen_proto_user clean_proto

gen_proto: gen_proto_common gen_proto_user

gen_proto_common:
	@protoc --go_out=. -I=. api/protobuf/common.proto

gen_proto_user:
	@protoc --go_out=. --go-grpc_out=. -I=. api/protobuf/user.proto

clean_proto:
	@rm -rf internal/genproto/