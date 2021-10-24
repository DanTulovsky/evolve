// Package world defines the entire world
package world

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// World is our world
type World struct {
	name    string
	config  Config
	objects []Object
	age     int
}

// NewWorld creates a new world with the provided config and initial objects
func NewWorld(wc Config, objects []Object) (*World, error) {
	log.Println("Initializing world...")
	log.Println(wc)
	return &World{
		name:    wc.Name,
		config:  wc,
		objects: objects,
	}, nil
}

func (w *World) String() string {
	var b = new(strings.Builder)

	fmt.Fprintf(b, "Name: %s\n", w.name)
	fmt.Fprintf(b, "Population: %d (alive: %d; dead: %d)\n", w.numObjects(), w.numLiveObjects(), w.numDeadObjects())
	fmt.Fprintf(b, "Age: %d\n", w.age)

	fmt.Fprintln(b, "Live Objects:")
	for _, o := range w.liveObjects() {
		fmt.Fprintf(b, "  [%s] %s (%d)\n", o.Type(), o.Name(), o.Age())
	}

	return b.String()
}

func (w *World) liveObjects() []Object {
	var os []Object
	for _, o := range w.objects {
		if o.IsAlive() {
			os = append(os, o)
		}
	}
	return os
}

// numDeadObjects returns the number of dead objects in the world
func (w *World) numDeadObjects() int {
	return len(w.objects) - len(w.liveObjects())
}

// numObjects returns the number of objects in the world
func (w *World) numObjects() int {
	return len(w.objects)
}

// numLiveObjects returns the number of live objects in the world
func (w *World) numLiveObjects() int {
	return len(w.liveObjects())
}

// spawnObject spawns a new object
func (w *World) spawnObject(o ObjectType) {
	var n Object

	switch o {
	case ObjectSimpleType:
		n = newObjectSimple(w.age)
	}

	w.objects = append(w.objects, n)
}

// randomEvents runs random events
func (w *World) randomEvents() {
	// every 10 ticks, spawn a new simpleObject
	if w.age%10 == 0 && w.numObjects() < w.config.MaxObjects {
		log.Println("Spawning new ObjectSimple...")
		w.spawnObject(ObjectSimpleType)
	}

}

func (w *World) Run(status chan string, done chan bool) {
	ticker := time.NewTicker(w.config.TickInterval)

	for {
		select {
		case <-done:
			log.Println("World ending...")
			ticker.Stop()
		case <-ticker.C:
			w.update()
		default:
			select {
			case status <- w.String():
			default:
			}
		}
	}
}

// update runs every w.config.TickInterval
func (w *World) update() {
	w.age++

	log.Println("tick")

	// Object actions
	for _, o := range w.objects {
		o.Update(w.age)
	}

	// random world events
	w.randomEvents()

}
