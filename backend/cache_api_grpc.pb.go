// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// CacherClient is the client API for Cacher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacherClient interface {
	Set(ctx context.Context, in *SetBody, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetBody, opts ...grpc.CallOption) (*GetResponse, error)
	Clear(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type cacherClient struct {
	cc grpc.ClientConnInterface
}

func NewCacherClient(cc grpc.ClientConnInterface) CacherClient {
	return &cacherClient{cc}
}

func (c *cacherClient) Set(ctx context.Context, in *SetBody, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Cacher/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) Get(ctx context.Context, in *GetBody, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.Cacher/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) Clear(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Cacher/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacherServer is the server API for Cacher service.
// All implementations must embed UnimplementedCacherServer
// for forward compatibility
type CacherServer interface {
	Set(context.Context, *SetBody) (*Response, error)
	Get(context.Context, *GetBody) (*GetResponse, error)
	Clear(context.Context, *Empty) (*Response, error)
	mustEmbedUnimplementedCacherServer()
}

// UnimplementedCacherServer must be embedded to have forward compatible implementations.
type UnimplementedCacherServer struct {
}

func (UnimplementedCacherServer) Set(context.Context, *SetBody) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacherServer) Get(context.Context, *GetBody) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacherServer) Clear(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedCacherServer) mustEmbedUnimplementedCacherServer() {}

// UnsafeCacherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacherServer will
// result in compilation errors.
type UnsafeCacherServer interface {
	mustEmbedUnimplementedCacherServer()
}

func RegisterCacherServer(s grpc.ServiceRegistrar, srv CacherServer) {
	s.RegisterService(&Cacher_ServiceDesc, srv)
}

func _Cacher_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cacher/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).Set(ctx, req.(*SetBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cacher/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).Get(ctx, req.(*GetBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cacher/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).Clear(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Cacher_ServiceDesc is the grpc.ServiceDesc for Cacher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cacher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cacher",
	HandlerType: (*CacherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Cacher_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cacher_Get_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _Cacher_Clear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache_api.proto",
}
