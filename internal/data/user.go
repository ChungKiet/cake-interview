package data

import (
	"context"
	"github.com/ChungKiet/cake-interview/ent"
	user2 "github.com/ChungKiet/cake-interview/ent/user"
	"github.com/ChungKiet/cake-interview/internal/biz"
	"github.com/ChungKiet/cake-interview/internal/constant"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.DefaultLogger),
	}
}

func (u userRepo) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return u.data.WithTx(ctx, fn)
}

func (u userRepo) CreateUser(ctx context.Context, req *biz.CreateUserRequest) (string, error) {
	userId := uuid.New().String()
	err := u.data.db.User.Create().
		SetUserID(userId).
		SetUserName(req.UserName).
		SetFullName(req.FullName).
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetPhoneNumer(req.PhoneNumber).
		SetBirthday(req.BirthDay).
		SetLatestLogin(time.Now().UTC()).
		SetAccount(req.Account).
		Exec(ctx)

	if err != nil {
		u.log.Error("CreateUser failed", "err", err)
		return constant.EMPTY_STRING, err
	}

	return userId, nil
}

func (u userRepo) GetUser(ctx context.Context, req *biz.GetUserRequest) (*ent.User, error) {
	user, err := u.data.db.User.Query().
		Where(user2.AccountEQ(req.Account)).
		Only(ctx)

	if err != nil {
		u.log.Error("GetUser failed", "err", err)
		return nil, err
	}

	return user, nil
}

func (u userRepo) UpdateLoginTime(ctx context.Context, account string) error {
	//TODO implement me
	err := u.data.db.User.Update().
		Where(user2.AccountEQ(account)).
		SetLatestLogin(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		Exec(ctx)

	if err != nil {
		u.log.Error("UpdateLoginTime failed", "err", err)
		return err
	}

	return nil
}
