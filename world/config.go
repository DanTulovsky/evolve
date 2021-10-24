package world

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Name         string
	MaxObjects   int
	TickInterval time.Duration
}

func (wc Config) String() string {
	var b = new(strings.Builder)
	fmt.Fprintln(b, "Config:")
	fmt.Fprintf(b, "  Name: %v\n", wc.Name)
	fmt.Fprintf(b, "  TickInterval: %v\n", wc.TickInterval)

	return b.String()
}

var DefaultConfig = Config{
	Name:         "default0",
	TickInterval: time.Second,
	MaxObjects:   10,
}
