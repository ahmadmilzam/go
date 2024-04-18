package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLog() Middleware {
	return func(ctx *gin.Context) {
		t1 := time.Now().UnixNano() / int64(time.Millisecond)
		ctx.Next()

		t2 := time.Now().UnixNano() / int64(time.Millisecond)
		diff := t2 - t1

		msg, msg_exist := ctx.Get("msg")
		if !msg_exist {
			msg = ""
		}
		err, err_exist := ctx.Get("err")
		if !err_exist {
			err = ""
		}

		// access the status we are sending
		status := ctx.Writer.Status()

		switch {
		case status >= 500:
			slog.Error(
				msg.(string),
				"type", "Unexpected Error",
				"error", err,
				"duration", float64(diff),
				"path", ctx.Request.URL.EscapedPath(),
				"method", ctx.Request.Method,
				"status", status)
		case status >= 400:
			slog.Warn(
				msg.(string),
				"type", "Expected Error",
				"error", err,
				"duration", float64(diff),
				"path", ctx.Request.URL.EscapedPath(),
				"method", ctx.Request.Method,
				"status", status)
		default:
			slog.Info(
				"Request processed",
				"duration", float64(diff),
				"path", ctx.Request.URL.EscapedPath(),
				"method", ctx.Request.Method,
				"status", status)
		}
	}
}
