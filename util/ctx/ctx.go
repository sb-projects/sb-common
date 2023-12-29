package ctx

import (
	"context"
	"log/slog"
)

type (
	ContextKey string
)

const (
	RequestIDStr string     = "x-request-id"
	RequestID    ContextKey = ContextKey(RequestIDStr)

	TransactionIDStr string     = "transaction-id"
	TransactionID    ContextKey = ContextKey(TransactionIDStr)
)

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestID, requestID)
}

func SetTransactionID(ctx context.Context, transactionID string) context.Context {
	return context.WithValue(ctx, TransactionID, transactionID)
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(RequestID).(string)
}

func GetTransactionID(ctx context.Context) string {
	return ctx.Value(TransactionID).(string)
}

func LogCtxFields(ctx context.Context, _ ...ContextKey) []any {
	fields := []any{
		slog.String(RequestIDStr, GetRequestID(ctx)),
		slog.String(TransactionIDStr, GetTransactionID(ctx)),
	}

	// TODO: add fields passed explicitly

	return fields
}
