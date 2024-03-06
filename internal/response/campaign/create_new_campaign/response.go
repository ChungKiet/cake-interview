package create_new_campaign

import (
	"errors"
	"github.com/ChungKiet/cake-interview/api/server"
)

func RaiseError(status int32, msg string) (*server.CreateCampaignResponse, error) {
	return &server.CreateCampaignResponse{
		Status:  status,
		Message: msg,
		MsgKey:  msg,
	}, errors.New(msg)
}

func RaiseSuccess(msg string, campaignId int64) (*server.CreateCampaignResponse, error) {
	return &server.CreateCampaignResponse{
		Status: 200,
		MsgKey: msg,
		Data: &server.CreateCampaignResponse_Data{
			Id: campaignId,
		},
	}, nil
}
