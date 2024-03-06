package biz

import (
	"context"
	"github.com/ChungKiet/cake-interview/api/server"
	"github.com/ChungKiet/cake-interview/ent"
)

type Tx interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type UserRepo interface {
	Tx
	CreateUser(ctx context.Context, req *CreateUserRequest) (string, error)
	GetUser(ctx context.Context, req *GetUserRequest) (*ent.User, error)
	UpdateLoginTime(ctx context.Context, account string) error
}

type VoucherRepo interface {
	Tx
	CreateVoucher(ctx context.Context, req *server.CreateVoucherRequest) error
	GetVoucherByIDs(ctx context.Context, ids []int) ([]*ent.Voucher, error)
}

type CampaignRepo interface {
	Tx
	CreateCampaign(ctx context.Context, request *CreateCampaignRequest) (int64, error)
	GetCampaignNotFull(ctx context.Context) (*ent.Campaign, error)
	IncreaseCampaignSlot(ctx context.Context, req *IncreaseCampaignSlotRequest) error
}

type VoucherCampaignRepo interface {
	Tx
	CreateVoucherCampaign(ctx context.Context, req *CreateVoucherCampaignRequest) error
	GetVoucherCampaign(ctx context.Context, campaignID int64) ([]*ent.VoucherCampaign, error)
}

type UserVoucherRepo interface {
	Tx
	CreateUserVoucher(ctx context.Context, req *CreateUserVoucherRequest) error
}
