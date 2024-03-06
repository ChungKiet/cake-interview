package login

import (
	"errors"
	"github.com/ChungKiet/cake-interview/api/server"
)

func RaiseError(status int32, msg string) (*server.LoginResponse, error) {
	return &server.LoginResponse{
		Status:  status,
		Message: msg,
		MsgKey:  msg,
	}, errors.New(msg)
}

func RaiseSuccess() (*server.LoginResponse, error) {
	return &server.LoginResponse{
		Status: 200,
		MsgKey: "OK",
	}, nil
}
