/*
 * @Author: lihuan
 * @Date: 2022-01-11 20:43:35
 * @LastEditors: lihuan
 * @LastEditTime: 2022-01-11 20:45:54
 * @Email: 17719495105@163.com
 */
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis cache.ClusterConf
	Mysql      Mysql
	Mode       string
}

type Mysql struct {
	Dns string
}
