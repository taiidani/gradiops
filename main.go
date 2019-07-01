package main

import (
	"flag"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Flags
var debug *bool
var fps *float64
var displayFps *bool

// Global state
var game *tl.Game
var score int

func init() {
	// Seed the faker library
	rand.Seed(time.Now().Unix())
}

func main() {
	debug = flag.Bool("debug", false, "Debug mode")
	fps = flag.Float64("fps", 30.0, "Set frames per second")
	displayFps = flag.Bool("fps-display", false, "Display frames per second on screen")
	flag.Parse()

	// Start the Game
	setupGame()
	newGame()

	// Let's A Go
	game.Start()
}

func setupGame() {
	game = tl.NewGame()
	game.SetDebugOn(*debug)
}

func newGame() {
	score = 0
	game.Log("Starting new game")

	if *displayFps {
		game.Screen().AddEntity(tl.NewFpsText(60, 0, tl.ColorWhite, tl.ColorBlack, 1))
	}
	game.Screen().SetFps(*fps)

	// Construct the level
	level := newBaseLevel()
	game.Screen().AddEntity(newHUD(level))
	game.Screen().SetLevel(level)
}
