/*
 * @Author: lihuan
 * @Date: 2021-12-14 20:56:34
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-16 22:32:12
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/models"
	"github.com/smartgreeting/store-rpc/user/user"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  修改用户信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoReq) (*user.UserReply, error) {
	err := l.svcCtx.UserDao.UpdateUserInfo(&models.User{
		ID:        in.Id,
		Username:  in.Username,
		Avatar:    in.Avatar,
		Gender:    in.Gender,
		Phone:     in.Phone,
		Email:     in.Email,
		Address:   in.Address,
		Hobbies:   in.Hobbies,
		UpdatedAt: in.UpdatedAt,
	})

	switch err {
	case nil:
		return &user.UserReply{}, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("未匹配到记录")
	default:
		return nil, errors.New("更新用户信息失败")
	}

}
