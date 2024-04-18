package trace

import (
	"time"

	"github.com/ahmadmilzam/go/config"
	"github.com/getsentry/sentry-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Init() {
	initDatadog()
	// initSentry()
}

func Stop() {
	tracer.Stop()
	sentry.Flush(2 * time.Second)
}

func initDatadog() {
	cfg := config.GetDatadogConfig()
	if !cfg.Enabled {
		return
	}
	tracer.Start(
		tracer.WithEnv(cfg.Env),
		tracer.WithServiceName(cfg.Name),
		tracer.WithServiceVersion(cfg.Version),
	)
}

// func initSentry() {
// 	cfg := config.GetSentry()
// 	if !cfg.Enabled {
// 		return
// 	}
// 	err := sentry.Init(sentry.ClientOptions{
// 		Dsn: cfg.DSN,
// 	})
// 	if err != nil {
// 		logger.ErrAttr(err)
// 	}
// }
