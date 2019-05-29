package timed

import (
	"fmt"
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
	"theMoon/meta"
	"theMoon/system/datatype"
	"theMoon/system/implement/base"
)

const (
	defaultStartDate datatype.EffectiveDate = 18000101 //@todo: may use default value property of field
	defaultEndDate   datatype.EffectiveDate = 99991231
)

type Instance struct {
	*base.Instance

	timeline *dll.List // list sorted by start date
	entries  *dll.List // list sorted by entry time
}

func (ins *Instance) Update(data map[string]interface{}) {

}

func (ins *Instance) Correct(data map[string]interface{}) {

}

func (ins *Instance) Remove(startDate datatype.EffectiveDate, entryTime uint) {

}

func (ins *Instance) HistoricalUpdate() {

}

func (ins *Instance) HistoricalCorrect() {

}

func (ins *Instance) HistoricalRemove() {

}

func (ins *Instance) SliceAt(date datatype.EffectiveDate, entryTime uint) (slice *Instance) {
	iter := ins.timeline.Iterator()
	for iter.Next() {
		snapshot, ok := iter.Value().(*Instance)
		if !ok {
			panic("snapshot is not a timed instance")
		}
		startDate := snapshot.FieldValueByName(FieldEffectiveStartDate).(datatype.EffectiveDate)
		endDate := snapshot.FieldValueByName(FieldEffectiveEndDate).(datatype.EffectiveDate)
		if startDate <= date && endDate >= date {
			if startDate == endDate && entryTime > 0 { // search in entries
				entryIter := snapshot.entries.Iterator()
				for entryIter.Next() {
					slice = entryIter.Value().(*Instance)
					if slice.FieldValueByName(FieldEffectiveEntryTime) == entryTime {
						return slice
					}
				}
				return nil
			}
			slice = snapshot
			break
		}
	}
	return slice
}

func NewTimedInstance(obj meta.Object) meta.Instance {
	if _, ok := obj.(Object); ok {
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
		timeline := dll.New()
		ins := &Instance{timeline: timeline, Instance: baseInstance.(*base.Instance), entries: dll.New()}
		timeline.Add(ins) // add self into snapshot list
		return ins
	}
	// will panic if meta object is not timeline supported
	panic(fmt.Errorf("meta object is not a timed object"))
}
