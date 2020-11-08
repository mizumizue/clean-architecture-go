package jsonp

import (
	"context"

	"clean-architecture-go/lib/actx"
)

type JsonWriter interface {
	JSON(code int, i interface{}) error
}

func withJsonWriter(ctx context.Context, fn func(jw JsonWriter) error) error {
	jw := actx.Get(ctx, actx.EchoContext).(JsonWriter)
	return fn(jw)
}
