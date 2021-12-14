// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"github.com/smartgreeting/store-rpc/user/user"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetSmsReq         = user.GetSmsReq
	GetUserInfoReq    = user.GetUserInfoReq
	LoginReq          = user.LoginReq
	RegisterReq       = user.RegisterReq
	UpdateUserInfoReq = user.UpdateUserInfoReq
	UserReply         = user.UserReply

	User interface {
		//  获取短信验证码
		GetSms(ctx context.Context, in *GetSmsReq, opts ...grpc.CallOption) (*UserReply, error)
		//  用户注册
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserReply, error)
		//  用户登陆
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserReply, error)
		//  获取用户信息
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*UserReply, error)
		//  修改用户信息
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UserReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

//  获取短信验证码
func (m *defaultUser) GetSms(ctx context.Context, in *GetSmsReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetSms(ctx, in, opts...)
}

//  用户注册
func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

//  用户登陆
func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

//  获取用户信息
func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

//  修改用户信息
func (m *defaultUser) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}
