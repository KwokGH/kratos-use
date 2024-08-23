package log

import (
	kratos_zap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLog(cfg Config) log.Logger {
	var writeSyncer zapcore.WriteSyncer
	var encoder zapcore.Encoder

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		TimeKey:        "T",
		LevelKey:       "L",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	if cfg.OutputType == OutputTypeFile {
		hook := NewHook(cfg)
		writeSyncer = zapcore.AddSync(hook)
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)

	zlogger := zap.New(core).WithOptions()
	logger := kratos_zap.NewLogger(zlogger)
	//l = log.With(l, "tid", TraceID(), "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	return log.With(logger, "tid", TraceID(), "caller", log.DefaultCaller)
}
