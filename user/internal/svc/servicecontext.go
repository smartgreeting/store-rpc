/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-16 22:28:21
 * @Email: 17719495105@163.com
 */
package svc

import (
	"context"

	"github.com/smartgreeting/store-rpc/user/internal/config"
	"github.com/smartgreeting/store-rpc/user/internal/dao"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RedisDB *redis.Redis
	UserDao *dao.UserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	// gorm
	db, err := gorm.Open(mysql.Open(c.Mysql.Dns), &gorm.Config{})
	// redis
	redisDB := redis.New(c.CacheRedis[0].Host)

	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:  c,
		DB:      db,
		RedisDB: redisDB,
		UserDao: dao.NewUserDao(context.TODO(), db),
	}
}
