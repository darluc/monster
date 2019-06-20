package meta

type InstanceCollector interface {
	Values() []Instance
	Add(val Instance)
	Remove(val Instance)
	Size() int
}
