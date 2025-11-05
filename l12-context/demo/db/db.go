package db

import (
	"context"
	"fmt"
)

type CtxKey string

const (
	// do not export the ctx key, the user package may modify that
	//RequestID CtxKey = "request_id"
	// make the ctxt key as unexported local
	requestID CtxKey = "request_id"
)

func RequestIDFrom(ctx context.Context) (string, error) {
	// get the request_id from the context
	// .(string) -> type assertion
	s, ok := ctx.Value(requestID).(string)
	if !ok {
		return "", fmt.Errorf("request_id not found in context")
	}
	return s, nil
}
	
func WithRequestID(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, requestID, "123")

	return ctx
}
