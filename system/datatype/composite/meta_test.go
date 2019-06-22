package composite

import (
	"monster/system/datatype"
	"monster/system/implement/base"
	"testing"
)

func TestExtends(t *testing.T) {
	person := base.NewBaseObject("person")
	nameField := base.NewBaseField("name", datatype.StringType)
	person.AddField(nameField)
	personType := NewMetaType(person)

	jim := base.NewBaseInstance(person)
	jim.SetFieldValue(nameField, "Ann")
	if personType.TypeCheck(jim) {
		t.Logf("%s is type of %s", jim.FieldValue(nameField), personType.Definition().Name())
	} else {
		t.FailNow()
	}
}
