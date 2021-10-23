package world

// Object is an object that can populate the world
type Object interface {
	Age(nowTick int) int
	Name() string
	Type() string
	Update()
}

type ObjectType int

const (
	ObjectSimpleType ObjectType = iota
)

var (
	ObjectTypeToName = map[ObjectType]string{
		ObjectSimpleType: "si",
	}
)

// PossibleObjects is a map from object type to the creation function
var PossibleObjects = map[ObjectType]func() Object{
	ObjectSimpleType: newObjectSimple,
}
