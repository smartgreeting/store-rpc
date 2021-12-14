// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/smartgreeting/store-rpc/user/internal/logic"
	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

//  获取短信验证码
func (s *UserServer) GetSms(ctx context.Context, in *user.GetSmsReq) (*user.UserReply, error) {
	l := logic.NewGetSmsLogic(ctx, s.svcCtx)
	return l.GetSms(in)
}

//  用户注册
func (s *UserServer) Register(ctx context.Context, in *user.RegisterReq) (*user.UserReply, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

//  用户登陆
func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.UserReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

//  获取用户信息
func (s *UserServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.UserReply, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

//  修改用户信息
func (s *UserServer) UpdateUserInfo(ctx context.Context, in *user.UpdateUserInfoReq) (*user.UserReply, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}
