package log

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

func TestLog(t *testing.T) {
	l := log.DefaultLogger
	l = log.With(l, "tid", tracing.TraceID(), "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
	h := log.NewHelper(l)

	h.Info("error")
}
