package meta

// Relation is just a specialized object
type Relation interface {
	Object
}

const (
	relationName string = "RELATION"
)

type baseRelation struct {
	Object
}

var relationObject Object

func NewBaseRelation(suite ConstructionSuite) Relation {
	r := &baseRelation{}
	return r
}

func initRelationObject(suite ConstructionSuite) {
	if relationObject == nil {
		//obj := suite.objectConstructor()

	}
}

func (*baseRelation) Name() string {
	return relationName
}

func (*baseRelation) RelationWith(object Object) Object {
	return nil // RelationObject has no relationship with any other objects
}
