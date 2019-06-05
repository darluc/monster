package timed

import (
	"testing"
	"theMoon/system/datatype"
	"theMoon/system/implement/base"
)

func TestTimeInstanceSlice(t *testing.T) {
	obj := NewTimedObject()
	nameField := base.NewBaseField("name", datatype.StringType)
	obj.AddField(nameField)
	ageField := NewTimedField("age", datatype.NumberType)
	obj.AddField(ageField)

	ins := NewTimedInstance(obj)
	ins.SetFieldValue(nameField, "bruce")
	ins.(*Instance).SetTimedFieldValue(ageField, 18, defaultStartDate, 0)
	someDate := 20190506
	var age int
	age = ins.(*Instance).TimedFieldValue(ageField, datatype.EffectiveDate(someDate), 0).(int)
	if age == 18 {
		t.Logf("age at %d is %d", someDate, age)
	} else {
		t.Errorf("failed: age at %d is %d", someDate, age)
	}

	ins.(*Instance).SetTimedFieldValue(ageField, 25, 18991212, 0)
	age = ins.(*Instance).TimedFieldValue(ageField, datatype.EffectiveDate(someDate), 0).(int)
	if age == 25 {
		t.Logf("age at %d is %d", someDate, age)
	} else {
		t.Errorf("failed: age at %d is %d", someDate, age)
	}

	ins.(*Instance).SetTimedFieldValue(ageField, 28, 20181212, 0)
	age = ins.(*Instance).FieldValue(ageField).(int)
	if age == 28 {
		t.Logf("current age is %d", age)
	} else {
		t.Errorf("failed: current age is %d", age)
	}
}
