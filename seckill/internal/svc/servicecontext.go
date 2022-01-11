/*
 * @Author: lihuan
 * @Date: 2022-01-11 20:43:35
 * @LastEditors: lihuan
 * @LastEditTime: 2022-01-11 20:50:41
 * @Email: 17719495105@163.com
 */
package svc

import (
	"context"

	"github.com/smartgreeting/store-rpc/seckill/internal/config"
	"github.com/smartgreeting/store-rpc/seckill/internal/dao"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	RedisDB    *redis.Redis
	SeckillDao *dao.SeckillDao
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
		Config:     c,
		DB:         db,
		RedisDB:    redisDB,
		SeckillDao: dao.NewSeckillDao(context.TODO(), db),
	}

}
