package service

import (
	"context"
	pb "github.com/ChungKiet/cake-interview/api/server"
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/ChungKiet/cake-interview/internal/response/campaign/create_new_campaign"
	"github.com/ChungKiet/cake-interview/internal/response/user/login"
	"github.com/ChungKiet/cake-interview/internal/response/user/register"
	"github.com/ChungKiet/cake-interview/internal/response/voucher/create_new_voucher"
	log2 "github.com/go-kratos/kratos/v2/log"
)

type CAKEServiceService struct {
	pb.UnimplementedCAKEServiceServer

	log             *log2.Helper
	userUseCase     *biz.UserUseCase
	campaignUseCase *biz.CampaignUseCase
	voucherUseCase  *biz.VoucherUseCase
}

func NewCAKEServiceService(userUseCase *biz.UserUseCase, voucherUseCase *biz.VoucherUseCase, campaignUseCase *biz.CampaignUseCase) *CAKEServiceService {
	return &CAKEServiceService{
		log:             log2.NewHelper(log2.DefaultLogger),
		userUseCase:     userUseCase,
		voucherUseCase:  voucherUseCase,
		campaignUseCase: campaignUseCase,
	}
}

func (s *CAKEServiceService) Register(ctx context.Context, req *pb.RegisterNewAccountRequest) (*pb.RegisterNewAccountResponse, error) {
	err := s.userUseCase.Register(ctx, req)
	if err != nil {
		s.log.Error("Register failed", "err", err)
		return register.RaiseError(500, "REGISTER_FAILED")
	}

	return register.RaiseSuccess()
}
func (s *CAKEServiceService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := s.userUseCase.Login(ctx, req)
	if err != nil {
		s.log.Error("Login failed", "err", err)
		return login.RaiseError(500, "LOGIN_FAILED")
	}

	return login.RaiseSuccess()
}

func (s *CAKEServiceService) CreateVoucher(ctx context.Context, req *pb.CreateVoucherRequest) (*pb.CreateVoucherResponse, error) {
	err := s.voucherUseCase.CreateVoucher(ctx, req)
	if err != nil {
		s.log.Error("Login failed", "err", err)
		return create_new_voucher.RaiseError(500, "CREATE_VOUCHER_FAILED")
	}

	return create_new_voucher.RaiseSuccess()
}

func (s *CAKEServiceService) CreateCampaign(ctx context.Context, req *pb.CreateCampaignRequest) (*pb.CreateCampaignResponse, error) {
	campaignID, err := s.campaignUseCase.CreateCampaign(ctx, req)
	if err != nil {
		s.log.Error("Login failed", "err", err)
		return create_new_campaign.RaiseError(500, "CREATE_CAMPAIGN_FAILED")
	}

	return create_new_campaign.RaiseSuccess("OK", campaignID)
}
