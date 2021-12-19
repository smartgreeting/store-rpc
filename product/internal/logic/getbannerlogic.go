/*
 * @Author: lihuan
 * @Date: 2021-12-18 12:48:16
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 20:29:27
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"github.com/smartgreeting/store-rpc/product/internal/svc"
	"github.com/smartgreeting/store-rpc/product/models"
	"github.com/smartgreeting/store-rpc/product/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func BannerToBannerListResp(banner *models.Banner) *product.Banner {
	return &product.Banner{
		Id:        banner.ID,
		ProductId: banner.ProductId,
		Url:       banner.Url,
		Order:     banner.Order,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}
func NewGetBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBannerLogic {
	return &GetBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取轮播图
func (l *GetBannerLogic) GetBanner(in *product.GetBannerReq) (*product.BannerReply, error) {
	res, err := l.svcCtx.BannerDao.GetBanner()

	switch err {
	case nil:
		if len(res) == 0 {
			return nil, errors.New("没有查找到记录")
		}
		list := make([]*product.Banner, 0)
		if len(res) > 0 {
			for _, banner := range res {
				list = append(list, BannerToBannerListResp(banner))
			}
		}
		return &product.BannerReply{
			List: list,
		}, nil
	default:
		return nil, err
	}

}
