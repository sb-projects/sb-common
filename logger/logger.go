package logger

import (
	"context"
	"log/slog"
	"os"
)

type (
	LoggerWrap interface {
		Debug(context.Context, string, ...any)
		Info(context.Context, string, ...any)
		Warn(context.Context, string, ...any)
		Error(context.Context, string, ...any)

		// Trace(context.Context, string, ...any)
		// Crit(context.Context, string, ...any)
	}
	Logger struct {
		inner *slog.Logger
	}
)

// NewLogger returns a logger with JSON handler and debug level
func NewLogger() *Logger {
	return &Logger{
		inner: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
}

func (l *Logger) Debug(_ context.Context, msg string, args ...any) {
	//args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Debug(msg, args...)
}

func (l *Logger) Info(_ context.Context, msg string, args ...any) {
	//args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Info(msg, args...)
}

func (l *Logger) Warn(_ context.Context, msg string, args ...any) {
	//args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Warn(msg, args...)
}

func (l *Logger) Error(_ context.Context, msg string, args ...any) {
	//args = append(args, ctxUtil.LogCtxFields(ctx))
	l.inner.Error(msg, args...)
}
