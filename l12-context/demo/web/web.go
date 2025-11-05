package web

import (
	"context"
	"demo/db"
)

type CtxKey string

const (
	RequestID CtxKey = "request_id"
)

func WithRequestID(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, RequestID, "456")

	// the exported RequestID in db package is not secure. below statement
	// modifies it. Change db package not to export ctx key RequestID
	ctx = context.WithValue(ctx, db.RequestID, "???")

	return ctx
}

