package datatype

import (
	"github.com/emirpasic/gods/sets/hashset"
	"monster/meta"
)

type InstanceHashSet struct {
	*hashset.Set
}

func (is *InstanceHashSet) Values() []meta.Instance {
	ret := make([]meta.Instance, 0, is.Set.Size())
	for _, v := range is.Set.Values() {
		ret = append(ret, v.(meta.Instance))
	}
	return ret
}

func (is *InstanceHashSet) Add(val meta.Instance) {
	is.Set.Add(val)
}

func (is *InstanceHashSet) Remove(val meta.Instance) {
	is.Set.Remove(val)
}

func NewInstanceHashSet() *InstanceHashSet {
	return &InstanceHashSet{Set: hashset.New()}
}
