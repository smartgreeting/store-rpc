/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 21:15:12
 * @Email: 17719495105@163.com
 */
package svc

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/smartgreeting/store-rpc/user/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RedCli *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// gorm
	db, err := gorm.Open(mysql.Open(c.Mysql.Dns), &gorm.Config{})
	// redis
	redisDb := redis.NewClient(&redis.Options{

		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: "",
		DB:       c.Redis.DB,
	})

	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		RedCli: redisDb,
	}
}
