package actx

import "context"

type ContextKey string

const (
	EchoContext ContextKey = "EchoContext"
)

func Get(ctx context.Context, key ContextKey) interface{} {
	return ctx.Value(key)
}

func Set(ctx context.Context, key ContextKey, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}
