.PHONY: gen_proto gen_proto_common gen_proto_user gen_proto_token clean_proto gen_client gen_client_common

gen_proto: gen_proto_common gen_proto_user gen_proto_token

gen_proto_common:
	@protoc --go_out=. -I=api/protobuf common.proto

gen_proto_user:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf user.proto

gen_proto_token:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf token.proto

clean_proto:
	@rm -rf internal/genproto/

# This requires the seatlect_mobile repo to be adjacent to this project repository
gen_client: gen_client_common gen_client_user gen_client_token

gen_client_common:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf common.proto

gen_client_user:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf user.proto

gen_client_token:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf token.proto
