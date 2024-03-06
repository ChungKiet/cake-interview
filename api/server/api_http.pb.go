// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.9
// source: server/api.proto

package server

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCAKEServiceCreateCampaign = "/CAKEService/CreateCampaign"
const OperationCAKEServiceCreateVoucher = "/CAKEService/CreateVoucher"
const OperationCAKEServiceLogin = "/CAKEService/Login"
const OperationCAKEServiceRegister = "/CAKEService/Register"

type CAKEServiceHTTPServer interface {
	CreateCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error)
	// CreateVoucher Internal api
	CreateVoucher(context.Context, *CreateVoucherRequest) (*CreateVoucherResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Register Sends a greeting
	Register(context.Context, *RegisterNewAccountRequest) (*RegisterNewAccountResponse, error)
}

func RegisterCAKEServiceHTTPServer(s *http.Server, srv CAKEServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/user/v1/register", _CAKEService_Register0_HTTP_Handler(srv))
	r.POST("/api/user/v1/login", _CAKEService_Login0_HTTP_Handler(srv))
	r.POST("/api/voucher/v1/create", _CAKEService_CreateVoucher0_HTTP_Handler(srv))
	r.POST("/api/campaign/v1/create", _CAKEService_CreateCampaign0_HTTP_Handler(srv))
}

func _CAKEService_Register0_HTTP_Handler(srv CAKEServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterNewAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCAKEServiceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterNewAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterNewAccountResponse)
		return ctx.Result(200, reply)
	}
}

func _CAKEService_Login0_HTTP_Handler(srv CAKEServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCAKEServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _CAKEService_CreateVoucher0_HTTP_Handler(srv CAKEServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateVoucherRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCAKEServiceCreateVoucher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateVoucher(ctx, req.(*CreateVoucherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateVoucherResponse)
		return ctx.Result(200, reply)
	}
}

func _CAKEService_CreateCampaign0_HTTP_Handler(srv CAKEServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCampaignRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCAKEServiceCreateCampaign)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCampaign(ctx, req.(*CreateCampaignRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCampaignResponse)
		return ctx.Result(200, reply)
	}
}

type CAKEServiceHTTPClient interface {
	CreateCampaign(ctx context.Context, req *CreateCampaignRequest, opts ...http.CallOption) (rsp *CreateCampaignResponse, err error)
	CreateVoucher(ctx context.Context, req *CreateVoucherRequest, opts ...http.CallOption) (rsp *CreateVoucherResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	Register(ctx context.Context, req *RegisterNewAccountRequest, opts ...http.CallOption) (rsp *RegisterNewAccountResponse, err error)
}

type CAKEServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCAKEServiceHTTPClient(client *http.Client) CAKEServiceHTTPClient {
	return &CAKEServiceHTTPClientImpl{client}
}

func (c *CAKEServiceHTTPClientImpl) CreateCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...http.CallOption) (*CreateCampaignResponse, error) {
	var out CreateCampaignResponse
	pattern := "/api/campaign/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCAKEServiceCreateCampaign))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CAKEServiceHTTPClientImpl) CreateVoucher(ctx context.Context, in *CreateVoucherRequest, opts ...http.CallOption) (*CreateVoucherResponse, error) {
	var out CreateVoucherResponse
	pattern := "/api/voucher/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCAKEServiceCreateVoucher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CAKEServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/api/user/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCAKEServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CAKEServiceHTTPClientImpl) Register(ctx context.Context, in *RegisterNewAccountRequest, opts ...http.CallOption) (*RegisterNewAccountResponse, error) {
	var out RegisterNewAccountResponse
	pattern := "/api/user/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCAKEServiceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}