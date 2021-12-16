package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/models"
	"github.com/smartgreeting/store-rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  用户注册
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.UserReply, error) {
	key := fmt.Sprintf("GetSms%s", in.Phone)
	smsCode, _ := l.svcCtx.RedisDB.Get(key)
	if smsCode != in.SmsCode {
		return nil, errors.New("验证不正确")
	}
	res, err := l.svcCtx.UserDao.Create(&models.User{
		Phone:    in.Phone,
		Password: in.Password,
		Username: in.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &user.UserReply{
		Id: res.ID,
	}, nil
}
