package executor

import (
	"context"
	"theMoon/meta"
)

type Executor interface {
	Open() error
	Next(ctx context.Context, instances *meta.Batch) error
	Close() error
}
