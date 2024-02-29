// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ibc/applications/perm/v1/query.proto

package permv1

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
	Query_PermissionedRelayer_FullMethodName  = "/ibc.applications.perm.v1.Query/PermissionedRelayer"
	Query_PermissionedRelayers_FullMethodName = "/ibc.applications.perm.v1.Query/PermissionedRelayers"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// PermissionedRelayer queries a permissioned ibc relayer for the specific channel.
	PermissionedRelayer(ctx context.Context, in *QueryPermissionedRelayerRequest, opts ...grpc.CallOption) (*QueryPermissionedRelayerResponse, error)
	// PermissionedRelayers queries a permissioned ibc relayers.
	PermissionedRelayers(ctx context.Context, in *QueryPermissionedRelayersRequest, opts ...grpc.CallOption) (*QueryPermissionedRelayersResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) PermissionedRelayer(ctx context.Context, in *QueryPermissionedRelayerRequest, opts ...grpc.CallOption) (*QueryPermissionedRelayerResponse, error) {
	out := new(QueryPermissionedRelayerResponse)
	err := c.cc.Invoke(ctx, Query_PermissionedRelayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PermissionedRelayers(ctx context.Context, in *QueryPermissionedRelayersRequest, opts ...grpc.CallOption) (*QueryPermissionedRelayersResponse, error) {
	out := new(QueryPermissionedRelayersResponse)
	err := c.cc.Invoke(ctx, Query_PermissionedRelayers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// PermissionedRelayer queries a permissioned ibc relayer for the specific channel.
	PermissionedRelayer(context.Context, *QueryPermissionedRelayerRequest) (*QueryPermissionedRelayerResponse, error)
	// PermissionedRelayers queries a permissioned ibc relayers.
	PermissionedRelayers(context.Context, *QueryPermissionedRelayersRequest) (*QueryPermissionedRelayersResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) PermissionedRelayer(context.Context, *QueryPermissionedRelayerRequest) (*QueryPermissionedRelayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionedRelayer not implemented")
}
func (UnimplementedQueryServer) PermissionedRelayers(context.Context, *QueryPermissionedRelayersRequest) (*QueryPermissionedRelayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionedRelayers not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_PermissionedRelayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPermissionedRelayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PermissionedRelayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PermissionedRelayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PermissionedRelayer(ctx, req.(*QueryPermissionedRelayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PermissionedRelayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPermissionedRelayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PermissionedRelayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PermissionedRelayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PermissionedRelayers(ctx, req.(*QueryPermissionedRelayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.perm.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PermissionedRelayer",
			Handler:    _Query_PermissionedRelayer_Handler,
		},
		{
			MethodName: "PermissionedRelayers",
			Handler:    _Query_PermissionedRelayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/perm/v1/query.proto",
}