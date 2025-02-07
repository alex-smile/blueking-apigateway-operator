// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package serverpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResourcesClient is the client API for Resources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourcesClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Healthz(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthzResponse, error)
	GetLeader(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetLeaderResponse, error)
}

type resourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewResourcesClient(cc grpc.ClientConnInterface) ResourcesClient {
	return &resourcesClient{cc}
}

func (c *resourcesClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/serverpb.Resources/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/serverpb.Resources/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/serverpb.Resources/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Healthz(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/serverpb.Resources/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) GetLeader(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetLeaderResponse, error) {
	out := new(GetLeaderResponse)
	err := c.cc.Invoke(ctx, "/serverpb.Resources/GetLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcesServer is the server API for Resources service.
// All implementations must embed UnimplementedResourcesServer
// for forward compatibility
type ResourcesServer interface {
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Healthz(context.Context, *empty.Empty) (*HealthzResponse, error)
	GetLeader(context.Context, *empty.Empty) (*GetLeaderResponse, error)
	mustEmbedUnimplementedResourcesServer()
}

// UnimplementedResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedResourcesServer struct {
}

func (UnimplementedResourcesServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedResourcesServer) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedResourcesServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResourcesServer) Healthz(context.Context, *empty.Empty) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedResourcesServer) GetLeader(context.Context, *empty.Empty) (*GetLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeader not implemented")
}
func (UnimplementedResourcesServer) mustEmbedUnimplementedResourcesServer() {}

// UnsafeResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourcesServer will
// result in compilation errors.
type UnsafeResourcesServer interface {
	mustEmbedUnimplementedResourcesServer()
}

func RegisterResourcesServer(s grpc.ServiceRegistrar, srv ResourcesServer) {
	s.RegisterService(&Resources_ServiceDesc, srv)
}

func _Resources_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Resources/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Resources/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Resources/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Resources/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Healthz(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_GetLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Resources/GetLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetLeader(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Resources_ServiceDesc is the grpc.ServiceDesc for Resources service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resources_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Resources",
	HandlerType: (*ResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _Resources_Sync_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _Resources_Diff_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Resources_List_Handler,
		},
		{
			MethodName: "Healthz",
			Handler:    _Resources_Healthz_Handler,
		},
		{
			MethodName: "GetLeader",
			Handler:    _Resources_GetLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/serverpb/rpc.proto",
}
