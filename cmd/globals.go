package main

import (
	"github.com/DanTulovsky/evolve/world"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"time"
)

var (
	primaryMonitor *pixelgl.Monitor
	window         *pixelgl.Window

	title                     string = "evolve"
	screenWidth, screenHeight int    = 1024, 768

	frameTick *time.Ticker
	fps       float64

	worldCanvas *pixelgl.Canvas

	// Our 'camera' targets (0,0) which will be the center of the screen.
	camPos = pixel.ZV
	cam    pixel.Matrix

	frames    uint64
	fpsText   *text.Text
	worldText *text.Text

	// the world itself
	w *world.World
)
