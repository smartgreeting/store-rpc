package logic

import (
	"context"

	"github.com/smartgreeting/store-rpc/seckill/internal/svc"
	"github.com/smartgreeting/store-rpc/seckill/seckill"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSeckillListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSeckillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeckillListLogic {
	return &GetSeckillListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取秒杀列表
func (l *GetSeckillListLogic) GetSeckillList(in *seckill.GetSeckillReq) (*seckill.GetSeckillListReply, error) {
	// todo: add your logic here and delete this line

	return &seckill.GetSeckillListReply{}, nil
}
