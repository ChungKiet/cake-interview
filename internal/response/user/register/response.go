package register

import (
	"errors"
	"github.com/ChungKiet/cake-interview/api/server"
)

func RaiseError(status int32, msg string) (*server.RegisterNewAccountResponse, error) {
	return &server.RegisterNewAccountResponse{
		Status:  status,
		Message: msg,
		MsgKey:  msg,
	}, errors.New(msg)
}

func RaiseSuccess() (*server.RegisterNewAccountResponse, error) {
	return &server.RegisterNewAccountResponse{
		Status: 200,
		MsgKey: "OK",
	}, nil
}
