// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pool/v1/pool_service.proto

package poolpb

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
	Pool_CreatePool_FullMethodName    = "/flipcash.pool.v1.Pool/CreatePool"
	Pool_GetPool_FullMethodName       = "/flipcash.pool.v1.Pool/GetPool"
	Pool_GetPagedPools_FullMethodName = "/flipcash.pool.v1.Pool/GetPagedPools"
	Pool_ClosePool_FullMethodName     = "/flipcash.pool.v1.Pool/ClosePool"
	Pool_ResolvePool_FullMethodName   = "/flipcash.pool.v1.Pool/ResolvePool"
	Pool_MakeBet_FullMethodName       = "/flipcash.pool.v1.Pool/MakeBet"
)

// PoolClient is the client API for Pool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoolClient interface {
	// CreatePool creates a new pool
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error)
	// GetPool gets pool metadata by its ID
	GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error)
	// GetPagedPools gets all pools for a user over a paging API
	//
	// Note: Only consolidated metadata is provided in the response
	GetPagedPools(ctx context.Context, in *GetPagedPoolsRequest, opts ...grpc.CallOption) (*GetPagedPoolsResponse, error)
	// ClosePool closes a pool from additional bets
	ClosePool(ctx context.Context, in *ClosePoolRequest, opts ...grpc.CallOption) (*ClosePoolResponse, error)
	// ResolvePool resolves a pool by declaring the pool's outcome. The pool creator
	// resolves a pool by calling this RPC first, then SubmitIntent to distribute funds
	// to the winning participants.
	ResolvePool(ctx context.Context, in *ResolvePoolRequest, opts ...grpc.CallOption) (*ResolvePoolResponse, error)
	// MakeBet creates a new bet against a pool. Pool participants make a bet by
	// calling MakeBet to create an initially unpaid bet, then SubmitIntent for
	// payment where:
	//  1. Intent ID == Bet.id
	//  2. Payment amount == PoolMetadata.buy_in
	//  3. Payment destination == PoolMetadata.funding_destination
	//
	// Bets can be changed as long as payment has not been received. Clients must
	// use the same Bet ID when updating their bet.
	MakeBet(ctx context.Context, in *MakeBetRequest, opts ...grpc.CallOption) (*MakeBetResponse, error)
}

type poolClient struct {
	cc grpc.ClientConnInterface
}

func NewPoolClient(cc grpc.ClientConnInterface) PoolClient {
	return &poolClient{cc}
}

func (c *poolClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePoolResponse)
	err := c.cc.Invoke(ctx, Pool_CreatePool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolClient) GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPoolResponse)
	err := c.cc.Invoke(ctx, Pool_GetPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolClient) GetPagedPools(ctx context.Context, in *GetPagedPoolsRequest, opts ...grpc.CallOption) (*GetPagedPoolsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPagedPoolsResponse)
	err := c.cc.Invoke(ctx, Pool_GetPagedPools_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolClient) ClosePool(ctx context.Context, in *ClosePoolRequest, opts ...grpc.CallOption) (*ClosePoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClosePoolResponse)
	err := c.cc.Invoke(ctx, Pool_ClosePool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolClient) ResolvePool(ctx context.Context, in *ResolvePoolRequest, opts ...grpc.CallOption) (*ResolvePoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolvePoolResponse)
	err := c.cc.Invoke(ctx, Pool_ResolvePool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolClient) MakeBet(ctx context.Context, in *MakeBetRequest, opts ...grpc.CallOption) (*MakeBetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MakeBetResponse)
	err := c.cc.Invoke(ctx, Pool_MakeBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolServer is the server API for Pool service.
// All implementations must embed UnimplementedPoolServer
// for forward compatibility.
type PoolServer interface {
	// CreatePool creates a new pool
	CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error)
	// GetPool gets pool metadata by its ID
	GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error)
	// GetPagedPools gets all pools for a user over a paging API
	//
	// Note: Only consolidated metadata is provided in the response
	GetPagedPools(context.Context, *GetPagedPoolsRequest) (*GetPagedPoolsResponse, error)
	// ClosePool closes a pool from additional bets
	ClosePool(context.Context, *ClosePoolRequest) (*ClosePoolResponse, error)
	// ResolvePool resolves a pool by declaring the pool's outcome. The pool creator
	// resolves a pool by calling this RPC first, then SubmitIntent to distribute funds
	// to the winning participants.
	ResolvePool(context.Context, *ResolvePoolRequest) (*ResolvePoolResponse, error)
	// MakeBet creates a new bet against a pool. Pool participants make a bet by
	// calling MakeBet to create an initially unpaid bet, then SubmitIntent for
	// payment where:
	//  1. Intent ID == Bet.id
	//  2. Payment amount == PoolMetadata.buy_in
	//  3. Payment destination == PoolMetadata.funding_destination
	//
	// Bets can be changed as long as payment has not been received. Clients must
	// use the same Bet ID when updating their bet.
	MakeBet(context.Context, *MakeBetRequest) (*MakeBetResponse, error)
	mustEmbedUnimplementedPoolServer()
}

// UnimplementedPoolServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPoolServer struct{}

func (UnimplementedPoolServer) CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedPoolServer) GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPool not implemented")
}
func (UnimplementedPoolServer) GetPagedPools(context.Context, *GetPagedPoolsRequest) (*GetPagedPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPagedPools not implemented")
}
func (UnimplementedPoolServer) ClosePool(context.Context, *ClosePoolRequest) (*ClosePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePool not implemented")
}
func (UnimplementedPoolServer) ResolvePool(context.Context, *ResolvePoolRequest) (*ResolvePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolvePool not implemented")
}
func (UnimplementedPoolServer) MakeBet(context.Context, *MakeBetRequest) (*MakeBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeBet not implemented")
}
func (UnimplementedPoolServer) mustEmbedUnimplementedPoolServer() {}
func (UnimplementedPoolServer) testEmbeddedByValue()              {}

// UnsafePoolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoolServer will
// result in compilation errors.
type UnsafePoolServer interface {
	mustEmbedUnimplementedPoolServer()
}

func RegisterPoolServer(s grpc.ServiceRegistrar, srv PoolServer) {
	// If the following call pancis, it indicates UnimplementedPoolServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Pool_ServiceDesc, srv)
}

func _Pool_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pool_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pool_GetPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServer).GetPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pool_GetPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServer).GetPool(ctx, req.(*GetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pool_GetPagedPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPagedPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServer).GetPagedPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pool_GetPagedPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServer).GetPagedPools(ctx, req.(*GetPagedPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pool_ClosePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServer).ClosePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pool_ClosePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServer).ClosePool(ctx, req.(*ClosePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pool_ResolvePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolvePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServer).ResolvePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pool_ResolvePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServer).ResolvePool(ctx, req.(*ResolvePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pool_MakeBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServer).MakeBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pool_MakeBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServer).MakeBet(ctx, req.(*MakeBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pool_ServiceDesc is the grpc.ServiceDesc for Pool service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pool_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipcash.pool.v1.Pool",
	HandlerType: (*PoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Pool_CreatePool_Handler,
		},
		{
			MethodName: "GetPool",
			Handler:    _Pool_GetPool_Handler,
		},
		{
			MethodName: "GetPagedPools",
			Handler:    _Pool_GetPagedPools_Handler,
		},
		{
			MethodName: "ClosePool",
			Handler:    _Pool_ClosePool_Handler,
		},
		{
			MethodName: "ResolvePool",
			Handler:    _Pool_ResolvePool_Handler,
		},
		{
			MethodName: "MakeBet",
			Handler:    _Pool_MakeBet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool/v1/pool_service.proto",
}
