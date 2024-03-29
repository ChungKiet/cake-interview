// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: server/api.proto

package server

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
	CAKEService_Register_FullMethodName       = "/CAKEService/Register"
	CAKEService_Login_FullMethodName          = "/CAKEService/Login"
	CAKEService_CreateVoucher_FullMethodName  = "/CAKEService/CreateVoucher"
	CAKEService_CreateCampaign_FullMethodName = "/CAKEService/CreateCampaign"
)

// CAKEServiceClient is the client API for CAKEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CAKEServiceClient interface {
	// Sends a greeting
	Register(ctx context.Context, in *RegisterNewAccountRequest, opts ...grpc.CallOption) (*RegisterNewAccountResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Internal api
	CreateVoucher(ctx context.Context, in *CreateVoucherRequest, opts ...grpc.CallOption) (*CreateVoucherResponse, error)
	CreateCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error)
}

type cAKEServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCAKEServiceClient(cc grpc.ClientConnInterface) CAKEServiceClient {
	return &cAKEServiceClient{cc}
}

func (c *cAKEServiceClient) Register(ctx context.Context, in *RegisterNewAccountRequest, opts ...grpc.CallOption) (*RegisterNewAccountResponse, error) {
	out := new(RegisterNewAccountResponse)
	err := c.cc.Invoke(ctx, CAKEService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAKEServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, CAKEService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAKEServiceClient) CreateVoucher(ctx context.Context, in *CreateVoucherRequest, opts ...grpc.CallOption) (*CreateVoucherResponse, error) {
	out := new(CreateVoucherResponse)
	err := c.cc.Invoke(ctx, CAKEService_CreateVoucher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAKEServiceClient) CreateCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error) {
	out := new(CreateCampaignResponse)
	err := c.cc.Invoke(ctx, CAKEService_CreateCampaign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CAKEServiceServer is the server API for CAKEService service.
// All implementations must embed UnimplementedCAKEServiceServer
// for forward compatibility
type CAKEServiceServer interface {
	// Sends a greeting
	Register(context.Context, *RegisterNewAccountRequest) (*RegisterNewAccountResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Internal api
	CreateVoucher(context.Context, *CreateVoucherRequest) (*CreateVoucherResponse, error)
	CreateCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error)
	mustEmbedUnimplementedCAKEServiceServer()
}

// UnimplementedCAKEServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCAKEServiceServer struct {
}

func (UnimplementedCAKEServiceServer) Register(context.Context, *RegisterNewAccountRequest) (*RegisterNewAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCAKEServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCAKEServiceServer) CreateVoucher(context.Context, *CreateVoucherRequest) (*CreateVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVoucher not implemented")
}
func (UnimplementedCAKEServiceServer) CreateCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCampaign not implemented")
}
func (UnimplementedCAKEServiceServer) mustEmbedUnimplementedCAKEServiceServer() {}

// UnsafeCAKEServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CAKEServiceServer will
// result in compilation errors.
type UnsafeCAKEServiceServer interface {
	mustEmbedUnimplementedCAKEServiceServer()
}

func RegisterCAKEServiceServer(s grpc.ServiceRegistrar, srv CAKEServiceServer) {
	s.RegisterService(&CAKEService_ServiceDesc, srv)
}

func _CAKEService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNewAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAKEServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAKEService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAKEServiceServer).Register(ctx, req.(*RegisterNewAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAKEService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAKEServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAKEService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAKEServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAKEService_CreateVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAKEServiceServer).CreateVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAKEService_CreateVoucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAKEServiceServer).CreateVoucher(ctx, req.(*CreateVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAKEService_CreateCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAKEServiceServer).CreateCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAKEService_CreateCampaign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAKEServiceServer).CreateCampaign(ctx, req.(*CreateCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CAKEService_ServiceDesc is the grpc.ServiceDesc for CAKEService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CAKEService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CAKEService",
	HandlerType: (*CAKEServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _CAKEService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CAKEService_Login_Handler,
		},
		{
			MethodName: "CreateVoucher",
			Handler:    _CAKEService_CreateVoucher_Handler,
		},
		{
			MethodName: "CreateCampaign",
			Handler:    _CAKEService_CreateCampaign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/api.proto",
}
