package create_new_voucher

import (
	"errors"
	"github.com/ChungKiet/cake-interview/api/server"
)

func RaiseError(status int64, msg string) (*server.CreateVoucherResponse, error) {
	return &server.CreateVoucherResponse{
		Status:  status,
		Message: msg,
		MsgKey:  msg,
	}, errors.New(msg)
}

func RaiseSuccess() (*server.CreateVoucherResponse, error) {
	return &server.CreateVoucherResponse{
		Status: 200,
		MsgKey: "OK",
	}, nil
}
