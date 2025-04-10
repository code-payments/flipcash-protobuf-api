// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: push/v1/push_service.proto

package pushpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Push_AddToken_FullMethodName     = "/flipcash.push.v1.Push/AddToken"
	Push_DeleteTokens_FullMethodName = "/flipcash.push.v1.Push/DeleteTokens"
)

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushClient interface {
	// AddToken adds a push token associated with a user.
	AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*AddTokenResponse, error)
	// DeleteTokens removes all push tokens within an app install for a user
	DeleteTokens(ctx context.Context, in *DeleteTokensRequest, opts ...grpc.CallOption) (*DeleteTokensResponse, error)
}

type pushClient struct {
	cc grpc.ClientConnInterface
}

func NewPushClient(cc grpc.ClientConnInterface) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*AddTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTokenResponse)
	err := c.cc.Invoke(ctx, Push_AddToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) DeleteTokens(ctx context.Context, in *DeleteTokensRequest, opts ...grpc.CallOption) (*DeleteTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTokensResponse)
	err := c.cc.Invoke(ctx, Push_DeleteTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServer is the server API for Push service.
// All implementations must embed UnimplementedPushServer
// for forward compatibility.
type PushServer interface {
	// AddToken adds a push token associated with a user.
	AddToken(context.Context, *AddTokenRequest) (*AddTokenResponse, error)
	// DeleteTokens removes all push tokens within an app install for a user
	DeleteTokens(context.Context, *DeleteTokensRequest) (*DeleteTokensResponse, error)
	mustEmbedUnimplementedPushServer()
}

// UnimplementedPushServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPushServer struct{}

func (UnimplementedPushServer) AddToken(context.Context, *AddTokenRequest) (*AddTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (UnimplementedPushServer) DeleteTokens(context.Context, *DeleteTokensRequest) (*DeleteTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTokens not implemented")
}
func (UnimplementedPushServer) mustEmbedUnimplementedPushServer() {}
func (UnimplementedPushServer) testEmbeddedByValue()              {}

// UnsafePushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServer will
// result in compilation errors.
type UnsafePushServer interface {
	mustEmbedUnimplementedPushServer()
}

func RegisterPushServer(s grpc.ServiceRegistrar, srv PushServer) {
	// If the following call pancis, it indicates UnimplementedPushServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Push_ServiceDesc, srv)
}

func _Push_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Push_AddToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).AddToken(ctx, req.(*AddTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_DeleteTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).DeleteTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Push_DeleteTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).DeleteTokens(ctx, req.(*DeleteTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Push_ServiceDesc is the grpc.ServiceDesc for Push service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Push_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipcash.push.v1.Push",
	HandlerType: (*PushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToken",
			Handler:    _Push_AddToken_Handler,
		},
		{
			MethodName: "DeleteTokens",
			Handler:    _Push_DeleteTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push/v1/push_service.proto",
}
