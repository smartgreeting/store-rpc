/*
 * @Author: lihuan
 * @Date: 2021-12-20 21:05:42
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-21 21:22:17
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"github.com/smartgreeting/store-rpc/product/internal/svc"
	"github.com/smartgreeting/store-rpc/product/models"
	"github.com/smartgreeting/store-rpc/product/product"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func ProductToMapListResp(res *models.Product) *product.ProductReply {
	return &product.ProductReply{
		Id:        res.ID,
		DetailId:  res.DetailId,
		CommentId: res.CommentId,
		Url:       res.Url,
		Des:       res.Des,
		Name:      res.Name,
		ShortName: res.ShortName,
		Type:      res.Type,
		Price:     res.Price,
		Sale:      res.Sale,
		Inventory: res.Inventory,
		Score:     res.Score,
		Discount:  res.Discount,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}
func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取产品列表
func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.ProductReply, error) {
	res, err := l.svcCtx.ProductDao.GetProduct(in.Id)

	switch err {
	case nil:
		return ProductToMapListResp(res), nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("未匹配到记录")
	default:
		return nil, err
	}

}
