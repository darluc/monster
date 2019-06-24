package domain

import (
	"monster/context"
	"monster/meta"
)

// Domain is a container for objects & instances
type Domain interface {
	context.Context
	Objects() []meta.Object
	Instances() []meta.Instance
}
