/*
 * @Author: lihuan
 * @Date: 2021-12-19 15:43:06
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-23 21:16:46
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

// 获取轮播图
func (p *ProductDao) GetBanner() ([]*models.Banner, error) {
	var banner []*models.Banner

	err := p.db.Order("`order` DESC").Find(&banner).Error

	return banner, err
}

// 获取商品
func (p *ProductDao) GetProduct(id int64) (*models.Product, error) {

	var product *models.Product
	err := p.db.First(&product, id).Error
	return product, err
}

// 获取商品列表
func (p *ProductDao) GetProductList() ([]*models.Product, error) {

	var product []*models.Product
	err := p.db.Find(&product).Error
	return product, err
}

// 新增商品
func (p *ProductDao) IncrementProduct(product *models.Product) error {
	err := p.db.Create(&product).Error
	return err
}

// 更新商品
func (p *ProductDao) UpdateProduct(product *models.Product) error {

	err := p.db.Model(&models.Product{}).Where("id = ? AND deleted = ?", product.ID, 0).Updates(&product).Error
	return err
}

//删除商品
func (p *ProductDao) DeleteProduct(id int64) error {
	err := p.db.Delete(&models.Product{}, id).Error

	return err
}
