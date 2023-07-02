package pkg

import (
	"context"
)

var _ JobHandler = (*JobHandlerFunc)(nil)

type JobHandler interface {
	Handle(context.Context, string)
}

type JobHandlerFunc func(context.Context, string)

func (j JobHandlerFunc) Handle(ctx context.Context, key string) {
	j(ctx, key)
}
