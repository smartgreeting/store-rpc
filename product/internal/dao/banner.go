/*
 * @Author: lihuan
 * @Date: 2021-12-19 15:43:06
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-21 20:25:44
 * @Email: 17719495105@163.com
 */
package dao

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/models"
	"gorm.io/gorm"
)

type ProductDao struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductDao(ctx context.Context, db *gorm.DB) *ProductDao {
	return &ProductDao{
		ctx: ctx,
		db:  db,
	}
}

func (p *ProductDao) GetBanner() ([]*models.Banner, error) {
	var banner []*models.Banner

	err := p.db.Order("`order` DESC").Find(&banner).Error

	return banner, err
}

func (p *ProductDao) GetProduct(id int64) (*models.Product, error) {

	var product *models.Product
	err := p.db.First(&product, id).Error
	return product, err
}

func (p *ProductDao) GetProductList() ([]*models.Product, error) {

	var product []*models.Product
	err := p.db.Find(&product).Error
	return product, err
}
