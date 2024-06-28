package main

import (
	"flag"

	"github.com/taiidani/gradiops/internal"
)

// Flags
var debug *bool
var fps *float64
var displayFps *bool

func main() {
	debug = flag.Bool("debug", false, "Debug mode")
	fps = flag.Float64("fps", 30.0, "Set frames per second")
	displayFps = flag.Bool("fps-display", false, "Display frames per second on screen")
	flag.Parse()

	// Start the Game
	game := internal.SetupGame(*debug, *displayFps, *fps)
	game.Restart()

	// Let's A Go
	game.Start()
}
