// Package world defines the entire world
package world

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Config struct {
	Name         string
	TickInterval time.Duration
}

func (wc Config) String() string {
	var b = new(strings.Builder)
	fmt.Fprintln(b, "Config:")
	fmt.Fprintf(b, "  Name: %v\n", wc.Name)
	fmt.Fprintf(b, "  TickInterval: %v\n", wc.TickInterval)

	return b.String()
}

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
	fmt.Fprintf(b, "Population: %d\n", w.NumObjects())
	fmt.Fprintf(b, "Age: %d\n", w.age)

	fmt.Fprintln(b, "Objects:")
	for _, o := range w.objects {
		fmt.Fprintf(b, "  [%s] %s (%d)\n", o.Type(), o.Name(), o.Age(w.age))
	}

	return b.String()
}

// NumObjects returns the number of objects in the world
func (w *World) NumObjects() int {
	return len(w.objects)
}

// SpawnObject spawns a new object
func (w *World) SpawnObject() {

}

// RandomEvents runs random events
func (w *World) RandomEvents() {

}

func (w *World) Run(status chan string, done chan bool) {
	ticker := time.NewTicker(w.config.TickInterval)

	for {
		select {
		case <-done:
			log.Println("World ending...")
			ticker.Stop()
		case <-ticker.C:
			w.Update()
		default:
			select {
			case status <- w.String():
			default:
			}
		}
	}
}

// Update runs every w.config.TickInterval
func (w *World) Update() {
	log.Println("tick")
	w.age++
}
