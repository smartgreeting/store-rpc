/*
 * @Author: lihuan
 * @Date: 2022-01-11 20:48:16
 * @LastEditors: lihuan
 * @LastEditTime: 2022-01-11 20:49:30
 * @Email: 17719495105@163.com
 */
package dao

import (
	"context"

	"gorm.io/gorm"
)

type SeckillDao struct {
	ctx context.Context
	db  *gorm.DB
}

func NewSeckillDao(ctx context.Context, db *gorm.DB) *SeckillDao {
	return &SeckillDao{
		ctx: ctx,
		db:  db,
	}
}
