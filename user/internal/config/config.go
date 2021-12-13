/*
 * @Author: lihuan
 * @Date: 2021-12-13 21:10:21
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 21:13:31
 * @Email: 17719495105@163.com
 */
package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql Mysql
	Redis Redis
	Sms   Sms
}
type Mysql struct {
	Dns string
}

type Redis struct {
	Host string
	Port int
	DB   int
}

type Sms struct {
	ExpireTime int
}
