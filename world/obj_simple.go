package world

import "github.com/Pallinder/go-randomdata"

// ObjectSimple is the simplest object in the world
// Implements the Object interface
type ObjectSimple struct {
	name string
}

func NewObjectSimple() *ObjectSimple {
	return &ObjectSimple{
		name: randomdata.SillyName(),
	}
}

func (o *ObjectSimple) Name() string {
	return o.name
}
