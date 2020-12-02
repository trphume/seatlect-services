.PHONY: gen_proto gen_proto_common gen_proto_user gen_proto_token gen_client gen_client_common clean_proto
.PHONY: gen_openapi gen_openapi_user  gen_openapi_business clean_openapi

# This section contains commands for working with proto files
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


# This section contains commands for working with openapi files
gen_openapi: gen_openapi_user gen_openapi_business

gen_openapi_user:
	@oapi-codegen -o internal/gen_openapi/user_api/user_api.gen.go -package user_api api/openapi/user.yml

gen_openapi_business:
	@oapi-codegen -o internal/gen_openapi/business_api/business_api.gen.go -package business_api api/openapi/business.yml

clean_openapi:
	@rm -rf internal/gen_openapi