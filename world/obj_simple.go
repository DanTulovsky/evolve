package world

import (
	"github.com/Pallinder/go-randomdata"
	"log"
)

// ObjectSimple is the simplest object in the world
// Implements the Object interface
type ObjectSimple struct {
	birthTick int
	deathTick int

	maxAge int

	// Updated every tick, this is the age of the world
	nowTick int

	name string
}

func newObjectSimple(nowTick int) Object {
	name := randomdata.SillyName()
	log.Printf("%s is born", name)

	return &ObjectSimple{
		name:      name,
		birthTick: nowTick,
		deathTick: 0,
		maxAge:    67,
		nowTick:   nowTick,
	}
}

func (o *ObjectSimple) Age() int {
	if o.IsDead() {
		return o.deathTick - o.birthTick
	}

	return o.nowTick - o.birthTick
}

func (o *ObjectSimple) IsAlive() bool {
	return !o.IsDead()
}

func (o *ObjectSimple) IsDead() bool {
	return o.deathTick != 0
}

func (o *ObjectSimple) Name() string {
	return o.name
}

// shouldDie returns true if the object should die this turn
func (o *ObjectSimple) shouldDie() bool {
	// Die if reached max age
	if o.IsAlive() && o.Age() >= o.maxAge {
		log.Printf("[%s] dies of old age", o.name)
		return true
	}

	return false
}

func (o *ObjectSimple) Type() string {
	return ObjectTypeToName[ObjectSimpleType]
}

// Update is run every tick
func (o *ObjectSimple) Update(nowTick int) {
	o.nowTick = nowTick

	// Check if time to die
	if o.shouldDie() {
		o.deathTick = nowTick
		return
	}

	//
}
