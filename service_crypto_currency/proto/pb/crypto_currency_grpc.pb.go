// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: crypto_currency.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CryptoCurrencyServiceClient is the client API for CryptoCurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoCurrencyServiceClient interface {
	GetCryptoCurrencies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
}

type cryptoCurrencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoCurrencyServiceClient(cc grpc.ClientConnInterface) CryptoCurrencyServiceClient {
	return &cryptoCurrencyServiceClient{cc}
}

func (c *cryptoCurrencyServiceClient) GetCryptoCurrencies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoCurrencyService/GetCryptoCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoCurrencyServiceServer is the server API for CryptoCurrencyService service.
// All implementations must embed UnimplementedCryptoCurrencyServiceServer
// for forward compatibility
type CryptoCurrencyServiceServer interface {
	GetCryptoCurrencies(context.Context, *emptypb.Empty) (*GetCurrenciesResponse, error)
	mustEmbedUnimplementedCryptoCurrencyServiceServer()
}

// UnimplementedCryptoCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoCurrencyServiceServer struct {
}

func (UnimplementedCryptoCurrencyServiceServer) GetCryptoCurrencies(context.Context, *emptypb.Empty) (*GetCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCryptoCurrencies not implemented")
}
func (UnimplementedCryptoCurrencyServiceServer) mustEmbedUnimplementedCryptoCurrencyServiceServer() {}

// UnsafeCryptoCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoCurrencyServiceServer will
// result in compilation errors.
type UnsafeCryptoCurrencyServiceServer interface {
	mustEmbedUnimplementedCryptoCurrencyServiceServer()
}

func RegisterCryptoCurrencyServiceServer(s grpc.ServiceRegistrar, srv CryptoCurrencyServiceServer) {
	s.RegisterService(&CryptoCurrencyService_ServiceDesc, srv)
}

func _CryptoCurrencyService_GetCryptoCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoCurrencyServiceServer).GetCryptoCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoCurrencyService/GetCryptoCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoCurrencyServiceServer).GetCryptoCurrencies(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoCurrencyService_ServiceDesc is the grpc.ServiceDesc for CryptoCurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoCurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CryptoCurrencyService",
	HandlerType: (*CryptoCurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCryptoCurrencies",
			Handler:    _CryptoCurrencyService_GetCryptoCurrencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto_currency.proto",
}
