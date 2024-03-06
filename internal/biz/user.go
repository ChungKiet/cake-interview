package biz

import (
	"context"
	"encoding/json"
	"github.com/ChungKiet/cake-interview/adapter"
	"github.com/ChungKiet/cake-interview/api/server"
	"github.com/ChungKiet/cake-interview/internal/constant"
	"github.com/ChungKiet/cake-interview/internal/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
	"time"
)

type UserUseCase struct {
	repo    UserRepo
	natsCli nats.JetStreamContext // stream context
	topic   string
	log     *log.Helper
}

func NewUserUserCase(repo UserRepo, natsCli nats.JetStreamContext, topic string) *UserUseCase {
	return &UserUseCase{
		repo:    repo,
		natsCli: natsCli,
		topic:   topic,
		log:     log.NewHelper(log.DefaultLogger),
	}
}

func (uc *UserUseCase) Register(ctx context.Context, req *server.RegisterNewAccountRequest) error {
	// find and
	if req.Account == constant.EMPTY_STRING {
		return constant.ErrUserAccountRequired
	}

	// user_name, phone_number & email can't both empty
	if req.UserName == constant.EMPTY_STRING &&
		req.PhoneNumber == constant.EMPTY_STRING &&
		req.Email == constant.EMPTY_STRING {
		return constant.ErrUserInfoRequired
	}

	// TODO: valid email format
	// TODO: valid phone number

	if req.Account != req.UserName &&
		req.Account != req.PhoneNumber &&
		req.Account != req.Email {
		return constant.ErrUserAccountInValid
	}

	// TODO: valid strong password
	if len(req.Password) < constant.MIN_PASSWORD_LENGTH {
		return constant.ErrUserPasswordInvalid
	}

	// valid other info
	if req.FullName == constant.EMPTY_STRING {
		return constant.ErrUserFullNameRequired
	}

	if req.Birthday <= 0 {
		return constant.ErrUserBirthdayInvalid
	}

	birthdayConverted := time.Unix(req.Birthday, 0)

	// check user already exists and insert in tx
	//var userId string
	userId, err := uc.repo.CreateUser(ctx, &CreateUserRequest{
		FullName:    req.FullName,
		UserName:    req.UserName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Account:     req.Account,
		Password:    utils.HashString(req.Password),
		BirthDay:    birthdayConverted,
	})

	if err != nil {
		uc.log.Error("[Register] create new account failed", "err", err.Error())
		return err
	}

	// Publish msg to create user_voucher
	data, _ := json.Marshal(&adapter.UserLoginEventData{
		UserID: userId,
	})

	_, err = uc.natsCli.Publish(uc.topic, data)
	if err != nil {
		uc.log.Error("[Register] Publish event user register failed", "err", err)
	}

	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, req *server.LoginRequest) error {
	if req.UserName == constant.EMPTY_STRING {
		return constant.ErrUserAccountRequired
	}

	if req.Password == constant.EMPTY_STRING {
		return constant.ErrUserPasswordInvalid
	}

	// query and update user latest login
	errTx := uc.repo.WithTx(ctx, func(ctx context.Context) error {
		user, err := uc.repo.GetUser(ctx, &GetUserRequest{Account: req.UserName})
		if err != nil {
			return err
		}

		if user.Password != utils.HashString(req.Password) {
			return constant.ErrUserPasswordIncorrect
		}

		err = uc.repo.UpdateLoginTime(ctx, req.UserName)

		return err
	})

	if errTx != nil {
		uc.log.Error("[Login] Login failed", "err", errTx.Error())
		return errTx
	}

	return nil
}
