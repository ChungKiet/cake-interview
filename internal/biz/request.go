package biz

import "time"

type CreateUserRequest struct {
	FullName    string
	UserName    string
	PhoneNumber string
	Email       string
	Password    string // encrypted password
	BirthDay    time.Time
	Account     string
}

type GetUserRequest struct {
	Account string
}

type CreateVoucherRequest struct {
	Name  string
	Type  string
	Value float32
}

type CreateCampaignRequest struct {
	Name      string
	StartDay  int64
	EndDay    int64
	TotalSlot int64
}

type CreateVoucherCampaignRequest struct {
	CampaignID int64
	VoucherID  int64
	StartDay   int64
	ValidTo    int64
}

type CreateUserVoucherRequest struct {
	VoucherCampaignID int64
	UserID            string
}

type IncreaseCampaignSlotRequest struct {
	CampaignID  int64
	IsFull      bool
	CurrentSlot int64
}
