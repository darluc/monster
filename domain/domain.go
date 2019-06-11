package domain

import (
	"monster/context"
	"monster/instance"
	"monster/meta"
)

// Domain is a container for objects & instances
type Domain interface {
	context.Context
	Objects() []meta.Object
	Instances() []instance.Instance
}
