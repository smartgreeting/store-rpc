/*
 * @Author: lihuan
 * @Date: 2021-12-14 21:11:59
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-14 21:52:42
 * @Email: 17719495105@163.com
 */
package dao

import (
	"context"

	"github.com/smartgreeting/store-rpc/user/models"
	"gorm.io/gorm"
)

type UserDao struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserDao(ctx context.Context, db *gorm.DB) *UserDao {
	return &UserDao{
		ctx: ctx,
		db:  db,
	}
}

func (u UserDao) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(user).Error
	return user, err
}

func (d UserDao) FindByPhone(phone string) (*models.User, error) {
	var user models.User

	err := d.db.Where("phone = ?", phone).First(&user).Error

	return &user, err
}
