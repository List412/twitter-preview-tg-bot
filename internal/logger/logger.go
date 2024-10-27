package logger

import (
	"context"
	"log/slog"
)

type CtxUUID struct {
}

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying
// handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attr, ok := ctx.Value(CtxUUID{}).(string); ok {
		r.AddAttrs(slog.String("uuid", attr))
	}

	return h.Handler.Handle(ctx, r)
}
