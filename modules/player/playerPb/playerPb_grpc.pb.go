// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: modules/player/playerPb/playerPb.proto

package shop_game

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
	PlayerGrpcService_CredentialSearch_FullMethodName              = "/PlayerGrpcService/CredentialSearch"
	PlayerGrpcService_FindOnePlayerProfileToRefresh_FullMethodName = "/PlayerGrpcService/FindOnePlayerProfileToRefresh"
	PlayerGrpcService_GetPlayerSavingAccount_FullMethodName        = "/PlayerGrpcService/GetPlayerSavingAccount"
)

// PlayerGrpcServiceClient is the client API for PlayerGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerGrpcServiceClient interface {
	CredentialSearch(ctx context.Context, in *CredentialSearchReq, opts ...grpc.CallOption) (*PlayerProfile, error)
	FindOnePlayerProfileToRefresh(ctx context.Context, in *FindOnePlayerProfileToRefreshReq, opts ...grpc.CallOption) (*PlayerProfile, error)
	GetPlayerSavingAccount(ctx context.Context, in *GetPlayerSavingAccountReq, opts ...grpc.CallOption) (*GetPlayerSavingAccountRes, error)
}

type playerGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerGrpcServiceClient(cc grpc.ClientConnInterface) PlayerGrpcServiceClient {
	return &playerGrpcServiceClient{cc}
}

func (c *playerGrpcServiceClient) CredentialSearch(ctx context.Context, in *CredentialSearchReq, opts ...grpc.CallOption) (*PlayerProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerProfile)
	err := c.cc.Invoke(ctx, PlayerGrpcService_CredentialSearch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerGrpcServiceClient) FindOnePlayerProfileToRefresh(ctx context.Context, in *FindOnePlayerProfileToRefreshReq, opts ...grpc.CallOption) (*PlayerProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerProfile)
	err := c.cc.Invoke(ctx, PlayerGrpcService_FindOnePlayerProfileToRefresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerGrpcServiceClient) GetPlayerSavingAccount(ctx context.Context, in *GetPlayerSavingAccountReq, opts ...grpc.CallOption) (*GetPlayerSavingAccountRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlayerSavingAccountRes)
	err := c.cc.Invoke(ctx, PlayerGrpcService_GetPlayerSavingAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerGrpcServiceServer is the server API for PlayerGrpcService service.
// All implementations must embed UnimplementedPlayerGrpcServiceServer
// for forward compatibility.
type PlayerGrpcServiceServer interface {
	CredentialSearch(context.Context, *CredentialSearchReq) (*PlayerProfile, error)
	FindOnePlayerProfileToRefresh(context.Context, *FindOnePlayerProfileToRefreshReq) (*PlayerProfile, error)
	GetPlayerSavingAccount(context.Context, *GetPlayerSavingAccountReq) (*GetPlayerSavingAccountRes, error)
	mustEmbedUnimplementedPlayerGrpcServiceServer()
}

// UnimplementedPlayerGrpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlayerGrpcServiceServer struct{}

func (UnimplementedPlayerGrpcServiceServer) CredentialSearch(context.Context, *CredentialSearchReq) (*PlayerProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialSearch not implemented")
}
func (UnimplementedPlayerGrpcServiceServer) FindOnePlayerProfileToRefresh(context.Context, *FindOnePlayerProfileToRefreshReq) (*PlayerProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOnePlayerProfileToRefresh not implemented")
}
func (UnimplementedPlayerGrpcServiceServer) GetPlayerSavingAccount(context.Context, *GetPlayerSavingAccountReq) (*GetPlayerSavingAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerSavingAccount not implemented")
}
func (UnimplementedPlayerGrpcServiceServer) mustEmbedUnimplementedPlayerGrpcServiceServer() {}
func (UnimplementedPlayerGrpcServiceServer) testEmbeddedByValue()                           {}

// UnsafePlayerGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerGrpcServiceServer will
// result in compilation errors.
type UnsafePlayerGrpcServiceServer interface {
	mustEmbedUnimplementedPlayerGrpcServiceServer()
}

func RegisterPlayerGrpcServiceServer(s grpc.ServiceRegistrar, srv PlayerGrpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedPlayerGrpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PlayerGrpcService_ServiceDesc, srv)
}

func _PlayerGrpcService_CredentialSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerGrpcServiceServer).CredentialSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerGrpcService_CredentialSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerGrpcServiceServer).CredentialSearch(ctx, req.(*CredentialSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerGrpcService_FindOnePlayerProfileToRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOnePlayerProfileToRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerGrpcServiceServer).FindOnePlayerProfileToRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerGrpcService_FindOnePlayerProfileToRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerGrpcServiceServer).FindOnePlayerProfileToRefresh(ctx, req.(*FindOnePlayerProfileToRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerGrpcService_GetPlayerSavingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerSavingAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerGrpcServiceServer).GetPlayerSavingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerGrpcService_GetPlayerSavingAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerGrpcServiceServer).GetPlayerSavingAccount(ctx, req.(*GetPlayerSavingAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerGrpcService_ServiceDesc is the grpc.ServiceDesc for PlayerGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PlayerGrpcService",
	HandlerType: (*PlayerGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialSearch",
			Handler:    _PlayerGrpcService_CredentialSearch_Handler,
		},
		{
			MethodName: "FindOnePlayerProfileToRefresh",
			Handler:    _PlayerGrpcService_FindOnePlayerProfileToRefresh_Handler,
		},
		{
			MethodName: "GetPlayerSavingAccount",
			Handler:    _PlayerGrpcService_GetPlayerSavingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/player/playerPb/playerPb.proto",
}
