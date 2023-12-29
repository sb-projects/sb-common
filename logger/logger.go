package logger

import (
	"context"
	"log/slog"
	"os"

	ctxUtil "github.com/sb-projects/sb-common/util/ctx"
)

type (
	Logger interface {
		Debug(context.Context, string, ...any)
		Info(context.Context, string, ...any)
		Warn(context.Context, string, ...any)
		Error(context.Context, string, ...any)

		// Trace(context.Context, string, ...any)
		// Crit(context.Context, string, ...any)
	}
	logger struct {
		inner *slog.Logger
	}
)

// NewLogger returns a logger with JSON handler and debug level
func NewLogger() Logger {
	return &logger{
		inner: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
			// AddSource: true,
		})),
	}
}

func (l *logger) Debug(ctx context.Context, msg string, args ...any) {
	args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Debug(msg, args...)
}

func (l *logger) Info(ctx context.Context, msg string, args ...any) {
	args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Info(msg, args...)
}

func (l *logger) Warn(ctx context.Context, msg string, args ...any) {
	args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Warn(msg, args...)
}

func (l *logger) Error(ctx context.Context, msg string, args ...any) {
	args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Error(msg, args...)
}
