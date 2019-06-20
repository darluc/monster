package relation

import (
	"context"
	"github.com/sirupsen/logrus"
	"monster/meta"
	"monster/system/datatype/composite"
)

type DestroyExec struct {
	Relation *composite.MetaDrivenType
}

func (*DestroyExec) Open() (err error) {
	return nil
}

func (*DestroyExec) Next(ctx context.Context, instances *meta.Batch) (err error) {
	for _, ins := range instances.Items {
		// @todo: search for related domain, and notify them
		logrus.Debugf("notify instance[%s] will be destroyed", ins.MetaObject().Name())
	}
	return
}

func (*DestroyExec) Close() (err error) {
	return
}