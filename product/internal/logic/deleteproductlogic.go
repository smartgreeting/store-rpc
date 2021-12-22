/*
 * @Author: lihuan
 * @Date: 2021-12-22 20:41:17
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 21:26:25
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/internal/svc"
	"github.com/smartgreeting/store-rpc/product/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  删除商品
func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductReq) (*product.DeleteProductReply, error) {
	err := l.svcCtx.ProductDao.DeleteProduct(in.Id)
	if err != nil {
		return nil, err
	}
	return &product.DeleteProductReply{}, nil
}
