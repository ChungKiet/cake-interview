package biz

import (
	"context"
	"github.com/ChungKiet/cake-interview/api/server"
	"github.com/ChungKiet/cake-interview/internal/constant"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type CampaignUseCase struct {
	voucherRepo         VoucherRepo
	voucherCampaignRepo VoucherCampaignRepo
	campaignRepo        CampaignRepo
	log                 *log.Helper
}

func NewCampaignUseCase(voucherRepo VoucherRepo, voucherCampaignRepo VoucherCampaignRepo, campaignRepo CampaignRepo) *CampaignUseCase {
	return &CampaignUseCase{
		voucherRepo:         voucherRepo,
		voucherCampaignRepo: voucherCampaignRepo,
		campaignRepo:        campaignRepo,
		log:                 log.NewHelper(log.DefaultLogger),
	}
}

func (uc *CampaignUseCase) CreateCampaign(ctx context.Context, req *server.CreateCampaignRequest) (int64, error) {
	if req.Name == constant.EMPTY_STRING {
		return -1, constant.ErrVoucherCampaignNameInvalid
	}

	now := time.Now().UTC()
	nowUnix := now.Unix()
	if req.StartDay < nowUnix {
		return -1, constant.ErrVoucherCampaignStartDayInvalid
	}

	if req.EndDay < now.Unix() {
		return -1, constant.ErrVoucherCampaignEndDayInvalid
	}

	if req.EndDay <= req.StartDay {
		return -1, constant.ErrVoucherCampaignLiveTimeInvalid
	}

	if req.TotalSlot <= constant.ZERO_LENGTH {
		return -1, constant.ErrVoucherCampaignTotalSlotInvalid
	}

	if len(req.Vouchers) == constant.ZERO_LENGTH {
		return -1, constant.ErrCampaignMustHavVouchers
	}

	var campaignId int64
	errTx := uc.voucherCampaignRepo.WithTx(ctx, func(ctx context.Context) error {
		reqVoucherIds := mappingVoucherRequestDataToVoucherIDs(req.Vouchers)
		vouchers, err := uc.voucherRepo.GetVoucherByIDs(ctx, reqVoucherIds)
		if err != nil {
			return err
		}

		if len(vouchers) < len(req.Vouchers) {
			var voucherIds []int
			for _, item := range vouchers {
				voucherIds = append(voucherIds, item.ID)
			}
			uc.log.Error("[CreateCampaign] GetVoucherByIDs failed", "request_ids", reqVoucherIds, "result_voucher_id", voucherIds)
			return constant.ErrVoucherNotFound
		}

		// create campaign
		campaignId, err = uc.campaignRepo.CreateCampaign(ctx, &CreateCampaignRequest{
			Name:      req.Name,
			StartDay:  req.StartDay,
			EndDay:    req.EndDay,
			TotalSlot: req.TotalSlot,
		})

		if err != nil {
			return err
		}

		// assume that voucher valid time is campaign valid time
		for _, item := range vouchers {
			err = uc.voucherCampaignRepo.CreateVoucherCampaign(ctx, &CreateVoucherCampaignRequest{
				CampaignID: campaignId,
				VoucherID:  int64(item.ID),
				StartDay:   req.StartDay,
				ValidTo:    req.EndDay,
			})

			if err != nil {
				return err
			}
		}

		return nil
		// create voucher campagin
	})
	if errTx != nil {
		uc.log.Error("[CreateCampaign] Create campaign failed", "err", errTx.Error())
		return -1, errTx
	}

	return campaignId, nil
}

func mappingVoucherRequestDataToVoucherIDs(vouchers []*server.CreateCampaignRequest_Data) []int {
	var res []int
	for _, item := range vouchers {
		res = append(res, int(item.VoucherId))
	}

	return res
}
