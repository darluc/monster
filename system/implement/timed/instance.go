package timed

import (
	"fmt"
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/sets/treeset"
	"monster/meta"
	"monster/system/datatype"
	"monster/system/implement/base"
	"monster/system/property"
	"monster/system/util"
)

const (
	defaultStartDate datatype.EffectiveDate = 18000101 //@todo: may use default value property of field
	defaultEndDate   datatype.EffectiveDate = 99991231
)

type Instance struct {
	*base.Instance

	points      *treeset.Set             // hashSet saved time point
	timedFields map[meta.Field]*sll.List // field : SinglyLinkedList<TimedValue>，value is ordered by date asc
}

// Points returns cut points of timed instance
func (ins *Instance) Points() *treeset.Set {
	return ins.points
}

func (ins *Instance) FieldValue(fld meta.Field) (val interface{}) {
	if fld.PropertyValue(property.TimelineEnabled) == true {
		val = ins.TimedFieldValue(fld, util.GetCurrentDate(), 0)
	} else {
		val = ins.Instance.FieldValue(fld)
	}
	return
}

// TimedFieldValue reads out value of timed field with the effective date and specified entry time.
// If entryTime is zero, just use the first one has the same effective date.
func (ins *Instance) TimedFieldValue(fld meta.Field, asOfDate datatype.EffectiveDate, entryTime uint) (val interface{}) {
	if fld.PropertyValue(property.TimelineEnabled) != true {
		panic("cannot read common field with TimedFieldValue method")
	}

	if list, ok := ins.timedFields[fld]; ok {
		iter := list.Iterator()
		for iter.Next() {
			current := iter.Value().(*timedValue)
			if current.Date <= asOfDate && (current.EntryTime == entryTime || entryTime == 0) {
				val = current.Value
			} else if current.Date > asOfDate {
				break
			}
		}
	}
	return
}

func (ins *Instance) SetFieldValue(fld meta.Field, val interface{}) {
	if fld.PropertyValue(property.TimelineEnabled) == true {
		panic("timed field cannot be updated by Instance.SetFieldValue")
	} else {
		ins.Instance.SetFieldValue(fld, val)
	}
}

// SetTimedFieldValue use to set value for timed field, and maintain the list of values
// which are ordered by [date asc, entry time desc]
func (ins *Instance) SetTimedFieldValue(fld meta.Field, val interface{}, startDate datatype.EffectiveDate, entryTime uint) {
	if fld.PropertyValue(property.TimelineEnabled) != true {
		panic("common field cannot be updated by Instance.SetTimedFieldValue")
	}

	list, ok := ins.timedFields[fld]
	if !ok {
		list = sll.New()
		ins.timedFields[fld] = list
	}

	if entryTime == 0 {
		iter := list.Iterator()
		for iter.Next() {
			if iter.Value().(*timedValue).Date == startDate { // find the first value which has the same effective date
				entryTime = iter.Value().(*timedValue).EntryTime + 1 // auto increase entry time
				break
			} else if iter.Value().(*timedValue).Date < startDate {
				break
			}
		}
	}
	list.Add(&timedValue{Date: startDate, EntryTime: entryTime, Value: val})

	// sort by date & entry time
	list.Sort(func(a, b interface{}) int {
		aVal := a.(*timedValue)
		bVal := b.(*timedValue)
		if aVal.Date == bVal.Date {
			if aVal.EntryTime > bVal.EntryTime {
				return 1
			} else {
				return -1
			}
		} else if aVal.Date > bVal.Date {
			return 1
		} else {
			return -1
		}
	})

	// remember date point
	ins.points.Add(startDate)
}

// SliceAt return a slice of timed instance value of all fields
func (ins *Instance) SliceAt(asOfDate datatype.EffectiveDate) map[meta.Field]interface{} {
	ret := make(map[meta.Field]interface{})
	for _, fld := range ins.MetaObject().Fields() {
		ret[fld] = ins.FieldValue(fld)
	}
	return ret
}

func NewTimedInstance(obj meta.Object) meta.Instance {
	if _, ok := obj.(*Object); ok {
		baseInstance := base.NewBaseInstance(obj)

		// set default start date
		startDateField := obj.Field(FieldEffectiveStartDate)
		if startDateField == nil {
			panic(fmt.Errorf("effective start date field not extists"))
		}
		baseInstance.SetFieldValue(startDateField, defaultStartDate)

		// set default end date
		endDateField := obj.Field(FieldEffectiveEndDate)
		if endDateField == nil {
			panic(fmt.Errorf("effective start date field not extists"))
		}
		baseInstance.SetFieldValue(endDateField, defaultEndDate)

		// initialize timeline list
		pointsTreeSet := treeset.NewWith(func(a interface{}, b interface{}) int {
			ai := a.(datatype.EffectiveDate)
			bi := b.(datatype.EffectiveDate)
			if ai > bi {
				return 1
			} else if ai < bi {
				return -1
			} else {
				return 0
			}
		})
		ins := &Instance{Instance: baseInstance.(*base.Instance), points: pointsTreeSet, timedFields: map[meta.Field]*sll.List{}}
		return ins
	}
	// will panic if meta object is not timeline supported
	panic(fmt.Errorf("meta object is not a timed object"))
}
