package logger

import (
	"io"
	"log/slog"
	"slices"
)

// "gitlab.com/flip-id/go-core/middleware/server"

const (
	timeFormat       = "2006-01-02T15:04:05.000"
	maskedString     = "*"
	correlationIDKey = "correlation_id"
	sourceKey        = "source"
	errorKey         = "error"
)

func ParseLevel(level string) slog.Level {
	logLevel := slog.Level(8)
	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		slog.Error("failed to parse log level", errAttr(err))
	}
	return logLevel
}

func InitializeLogger(option *Option) {
	opts := &slog.HandlerOptions{
		AddSource:   false,
		Level:       ParseLevel(option.Level),
		ReplaceAttr: formatTimeAttrFunc(timeFormat),
	}
	logger := slog.New(newCustomHandler(option.Writer, opts))
	slog.SetDefault(logger)
}

func errAttr(err error) slog.Attr {
	return slog.String(errorKey, err.Error())
}

func formatTimeAttrFunc(timeFormat string) func(groups []string, a slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			a.Value = slog.StringValue(a.Value.Time().Format(timeFormat))
		}
		if isSensitiveKey(a.Key) {
			a.Value = slog.StringValue(maskedString)
		}
		return a
	}
}

func isSensitiveKey(key string) bool {
	sensitiveKeys := []string{"pin", "password", "secret", "token", "credential"}
	return slices.Contains(sensitiveKeys, key)
}

type customHandler struct {
	slog.Handler
}

func newCustomHandler(w io.Writer, opts *slog.HandlerOptions) *customHandler {
	return &customHandler{
		slog.NewJSONHandler(w, opts),
	}
}
