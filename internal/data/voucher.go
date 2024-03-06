package data

import (
	"context"
	"github.com/ChungKiet/cake-interview/api/server"
	"github.com/ChungKiet/cake-interview/ent"
	"github.com/ChungKiet/cake-interview/ent/voucher"
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type voucherRepo struct {
	data *Data
	log  *log.Helper
}

// NewVoucherRepo .
func NewVoucherRepo(data *Data) biz.VoucherRepo {
	return &voucherRepo{
		data: data,
		log:  log.NewHelper(log.DefaultLogger),
	}
}

func (v voucherRepo) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return v.data.WithTx(ctx, fn)
}

func (v voucherRepo) CreateVoucher(ctx context.Context, req *server.CreateVoucherRequest) error {
	err := v.data.db.Voucher.Create().
		SetName(req.Name).
		SetType(req.Type).
		SetValue(float64(req.Value)).
		SetCreatedAt(time.Now().UTC()).
		Exec(ctx)

	if err != nil {
		v.log.Error("CreateVoucher failed", "err", err)
		return err
	}

	return nil
}

func (v voucherRepo) GetVoucherByIDs(ctx context.Context, ids []int) ([]*ent.Voucher, error) {
	vouchers, err := v.data.db.Voucher.Query().
		Where(voucher.IDIn(ids...)).
		All(ctx)

	if err != nil {
		v.log.Error("GetVoucherByIDs failed", "err", err)
		return nil, err
	}

	return vouchers, nil
}
