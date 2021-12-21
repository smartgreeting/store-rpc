/*
 * @Author: lihuan
 * @Date: 2021-12-21 19:52:35
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-21 21:22:21
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"github.com/smartgreeting/store-rpc/product/internal/svc"
	"github.com/smartgreeting/store-rpc/product/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取产品列表
func (l *GetProductListLogic) GetProductList(in *product.GetProductListReq) (*product.ProductListReply, error) {
	res, err := l.svcCtx.ProductDao.GetProductList()
	switch err {
	case nil:
		if len(res) == 0 {
			return nil, errors.New("没有查找到记录")
		}
		list := make([]*product.ProductReply, 0)
		if len(res) > 0 {
			for _, v := range res {
				list = append(list, ProductToMapListResp(v))
			}
		}
		return &product.ProductListReply{
			List: list,
		}, nil
	default:
		return nil, err
	}

}
