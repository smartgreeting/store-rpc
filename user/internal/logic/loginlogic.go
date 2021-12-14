/*
 * @Author: lihuan
 * @Date: 2021-12-14 20:56:34
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-14 21:53:30
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"github.com/smartgreeting/store-rpc/user/internal/dao"
	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/user"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.DB),
	}
}

//  用户登陆
func (l *LoginLogic) Login(in *user.LoginReq) (*user.UserReply, error) {

	u, err := l.userDao.FindByPhone(in.Phone)

	switch err {
	case nil:
		if u.Password != in.Password {
			return nil, errors.New("密码错误")
		}
		return &user.UserReply{
			Id: u.ID,
		}, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("未匹配到记录")
	default:
		return nil, err
	}

}
