/*
 * @Author: lihuan
 * @Date: 2022-01-11 20:43:35
 * @LastEditors: lihuan
 * @LastEditTime: 2022-01-11 21:05:06
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"

	"github.com/smartgreeting/store-rpc/seckill/internal/svc"
	"github.com/smartgreeting/store-rpc/seckill/seckill"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSeckillHomeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSeckillHomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeckillHomeListLogic {
	return &GetSeckillHomeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取首页秒杀列表
func (l *GetSeckillHomeListLogic) GetSeckillHomeList(in *seckill.GetSeckillHomeReq) (*seckill.GetSeckillHomeListReply, error) {
	// todo: add your logic here and delete this line
	list := make([]*seckill.GetSeckillHomeReply, 0)
	return &seckill.GetSeckillHomeListReply{
		List: list,
	}, nil
}
