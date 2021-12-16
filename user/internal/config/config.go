/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-16 22:17:35
 * @Email: 17719495105@163.com
 */
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      Mysql
	CacheRedis cache.ClusterConf
	Sms        Sms
}
type Mysql struct {
	Dns string
}
type Sms struct {
	ExpireTime int
}
