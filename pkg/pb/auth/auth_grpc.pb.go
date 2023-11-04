// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/pb/auth/auth.proto

package auth_pb_service

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

// AuthPbServiceClient is the client API for AuthPbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthPbServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type authPbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthPbServiceClient(cc grpc.ClientConnInterface) AuthPbServiceClient {
	return &authPbServiceClient{cc}
}

func (c *authPbServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthPbService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthPbServiceServer is the server API for AuthPbService service.
// All implementations must embed UnimplementedAuthPbServiceServer
// for forward compatibility
type AuthPbServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedAuthPbServiceServer()
}

// UnimplementedAuthPbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthPbServiceServer struct {
}

func (UnimplementedAuthPbServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthPbServiceServer) mustEmbedUnimplementedAuthPbServiceServer() {}

// UnsafeAuthPbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthPbServiceServer will
// result in compilation errors.
type UnsafeAuthPbServiceServer interface {
	mustEmbedUnimplementedAuthPbServiceServer()
}

func RegisterAuthPbServiceServer(s grpc.ServiceRegistrar, srv AuthPbServiceServer) {
	s.RegisterService(&AuthPbService_ServiceDesc, srv)
}

func _AuthPbService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthPbServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthPbService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthPbServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthPbService_ServiceDesc is the grpc.ServiceDesc for AuthPbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthPbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthPbService",
	HandlerType: (*AuthPbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthPbService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth/auth.proto",
}
