package world

import "github.com/Pallinder/go-randomdata"

// ObjectSimple is the simplest object in the world
// Implements the Object interface
type ObjectSimple struct {
	name      string
	birthTick int
}

func newObjectSimple() Object {
	return &ObjectSimple{
		name:      randomdata.SillyName(),
		birthTick: 0,
	}
}

func (o *ObjectSimple) Age(nowTick int) int {
	return nowTick - o.birthTick
}

func (o *ObjectSimple) Name() string {
	return o.name
}

func (o *ObjectSimple) Type() string {
	return ObjectTypeToName[ObjectSimpleType]
}

// Update is run every tick
func (o *ObjectSimple) Update() {
}
