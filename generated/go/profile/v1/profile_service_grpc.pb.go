// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: profile/v1/profile_service.proto

package profilepb

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
	Profile_GetProfile_FullMethodName          = "/flipcash.profile.v1.Profile/GetProfile"
	Profile_SetDisplayName_FullMethodName      = "/flipcash.profile.v1.Profile/SetDisplayName"
	Profile_LinkSocialAccount_FullMethodName   = "/flipcash.profile.v1.Profile/LinkSocialAccount"
	Profile_UnlinkSocialAccount_FullMethodName = "/flipcash.profile.v1.Profile/UnlinkSocialAccount"
)

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	SetDisplayName(ctx context.Context, in *SetDisplayNameRequest, opts ...grpc.CallOption) (*SetDisplayNameResponse, error)
	// LinkSocialAccount links a social account to a user
	LinkSocialAccount(ctx context.Context, in *LinkSocialAccountRequest, opts ...grpc.CallOption) (*LinkSocialAccountResponse, error)
	// UnlinkSocialAccount removes a social account link from a user
	UnlinkSocialAccount(ctx context.Context, in *UnlinkSocialAccountRequest, opts ...grpc.CallOption) (*UnlinkSocialAccountResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) SetDisplayName(ctx context.Context, in *SetDisplayNameRequest, opts ...grpc.CallOption) (*SetDisplayNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetDisplayNameResponse)
	err := c.cc.Invoke(ctx, Profile_SetDisplayName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) LinkSocialAccount(ctx context.Context, in *LinkSocialAccountRequest, opts ...grpc.CallOption) (*LinkSocialAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LinkSocialAccountResponse)
	err := c.cc.Invoke(ctx, Profile_LinkSocialAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UnlinkSocialAccount(ctx context.Context, in *UnlinkSocialAccountRequest, opts ...grpc.CallOption) (*UnlinkSocialAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnlinkSocialAccountResponse)
	err := c.cc.Invoke(ctx, Profile_UnlinkSocialAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility.
type ProfileServer interface {
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error)
	// LinkSocialAccount links a social account to a user
	LinkSocialAccount(context.Context, *LinkSocialAccountRequest) (*LinkSocialAccountResponse, error)
	// UnlinkSocialAccount removes a social account link from a user
	UnlinkSocialAccount(context.Context, *UnlinkSocialAccountRequest) (*UnlinkSocialAccountResponse, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfileServer struct{}

func (UnimplementedProfileServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfileServer) SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDisplayName not implemented")
}
func (UnimplementedProfileServer) LinkSocialAccount(context.Context, *LinkSocialAccountRequest) (*LinkSocialAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkSocialAccount not implemented")
}
func (UnimplementedProfileServer) UnlinkSocialAccount(context.Context, *UnlinkSocialAccountRequest) (*UnlinkSocialAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkSocialAccount not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}
func (UnimplementedProfileServer) testEmbeddedByValue()                 {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	// If the following call pancis, it indicates UnimplementedProfileServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_SetDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDisplayNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).SetDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_SetDisplayName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).SetDisplayName(ctx, req.(*SetDisplayNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_LinkSocialAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkSocialAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).LinkSocialAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_LinkSocialAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).LinkSocialAccount(ctx, req.(*LinkSocialAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UnlinkSocialAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkSocialAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UnlinkSocialAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UnlinkSocialAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UnlinkSocialAccount(ctx, req.(*UnlinkSocialAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipcash.profile.v1.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _Profile_GetProfile_Handler,
		},
		{
			MethodName: "SetDisplayName",
			Handler:    _Profile_SetDisplayName_Handler,
		},
		{
			MethodName: "LinkSocialAccount",
			Handler:    _Profile_LinkSocialAccount_Handler,
		},
		{
			MethodName: "UnlinkSocialAccount",
			Handler:    _Profile_UnlinkSocialAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile/v1/profile_service.proto",
}
