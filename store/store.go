package store

import (
	"theMoon/domain"
	"theMoon/instance"
	"theMoon/meta"
	"theMoon/util"
)

// ObjectSaver can save single object for persistence
type ObjectSaver interface {
	SaveObject(object meta.Object) *util.Progress
}

// InstanceSaver can save single instance for persistence
type InstanceSaver interface {
	SaveInstance(instance instance.Instance) *util.Progress
	RemoveInstance(instance instance.Instance) *util.Progress
}

// DomainSaver can save all objects & instances in the specified domain
type DomainSaver interface {
	InstanceSaver
	ObjectSaver
	DomainInfo() domain.Domain
}
