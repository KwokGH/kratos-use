package log

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
)

func NewKratosLog(cfg Config) log.Logger {
	l := log.DefaultLogger
	if cfg.OutputType == OutputTypeFile {
		hook := NewHook(cfg)
		l = log.NewStdLogger(hook)
	}

	l = log.With(l, "tid", TraceID(), "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	return l
}

func TraceID() log.Valuer {
	return func(ctx context.Context) interface{} {
		return ctx.Value("tid")
	}
}

func NewHook(cfg Config) *lumberjack.Logger {
	hook := &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: 10,
		MaxAge:     cfg.MaxAge,
		Compress:   false,
	}

	return hook
}
