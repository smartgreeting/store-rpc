/*
 * @Author: lihuan
 * @Date: 2021-12-14 20:56:34
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-15 21:31:42
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"github.com/smartgreeting/store-rpc/user/internal/dao"
	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/models"
	"github.com/smartgreeting/store-rpc/user/user"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.DB),
	}
}

//  获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.UserReply, error) {
	res, err := l.userDao.FindUserInfoById(in.Id)

	switch err {
	case nil:
		return MapToUserInfo(res), nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("未匹配到记录")
	default:
		return nil, err
	}
}

func MapToUserInfo(u *models.User) *user.UserReply {
	return &user.UserReply{
		Id:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Email:     u.Email,
		Address:   u.Address,
		Hobbies:   u.Hobbies,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
