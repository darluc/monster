package executor

import (
	"context"
	"monster/meta"
)

type Executor interface {
	Open() (err error)
	Next(ctx context.Context, instances *meta.Batch) (err error)
	Close() (err error)
}
