// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: activity/v1/activity_feed_service.proto

package activitypb

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
	ActivityFeed_GetLatestNotifications_FullMethodName = "/flipcash.activity.v1.ActivityFeed/GetLatestNotifications"
	ActivityFeed_GetPagedNotifications_FullMethodName  = "/flipcash.activity.v1.ActivityFeed/GetPagedNotifications"
	ActivityFeed_GetBatchNotifications_FullMethodName  = "/flipcash.activity.v1.ActivityFeed/GetBatchNotifications"
)

// ActivityFeedClient is the client API for ActivityFeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityFeedClient interface {
	// GetLatestNotifications gets the latest N notifications in a user's
	// activity feed. Results will be ordered by descending timestamp.
	GetLatestNotifications(ctx context.Context, in *GetLatestNotificationsRequest, opts ...grpc.CallOption) (*GetLatestNotificationsResponse, error)
	// GetPagedNotifications gets all notifications using a paging API.
	GetPagedNotifications(ctx context.Context, in *GetPagedNotificationsRequest, opts ...grpc.CallOption) (*GetPagedNotificationsResponse, error)
	// GetBatchNotifications gets a batch of notifications by ID.
	GetBatchNotifications(ctx context.Context, in *GetBatchNotificationsRequest, opts ...grpc.CallOption) (*GetBatchNotificationsResponse, error)
}

type activityFeedClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityFeedClient(cc grpc.ClientConnInterface) ActivityFeedClient {
	return &activityFeedClient{cc}
}

func (c *activityFeedClient) GetLatestNotifications(ctx context.Context, in *GetLatestNotificationsRequest, opts ...grpc.CallOption) (*GetLatestNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLatestNotificationsResponse)
	err := c.cc.Invoke(ctx, ActivityFeed_GetLatestNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityFeedClient) GetPagedNotifications(ctx context.Context, in *GetPagedNotificationsRequest, opts ...grpc.CallOption) (*GetPagedNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPagedNotificationsResponse)
	err := c.cc.Invoke(ctx, ActivityFeed_GetPagedNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityFeedClient) GetBatchNotifications(ctx context.Context, in *GetBatchNotificationsRequest, opts ...grpc.CallOption) (*GetBatchNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBatchNotificationsResponse)
	err := c.cc.Invoke(ctx, ActivityFeed_GetBatchNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityFeedServer is the server API for ActivityFeed service.
// All implementations must embed UnimplementedActivityFeedServer
// for forward compatibility.
type ActivityFeedServer interface {
	// GetLatestNotifications gets the latest N notifications in a user's
	// activity feed. Results will be ordered by descending timestamp.
	GetLatestNotifications(context.Context, *GetLatestNotificationsRequest) (*GetLatestNotificationsResponse, error)
	// GetPagedNotifications gets all notifications using a paging API.
	GetPagedNotifications(context.Context, *GetPagedNotificationsRequest) (*GetPagedNotificationsResponse, error)
	// GetBatchNotifications gets a batch of notifications by ID.
	GetBatchNotifications(context.Context, *GetBatchNotificationsRequest) (*GetBatchNotificationsResponse, error)
	mustEmbedUnimplementedActivityFeedServer()
}

// UnimplementedActivityFeedServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityFeedServer struct{}

func (UnimplementedActivityFeedServer) GetLatestNotifications(context.Context, *GetLatestNotificationsRequest) (*GetLatestNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestNotifications not implemented")
}
func (UnimplementedActivityFeedServer) GetPagedNotifications(context.Context, *GetPagedNotificationsRequest) (*GetPagedNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPagedNotifications not implemented")
}
func (UnimplementedActivityFeedServer) GetBatchNotifications(context.Context, *GetBatchNotificationsRequest) (*GetBatchNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchNotifications not implemented")
}
func (UnimplementedActivityFeedServer) mustEmbedUnimplementedActivityFeedServer() {}
func (UnimplementedActivityFeedServer) testEmbeddedByValue()                      {}

// UnsafeActivityFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityFeedServer will
// result in compilation errors.
type UnsafeActivityFeedServer interface {
	mustEmbedUnimplementedActivityFeedServer()
}

func RegisterActivityFeedServer(s grpc.ServiceRegistrar, srv ActivityFeedServer) {
	// If the following call pancis, it indicates UnimplementedActivityFeedServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActivityFeed_ServiceDesc, srv)
}

func _ActivityFeed_GetLatestNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityFeedServer).GetLatestNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityFeed_GetLatestNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityFeedServer).GetLatestNotifications(ctx, req.(*GetLatestNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityFeed_GetPagedNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPagedNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityFeedServer).GetPagedNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityFeed_GetPagedNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityFeedServer).GetPagedNotifications(ctx, req.(*GetPagedNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityFeed_GetBatchNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityFeedServer).GetBatchNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityFeed_GetBatchNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityFeedServer).GetBatchNotifications(ctx, req.(*GetBatchNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityFeed_ServiceDesc is the grpc.ServiceDesc for ActivityFeed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityFeed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipcash.activity.v1.ActivityFeed",
	HandlerType: (*ActivityFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestNotifications",
			Handler:    _ActivityFeed_GetLatestNotifications_Handler,
		},
		{
			MethodName: "GetPagedNotifications",
			Handler:    _ActivityFeed_GetPagedNotifications_Handler,
		},
		{
			MethodName: "GetBatchNotifications",
			Handler:    _ActivityFeed_GetBatchNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity/v1/activity_feed_service.proto",
}
