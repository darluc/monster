package domain

import (
	"theMoon/context"
	"theMoon/instance"
	"theMoon/meta"
)

// Domain is a container for objects & instances
type Domain interface {
	context.Context
	Objects() []meta.Object
	Instances() []instance.Instance
}
