// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reservationpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	ListReservation(ctx context.Context, in *ListReservationRequest, opts ...grpc.CallOption) (*ListReservationResponse, error)
	ReserveSeats(ctx context.Context, in *ReserveSeatsRequest, opts ...grpc.CallOption) (*ReserveSeatsResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) ListReservation(ctx context.Context, in *ListReservationRequest, opts ...grpc.CallOption) (*ListReservationResponse, error) {
	out := new(ListReservationResponse)
	err := c.cc.Invoke(ctx, "/seatlect.ReservationService/ListReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ReserveSeats(ctx context.Context, in *ReserveSeatsRequest, opts ...grpc.CallOption) (*ReserveSeatsResponse, error) {
	out := new(ReserveSeatsResponse)
	err := c.cc.Invoke(ctx, "/seatlect.ReservationService/ReserveSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	ListReservation(context.Context, *ListReservationRequest) (*ListReservationResponse, error)
	ReserveSeats(context.Context, *ReserveSeatsRequest) (*ReserveSeatsResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) ListReservation(context.Context, *ListReservationRequest) (*ListReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReservation not implemented")
}
func (UnimplementedReservationServiceServer) ReserveSeats(context.Context, *ReserveSeatsRequest) (*ReserveSeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveSeats not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&_ReservationService_serviceDesc, srv)
}

func _ReservationService_ListReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ListReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seatlect.ReservationService/ListReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ListReservation(ctx, req.(*ListReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ReserveSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveSeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ReserveSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seatlect.ReservationService/ReserveSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ReserveSeats(ctx, req.(*ReserveSeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReservationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seatlect.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListReservation",
			Handler:    _ReservationService_ListReservation_Handler,
		},
		{
			MethodName: "ReserveSeats",
			Handler:    _ReservationService_ReserveSeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation.proto",
}
