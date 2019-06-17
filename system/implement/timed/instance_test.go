package timed

import (
	"monster/system/datatype"
	"monster/system/implement/base"
	"testing"
)

func TestTimeInstanceSlice(t *testing.T) {
	obj := NewTimedObject()
	nameField := base.NewBaseField("name", datatype.StringType)
	obj.AddField(nameField)
	ageField := NewTimedField("age", datatype.NumberType)
	obj.AddField(ageField)

	ins := NewTimedInstance(obj)
	timedIns := ins.(*Instance)
	ins.SetFieldValue(nameField, "bruce")
	// set default start value
	timedIns.SetTimedFieldValue(ageField, 18, defaultStartDate, 0)
	someDate := 20190506
	var age int
	age = ins.(*Instance).TimedFieldValue(ageField, datatype.EffectiveDate(someDate), 0).(int)
	if age == 18 {
		t.Logf("age at %d is %d", someDate, age)
	} else {
		t.Errorf("failed: age at %d is %d", someDate, age)
	}

	// set value at 1899/12/12
	timedIns.SetTimedFieldValue(ageField, 25, 18991212, 0)
	age = timedIns.TimedFieldValue(ageField, datatype.EffectiveDate(someDate), 0).(int)
	if age == 25 {
		t.Logf("age at %d is %d", someDate, age)
	} else {
		t.Errorf("failed: age at %d is %d", someDate, age)
	}

	// set value at 2018/12/12
	timedIns.SetTimedFieldValue(ageField, 28, 20181212, 0)
	timedIns.SetTimedFieldValue(ageField, 29, 20181212, 999)
	age = timedIns.FieldValue(ageField).(int)
	if age == 29 {
		t.Logf("current age is %d", age)
	} else {
		t.Errorf("failed: current age is %d", age)
	}

	// test cut points for instances
	pointsCount := timedIns.Points().Size()
	if pointsCount == 3 {
		t.Logf("instance got 3 cut points")
	} else {
		t.Errorf("instance cut points is %d", pointsCount)
	}
}
