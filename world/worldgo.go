// world defines the entire world
package world

import (
	"fmt"
	"strings"
)

type WorldConfig struct {
	Name string
}

// World is our world
type World struct {
	name string
}

func NewWorld(wc WorldConfig) (*World, error) {
	return &World{
		name: wc.Name,
	}, nil
}

func (w *World) String() string {
	var b *strings.Builder = new(strings.Builder)

	fmt.Fprintf(b, "Name: %s", w.name)

	return b.String()
}

// Update runs every frame
// dt is the duration of the last frame
func (w *World) Update(dt float64) {

}
