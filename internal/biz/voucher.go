package biz

import (
	"context"
	"github.com/ChungKiet/cake-interview/api/server"
	"github.com/ChungKiet/cake-interview/internal/constant"
	"github.com/go-kratos/kratos/v2/log"
)

type VoucherUseCase struct {
	repo VoucherRepo
	log  *log.Helper
}

func NewVoucherUseCase(repo VoucherRepo) *VoucherUseCase {
	return &VoucherUseCase{
		repo: repo,
		log:  log.NewHelper(log.DefaultLogger),
	}
}

func (uc *VoucherUseCase) CreateVoucher(ctx context.Context, req *server.CreateVoucherRequest) error {
	if req.Name == constant.EMPTY_STRING {
		return constant.ErrVoucherNameRequired
	}

	if req.Type == constant.EMPTY_STRING {
		return constant.ErrVoucherTypeRequired
	}

	if ok := constant.MapVoucherType[req.Type]; !ok {
		return constant.ErrVoucherTypeInvalid
	}

	if req.Value < constant.MIN_VOUCHER_VALUE || req.Value > constant.MAX_VOUCHER_VALUE {
		return constant.ErrVoucherValueInvalid
	}

	return uc.repo.CreateVoucher(ctx, req)
}
