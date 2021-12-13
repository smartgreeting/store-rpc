/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 21:14:21
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/smartgreeting/store-rpc/user/internal/svc"
	"github.com/smartgreeting/store-rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogic {
	return &GetSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取短信验证码
func (l *GetSmsLogic) GetSms(in *user.GetSmsReq) (*user.UserReply, error) {

	key := fmt.Sprintf("GetSms%s", in.Phone)
	code := genCode(6)
	err := l.svcCtx.RedCli.Set(key, code, time.Duration(l.svcCtx.Config.Sms.ExpireTime)*time.Minute).Err()
	if err != nil {
		return nil, errors.New("生成验证码失败")
	}

	return &user.UserReply{
		SmsCode: code,
	}, nil

}

func genCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()

}
