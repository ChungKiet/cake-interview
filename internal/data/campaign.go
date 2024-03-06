package data

import (
	"context"
	"github.com/ChungKiet/cake-interview/ent"
	campaign2 "github.com/ChungKiet/cake-interview/ent/campaign"
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type campaignRepo struct {
	data *Data
	log  *log.Helper
}

// NewCampaignRepo .
func NewCampaignRepo(data *Data) biz.CampaignRepo {
	return &campaignRepo{
		data: data,
		log:  log.NewHelper(log.DefaultLogger),
	}
}

func (c campaignRepo) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	//TODO implement me
	return c.data.WithTx(ctx, fn)
}

func (c campaignRepo) CreateCampaign(ctx context.Context, req *biz.CreateCampaignRequest) (int64, error) {
	//TODO implement me
	now := time.Now().UTC()
	campaign, err := c.data.db.Campaign.Create().
		SetName(req.Name).
		SetStartDay(time.Unix(req.StartDay, 0)).
		SetEndDay(time.Unix(req.EndDay, 0)).
		SetCurrentSlot(0).
		SetIsFull(false).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetTotalSlot(req.TotalSlot).
		Save(ctx)

	if err != nil {
		c.log.Error("CreateCampaign failed", "err", err)
		return -1, err
	}

	return int64(campaign.ID), nil
}

// assume that just exec only 1 campain / event
func (c campaignRepo) GetCampaignNotFull(ctx context.Context) (*ent.Campaign, error) {
	campaign, err := c.data.db.Campaign.Query().
		Where(campaign2.IsFullEQ(false)).
		Only(ctx)

	if err != nil {
		c.log.Error("GetCampaignNotFull failed", "err", err)
		return nil, err
	}

	return campaign, nil
}

func (c campaignRepo) IncreaseCampaignSlot(ctx context.Context, req *biz.IncreaseCampaignSlotRequest) error {
	//TODO implement me
	err := c.data.db.Campaign.Update().
		Where(campaign2.IDEQ(int(req.CampaignID))).
		SetIsFull(req.IsFull).
		SetCurrentSlot(req.CurrentSlot).
		Exec(ctx)

	if err != nil {
		c.log.Error("IncreaseCampaignSlot failed", "err", err)
		return err
	}

	return nil
}
