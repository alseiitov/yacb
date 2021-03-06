// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: rate.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RateServiceClient is the client API for RateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateServiceClient interface {
	GetCurrentRate(ctx context.Context, in *GetCurrentRateRequest, opts ...grpc.CallOption) (*Rate, error)
	GetRateChange(ctx context.Context, in *GetRateChangeRequest, opts ...grpc.CallOption) (*GetRateChangeResponse, error)
}

type rateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateServiceClient(cc grpc.ClientConnInterface) RateServiceClient {
	return &rateServiceClient{cc}
}

func (c *rateServiceClient) GetCurrentRate(ctx context.Context, in *GetCurrentRateRequest, opts ...grpc.CallOption) (*Rate, error) {
	out := new(Rate)
	err := c.cc.Invoke(ctx, "/proto.RateService/GetCurrentRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateServiceClient) GetRateChange(ctx context.Context, in *GetRateChangeRequest, opts ...grpc.CallOption) (*GetRateChangeResponse, error) {
	out := new(GetRateChangeResponse)
	err := c.cc.Invoke(ctx, "/proto.RateService/GetRateChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServiceServer is the server API for RateService service.
// All implementations must embed UnimplementedRateServiceServer
// for forward compatibility
type RateServiceServer interface {
	GetCurrentRate(context.Context, *GetCurrentRateRequest) (*Rate, error)
	GetRateChange(context.Context, *GetRateChangeRequest) (*GetRateChangeResponse, error)
	mustEmbedUnimplementedRateServiceServer()
}

// UnimplementedRateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRateServiceServer struct {
}

func (UnimplementedRateServiceServer) GetCurrentRate(context.Context, *GetCurrentRateRequest) (*Rate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentRate not implemented")
}
func (UnimplementedRateServiceServer) GetRateChange(context.Context, *GetRateChangeRequest) (*GetRateChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRateChange not implemented")
}
func (UnimplementedRateServiceServer) mustEmbedUnimplementedRateServiceServer() {}

// UnsafeRateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateServiceServer will
// result in compilation errors.
type UnsafeRateServiceServer interface {
	mustEmbedUnimplementedRateServiceServer()
}

func RegisterRateServiceServer(s grpc.ServiceRegistrar, srv RateServiceServer) {
	s.RegisterService(&RateService_ServiceDesc, srv)
}

func _RateService_GetCurrentRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetCurrentRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RateService/GetCurrentRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetCurrentRate(ctx, req.(*GetCurrentRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateService_GetRateChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServiceServer).GetRateChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RateService/GetRateChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServiceServer).GetRateChange(ctx, req.(*GetRateChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RateService_ServiceDesc is the grpc.ServiceDesc for RateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RateService",
	HandlerType: (*RateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentRate",
			Handler:    _RateService_GetCurrentRate_Handler,
		},
		{
			MethodName: "GetRateChange",
			Handler:    _RateService_GetRateChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
