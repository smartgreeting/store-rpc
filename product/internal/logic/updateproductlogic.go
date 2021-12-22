/*
 * @Author: lihuan
 * @Date: 2021-12-22 20:41:17
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 21:25:40
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/internal/svc"
	"github.com/smartgreeting/store-rpc/product/models"
	"github.com/smartgreeting/store-rpc/product/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  修改产品
func (l *UpdateProductLogic) UpdateProduct(in *product.ProductReq) (*product.UpdateProductReply, error) {
	err := l.svcCtx.ProductDao.UpdateProduct(&models.Product{
		ID:        in.Id,
		DetailId:  in.DetailId,
		CommentId: in.CommentId,
		Url:       in.Url,
		Des:       in.Des,
		Name:      in.Name,
		ShortName: in.ShortName,
		Type:      in.Type,
		Price:     in.Price,
		Sale:      in.Sale,
		Inventory: in.Inventory,
		Score:     in.Score,
		Discount:  in.Discount,
	})

	if err != nil {
		return nil, err
	}

	return &product.UpdateProductReply{}, nil
}
