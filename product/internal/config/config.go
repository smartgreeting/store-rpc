/*
 * @Author: lihuan
 * @Date: 2021-12-18 12:48:16
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 16:52:23
 * @Email: 17719495105@163.com
 */
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
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
