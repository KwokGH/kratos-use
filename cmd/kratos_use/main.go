package main

import (
	"flag"
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
			env.NewSource(""),
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
