// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package businesspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BusinessServiceClient is the client API for BusinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessServiceClient interface {
	ListBusiness(ctx context.Context, in *ListBusinessRequest, opts ...grpc.CallOption) (*ListBusinessResponse, error)
	ListBusinessById(ctx context.Context, in *ListBusinessByIdRequest, opts ...grpc.CallOption) (*ListBusinessByIdResponse, error)
	GetBusinessById(ctx context.Context, in *GetBusinessByIdRequest, opts ...grpc.CallOption) (*GetBusinessByIdResponse, error)
}

type businessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessServiceClient(cc grpc.ClientConnInterface) BusinessServiceClient {
	return &businessServiceClient{cc}
}

func (c *businessServiceClient) ListBusiness(ctx context.Context, in *ListBusinessRequest, opts ...grpc.CallOption) (*ListBusinessResponse, error) {
	out := new(ListBusinessResponse)
	err := c.cc.Invoke(ctx, "/seatlect.BusinessService/ListBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) ListBusinessById(ctx context.Context, in *ListBusinessByIdRequest, opts ...grpc.CallOption) (*ListBusinessByIdResponse, error) {
	out := new(ListBusinessByIdResponse)
	err := c.cc.Invoke(ctx, "/seatlect.BusinessService/ListBusinessById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) GetBusinessById(ctx context.Context, in *GetBusinessByIdRequest, opts ...grpc.CallOption) (*GetBusinessByIdResponse, error) {
	out := new(GetBusinessByIdResponse)
	err := c.cc.Invoke(ctx, "/seatlect.BusinessService/GetBusinessById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServiceServer is the server API for BusinessService service.
// All implementations must embed UnimplementedBusinessServiceServer
// for forward compatibility
type BusinessServiceServer interface {
	ListBusiness(context.Context, *ListBusinessRequest) (*ListBusinessResponse, error)
	ListBusinessById(context.Context, *ListBusinessByIdRequest) (*ListBusinessByIdResponse, error)
	GetBusinessById(context.Context, *GetBusinessByIdRequest) (*GetBusinessByIdResponse, error)
	mustEmbedUnimplementedBusinessServiceServer()
}

// UnimplementedBusinessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBusinessServiceServer struct {
}

func (UnimplementedBusinessServiceServer) ListBusiness(context.Context, *ListBusinessRequest) (*ListBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) ListBusinessById(context.Context, *ListBusinessByIdRequest) (*ListBusinessByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBusinessById not implemented")
}
func (UnimplementedBusinessServiceServer) GetBusinessById(context.Context, *GetBusinessByIdRequest) (*GetBusinessByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessById not implemented")
}
func (UnimplementedBusinessServiceServer) mustEmbedUnimplementedBusinessServiceServer() {}

// UnsafeBusinessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServiceServer will
// result in compilation errors.
type UnsafeBusinessServiceServer interface {
	mustEmbedUnimplementedBusinessServiceServer()
}

func RegisterBusinessServiceServer(s grpc.ServiceRegistrar, srv BusinessServiceServer) {
	s.RegisterService(&_BusinessService_serviceDesc, srv)
}

func _BusinessService_ListBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).ListBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seatlect.BusinessService/ListBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).ListBusiness(ctx, req.(*ListBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_ListBusinessById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusinessByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).ListBusinessById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seatlect.BusinessService/ListBusinessById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).ListBusinessById(ctx, req.(*ListBusinessByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_GetBusinessById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).GetBusinessById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seatlect.BusinessService/GetBusinessById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).GetBusinessById(ctx, req.(*GetBusinessByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BusinessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seatlect.BusinessService",
	HandlerType: (*BusinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBusiness",
			Handler:    _BusinessService_ListBusiness_Handler,
		},
		{
			MethodName: "ListBusinessById",
			Handler:    _BusinessService_ListBusinessById_Handler,
		},
		{
			MethodName: "GetBusinessById",
			Handler:    _BusinessService_GetBusinessById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "business.proto",
}
