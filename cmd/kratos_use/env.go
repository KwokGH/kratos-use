package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"kratos-use/internal/conf"
	pkgLog "kratos-use/pkg/log"
)

const (
	DatabaseSourceKey   = "KRATOS_DATABASE_SOURCE"
	AuthAccessSecretKey = "KRATOS_AUTH_ACCESS_SECRET"
)

func setConfig(c config.Config, bc *conf.Bootstrap) {
	// 服务
	bc.Server = &conf.Server{Http: &conf.Server_HTTP{
		Network: "",
		Addr:    "0.0.0.0:8000",
		Timeout: 10,
	}}

	// 数据库
	bc.Data = &conf.Data{Database: &conf.Data_Database{
		Driver: "mysql",
		Source: mustGetString(c, DatabaseSourceKey),
	}}

	// 日志
	bc.Log = &conf.Log{
		OutType:  string(pkgLog.OutputTypeConsole),
		FilePath: "",
		MaxSize:  0,
		MaxAge:   0,
	}

	// 其他一些配置
	bc.App = &conf.App{
		// jwt
		Auth: &conf.App_Auth{
			AccessSecret: mustGetString(c, AuthAccessSecretKey),
			AccessExpire: 1296000,
		},
	}
}

func mustGetString(c config.Config, key string) string {
	if val, err := c.Value(key).String(); err != nil {
		fmt.Println(key)
		panic(err)
	} else {
		return val
	}
}
func mustGetInt(c config.Config, key string) int64 {
	if val, err := c.Value(key).Int(); err != nil {
		panic(err)
	} else {
		return val
	}
}
