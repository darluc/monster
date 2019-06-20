package composite

import (
	"monster/meta"
	"testing"
)

func TestGenesisType(t *testing.T) {
	obj := new(meta.Object)
	if *obj == GenesisType.Object {
		t.Logf("we got same struct instance with [new] function")
	} else {
		t.Errorf("%p is not %p", obj, GenesisType.Object)
	}
}
