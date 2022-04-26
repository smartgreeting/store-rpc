/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 16:54:36
 * @Email: 17719495105@163.com
 */
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      Mysql
	CacheRedis cache.ClusterConf
	Sms        Sms
	Mode       string
}
type Mysql struct {
	Dns string
}
type Sms struct {
	ExpireTime int
}
