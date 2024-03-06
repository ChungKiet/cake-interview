package data

import (
	"context"
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type userVoucherRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserVoucherRepo .
func NewUserVoucherRepo(data *Data) biz.UserVoucherRepo {
	return &userVoucherRepo{
		data: data,
		log:  log.NewHelper(log.DefaultLogger),
	}
}

func (u userVoucherRepo) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return u.data.WithTx(ctx, fn)
}

func (u userVoucherRepo) CreateUserVoucher(ctx context.Context, req *biz.CreateUserVoucherRequest) error {
	now := time.Now().UTC()
	err := u.data.db.UserVoucher.Create().
		SetUserID(req.UserID).
		SetVoucherCampaignID(req.VoucherCampaignID).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Exec(ctx)

	if err != nil {
		u.log.Error("CreateUserVoucher failed", "err", err)
		return err
	}

	return nil
}
