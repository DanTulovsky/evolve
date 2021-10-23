package main

import (
	"fmt"
	"github.com/DanTulovsky/evolve/world"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	pixelgl.Run(run)
}

func initScreen() {
	primaryMonitor = pixelgl.PrimaryMonitor()
	cfg := pixelgl.WindowConfig{
		Title:   title,
		Bounds:  pixel.R(0, 0, float64(screenWidth), float64(screenHeight)),
		Monitor: nil,
		VSync:   true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	window = win
	worldCanvas = pixelgl.NewCanvas(win.Bounds())

	// Set the camera to look at camPos.
	cam = pixel.IM.Moved(worldCanvas.Bounds().Center().Sub(camPos))
	worldCanvas.SetMatrix(cam)
}

func initText() {
	fpsText = text.New(pixel.V(10, window.Bounds().H()-20), text.Atlas7x13)
	fpsText.Color = colornames.Antiquewhite

	worldAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	worldText = text.New(pixel.V(window.Bounds().W()-200, window.Bounds().H()-20), worldAtlas)
	worldText.Color = colornames.Black
}

// setFPS allows us to set max frames per second.
// Disable any maximum by passing 0.
func setFPS(fps int) {
	if fps <= 0 {
		frameTick = nil
	} else {
		frameTick = time.NewTicker(time.Second / time.Duration(fps))
	}
}

// draw is called after update and just draws
// everything visible to the screen.
func draw() {
	window.Clear(colornames.Skyblue)
	worldCanvas.Clear(pixel.Alpha(0))

	// Draw HUD
	fpsText.Clear()
	fpsText.WriteString(fmt.Sprintf("FPS: %d", int(math.Round(fps))))
	fpsText.Draw(window, pixel.IM)

	worldText.Clear()
	worldText.WriteString(worldStatus)
	worldText.Draw(window, pixel.IM)
}

// update handles all logic changes in the game. This
// includes moving objects or handling input.
func update(dt float64) {}

func run() {
	rand.Seed(time.Now().UnixNano())

	initScreen()
	initText()
	setFPS(0)

	window.Clear(colornames.Skyblue)

	objects := make([]world.Object, 0)
	objects = append(objects, world.PossibleObjects[world.ObjectSimpleType]())

	var err error
	w, err = world.NewWorld(world.DefaultConfig, objects)
	if err != nil {
		log.Fatalf("initializing world: %v", err)
	}
	doneCh := make(chan bool)
	statusCh := make(chan string, 1)
	go w.Run(statusCh, doneCh)

	last := time.Now()
	start := time.Now()

	for !window.Closed() {
		// duration of the last frame
		dt := time.Since(last).Seconds()
		last = time.Now()

		fps = float64(frames) / last.Sub(start).Seconds()

		// get latest statusCh from the world
		select {
		case s := <-statusCh:
			worldStatus = s
		default:
		}

		// dt is passed to the update() function (3) which can then interpolate if the frames don't have a consistent duration.
		// https://gafferongames.com/post/fix_your_timestep/
		update(dt)

		// draw everything visible onto the screen
		draw()

		frames++
		window.Update()

		if frameTick != nil {
			<-frameTick.C
		}
	}

	doneCh <- true
}
