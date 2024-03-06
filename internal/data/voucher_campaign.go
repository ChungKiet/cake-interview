package data

import (
	"context"
	"github.com/ChungKiet/cake-interview/ent"
	"github.com/ChungKiet/cake-interview/ent/vouchercampaign"
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type voucherCampaignRepo struct {
	data *Data
	log  *log.Helper
}

// NewVoucherCampaignRepo .
func NewVoucherCampaignRepo(data *Data) biz.VoucherCampaignRepo {
	return &voucherCampaignRepo{
		data: data,
		log:  log.NewHelper(log.DefaultLogger),
	}
}

func (v voucherCampaignRepo) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return v.data.WithTx(ctx, fn)
}

func (v voucherCampaignRepo) CreateVoucherCampaign(ctx context.Context, req *biz.CreateVoucherCampaignRequest) error {
	//TODO implement me
	now := time.Now().UTC()
	err := v.data.db.VoucherCampaign.Create().
		SetVoucherID(req.VoucherID).
		SetCampaignID(req.CampaignID).
		SetStartDay(time.Unix(req.StartDay, 0)).
		SetValidTo(time.Unix(req.ValidTo, 0)).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Exec(ctx)

	if err != nil {
		v.log.Error("CreateVoucherCampaign failed", "err", err)
		return err
	}

	return nil
}

func (v voucherCampaignRepo) GetVoucherCampaign(ctx context.Context, campaignID int64) ([]*ent.VoucherCampaign, error) {
	vouchers, err := v.data.db.VoucherCampaign.Query().
		Where(vouchercampaign.CampaignIDEQ(campaignID)).
		All(ctx)

	if err != nil {
		v.log.Error("GetVoucherCampaign failed", "err", err)
		return nil, err
	}

	return vouchers, err
}
