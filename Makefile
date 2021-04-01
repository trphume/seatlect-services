.PHONY: gen_proto clean_proto gen_proto_common gen_proto_user gen_proto_token gen_proto_business gen_proto_order gen_proto_reservation
.PHONY: gen_client clean_openapi gen_client_common gen_client_business gen_client_order gen_client_reservation
.PHONY: gen_openapi gen_openapi_user  gen_openapi_business gen_openapi_admin gen_openapi_request gen_openapi_reservation gen_openapi_placement clean_openapi

# This section contains commands for working with proto files
gen_proto: gen_proto_common gen_proto_user gen_proto_token gen_proto_business gen_proto_order gen_proto_reservation

gen_proto_common:
	@protoc --go_out=. -I=api/protobuf --go_opt=module=github.com/tphume/seatlect-services common.proto

gen_proto_user:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf --go_opt=module=github.com/tphume/seatlect-services --go-grpc_opt=module=github.com/tphume/seatlect-services user.proto

gen_proto_token:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf --go_opt=module=github.com/tphume/seatlect-services --go-grpc_opt=module=github.com/tphume/seatlect-services token.proto

gen_proto_business:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf --go_opt=module=github.com/tphume/seatlect-services --go-grpc_opt=module=github.com/tphume/seatlect-services business.proto

gen_proto_order:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf --go_opt=module=github.com/tphume/seatlect-services --go-grpc_opt=module=github.com/tphume/seatlect-services order.proto

gen_proto_reservation:
	@protoc --go_out=. --go-grpc_out=. -I=api/protobuf --go_opt=module=github.com/tphume/seatlect-services --go-grpc_opt=module=github.com/tphume/seatlect-services reservation.proto

clean_proto:
	@rm -rf internal/genproto/

# This requires the seatlect_mobile repo to be adjacent to this project repository
gen_client: gen_client_common gen_client_user gen_client_token gen_client_business gen_client_order gen_client_reservation

gen_client_common:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf common.proto

gen_client_user:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf user.proto

gen_client_token:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf token.proto

gen_client_business:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf business.proto

gen_client_order:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf order.proto

gen_client_reservation:
	@protoc --dart_out=grpc:../seatlect_mobile/packages/genproto/lib/src -I=api/protobuf reservation.proto

# This section contains commands for working with openapi files
gen_openapi: gen_openapi_user gen_openapi_business gen_openapi_admin gen_openapi_request gen_openapi_reservation gen_openapi_placement

gen_openapi_user:
	@oapi-codegen -o internal/gen_openapi/user_api/user_api.gen.go -package user_api -generate "types,server,spec" api/openapi/user.yml

gen_openapi_business:
	@oapi-codegen -o internal/gen_openapi/business_api/business_api.gen.go -package business_api -generate "types,server,spec" api/openapi/business.yml

gen_openapi_admin:
	@oapi-codegen -o internal/gen_openapi/admin_api/admin_api.gen.go -package admin_api -generate "types,server,spec" api/openapi/admin.yml

gen_openapi_request:
	@oapi-codegen -o internal/gen_openapi/request_api/request_api.gen.go -package request_api -generate "types,server,spec" api/openapi/request.yml

gen_openapi_reservation:
	@oapi-codegen -o internal/gen_openapi/reservation_api/reservation_api.gen.go -package reservation_api -generate "types,server,spec" api/openapi/reservation.yml

gen_openapi_placement:
	@oapi-codegen -o internal/gen_openapi/placement_api/placement_api.gen.go -package placement_api -generate "types,server,spec" api/openapi/placement.yml

clean_openapi:
	@rm -rf internal/gen_openapi