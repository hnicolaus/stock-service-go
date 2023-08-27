// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: bibit.proto

package proto

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

const (
	Bibit_GetStockSummary_FullMethodName = "/proto.Bibit/GetStockSummary"
)

// BibitClient is the client API for Bibit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BibitClient interface {
	GetStockSummary(ctx context.Context, in *GetStockSummaryRequest, opts ...grpc.CallOption) (*GetStockSummaryResponse, error)
}

type bibitClient struct {
	cc grpc.ClientConnInterface
}

func NewBibitClient(cc grpc.ClientConnInterface) BibitClient {
	return &bibitClient{cc}
}

func (c *bibitClient) GetStockSummary(ctx context.Context, in *GetStockSummaryRequest, opts ...grpc.CallOption) (*GetStockSummaryResponse, error) {
	out := new(GetStockSummaryResponse)
	err := c.cc.Invoke(ctx, Bibit_GetStockSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BibitServer is the server API for Bibit service.
// All implementations must embed UnimplementedBibitServer
// for forward compatibility
type BibitServer interface {
	GetStockSummary(context.Context, *GetStockSummaryRequest) (*GetStockSummaryResponse, error)
	mustEmbedUnimplementedBibitServer()
}

// UnimplementedBibitServer must be embedded to have forward compatible implementations.
type UnimplementedBibitServer struct {
}

func (UnimplementedBibitServer) GetStockSummary(context.Context, *GetStockSummaryRequest) (*GetStockSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockSummary not implemented")
}
func (UnimplementedBibitServer) mustEmbedUnimplementedBibitServer() {}

// UnsafeBibitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BibitServer will
// result in compilation errors.
type UnsafeBibitServer interface {
	mustEmbedUnimplementedBibitServer()
}

func RegisterBibitServer(s grpc.ServiceRegistrar, srv BibitServer) {
	s.RegisterService(&Bibit_ServiceDesc, srv)
}

func _Bibit_GetStockSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BibitServer).GetStockSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bibit_GetStockSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BibitServer).GetStockSummary(ctx, req.(*GetStockSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bibit_ServiceDesc is the grpc.ServiceDesc for Bibit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bibit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Bibit",
	HandlerType: (*BibitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStockSummary",
			Handler:    _Bibit_GetStockSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bibit.proto",
}
