package mysql

import (
	"context"
	"theMoon/domain"
	"theMoon/instance"
	"theMoon/meta"
	"theMoon/util"
)

type StoreExec struct {
	ds *DataSaver
}

func (*StoreExec) Open() error {
	//initialize mysql datatype saver
}

func (store *StoreExec) Next(ctx context.Context, instances *instance.Batch) error {
	for _, ins := range instances.Items {
		store.ds.SaveObject(ins.MetaObject())
		store.ds.SaveInstance(ins)
	}
	return nil
}

func (*StoreExec) Close() error {
	//@todo close mysql datatype saver close connection to mysql server
	panic("implement me")
}

type DataSaver struct {
}

func (*DataSaver) SaveInstance(instance instance.Instance) *util.Progress {
	panic("implement me")
}

func (*DataSaver) RemoveInstance(instance instance.Instance) *util.Progress {
	panic("implement me")
}

func (*DataSaver) SaveObject(object meta.Object) *util.Progress {
	panic("implement me")
}

func (*DataSaver) SaveAll(domain domain.Domain) *util.Progress {
	panic("implement me")
}

func NewDataSaver() *DataSaver {
	return &DataSaver{}
}
