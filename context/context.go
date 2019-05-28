package context

type name string

type Context interface {
	Parent() Context
	Get(name name) interface{}
}
