package datatype

import (
	"reflect"
	"strconv"
)

var EffectiveDateType EffectiveDate

type EffectiveDate uint

func (EffectiveDate) TypeCheck(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Uint {
		if s := strconv.Itoa(value.(int)); len(s) == 8 { // 20190603 = 2019/06/03
			return true
		}
	}
	return false
}

func (EffectiveDate) ReflectType() reflect.Type {
	return reflect.TypeOf(EffectiveDateType)
}
