package main

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config/env"
	"kratos-use/internal/conf"
	"os"

	pkgLog "kratos-use/pkg/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
	_ "kratos-use/ent/runtime"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	//flagconf string

	id, _ = os.Hostname()
)

func init() {
	//flag.StringVar(&flagconf, "conf", "../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			env.NewSource("KRATOS_"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	setConfig(c, &bc)

	// 初始化日志信息
	logger := pkgLog.NewZapLog(pkgLog.Config{
		FilePath:   bc.Log.FilePath,
		MaxSize:    int(bc.Log.MaxSize),
		MaxAge:     int(bc.Log.MaxAge),
		OutputType: pkgLog.OutputType(bc.Log.OutType),
	})
	logger = log.With(logger, "version", Version)

	app, cleanup, err := wireApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

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
		Source: mustGetString(c, "DATABASE_SOURCE"),
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
			AccessSecret: mustGetString(c, "AUTH_ACCESS_SECRET"),
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
