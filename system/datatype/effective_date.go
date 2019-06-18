package datatype

import (
	"reflect"
	"strconv"
)

var EffectiveDateType EffectiveDate

type EffectiveDate uint

func (EffectiveDate) TypeCheck(value interface{}) bool {
	if date, ok := value.(EffectiveDate); ok && date > 0 {
		return len(strconv.Itoa(int(date))) <= 8
	}
	return false
}

func (EffectiveDate) ReflectType() reflect.Type {
	return reflect.TypeOf(EffectiveDateType)
}
