/*
 * @Author: lihuan
 * @Date: 2021-12-14 21:11:59
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-16 20:16:42
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

func (u UserDao) FindByPhone(phone string) (*models.User, error) {
	var user models.User

	err := u.db.Where("phone = ? AND deleted = ?", phone, 0).First(&user).Error

	return &user, err
}

func (u UserDao) FindUserInfoById(id uint64) (*models.User, error) {
	var user models.User

	err := u.db.Where("id = ? AND deleted = ?", id, 0).First(&user).Error

	return &user, err
}

func (u UserDao) UpdateUserInfo(user *models.User) error {

	err := u.db.Model(&models.User{}).Where("id = ? AND deleted = ?", user.ID, 0).Updates(&user).Error

	return err
}
