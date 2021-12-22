package logic

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/internal/svc"
	"github.com/smartgreeting/store-rpc/product/models"
	"github.com/smartgreeting/store-rpc/product/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type IncrementProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrementProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrementProductLogic {
	return &IncrementProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  新增产品
func (l *IncrementProductLogic) IncrementProduct(in *product.ProductReq) (*product.IncrementProductReply, error) {
	err := l.svcCtx.ProductDao.IncrementProduct(&models.Product{
		Url:       in.Url,
		Des:       in.Des,
		Name:      in.Name,
		ShortName: in.ShortName,
		Type:      in.Type,
		Price:     in.Price,
		Inventory: in.Inventory,
		Discount:  in.Discount,
	})
	if err != nil {
		return nil, err
	}
	return &product.IncrementProductReply{}, nil
}
