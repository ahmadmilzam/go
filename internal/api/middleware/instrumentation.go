package middleware

import (
	"fmt"
	"time"

	"github.com/ahmadmilzam/go/pkg/statsd"
	"github.com/gin-gonic/gin"
)

func InstrumentStatsD() Middleware {
	return func(ctx *gin.Context) {
		t1 := time.Now().UnixNano() / int64(time.Millisecond)
		ctx.Next()

		t2 := time.Now().UnixNano() / int64(time.Millisecond)
		diff := t2 - t1
		_ = statsd.Gauge(fmt.Sprintf("%s.%s", ctx.Request.Method, ctx.Request.URL.EscapedPath()), float64(diff))
	}
}
