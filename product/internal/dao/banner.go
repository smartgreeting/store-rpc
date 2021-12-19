/*
 * @Author: lihuan
 * @Date: 2021-12-19 15:43:06
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 17:37:45
 * @Email: 17719495105@163.com
 */
package dao

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/models"
	"gorm.io/gorm"
)

type BannerDao struct {
	ctx context.Context
	db  *gorm.DB
}

func NewBannerDao(ctx context.Context, db *gorm.DB) *BannerDao {
	return &BannerDao{
		ctx: ctx,
		db:  db,
	}
}

func (b *BannerDao) GetBanner() ([]*models.Banner, error) {
	var banner []*models.Banner

	err := b.db.Order("`order` DESC").Find(&banner).Error

	return banner, err
}
