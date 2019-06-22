package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var game *tl.Game
var score int

func init() {
	// Seed the faker library
	rand.Seed(time.Now().Unix())
}

func main() {
	// Start the Game
	game = newGame()

	// Construct the level
	level := newBaseLevel()
	game.Screen().AddEntity(newHUD(level))
	game.Screen().SetLevel(level)

	// Let's A Go
	game.Start()
}

func newGame() *tl.Game {
	debug := flag.Bool("debug", false, "Debug mode")
	fps := flag.Float64("fps", 30.0, "Set frames per second")
	displayFps := flag.Bool("fps-display", false, "Display frames per second on screen")
	flag.Parse()

	g := tl.NewGame()
	g.SetDebugOn(*debug)

	if *displayFps {
		g.Screen().AddEntity(tl.NewFpsText(60, 0, tl.ColorWhite, tl.ColorBlack, 1))
	}

	g.Screen().SetFps(*fps)
	return g
}

func gameOver() {
	screen := tl.NewScreen()
	screen.AddEntity(tl.NewText(0, 1, "GAME OVER", tl.ColorWhite, tl.ColorDefault))
	screen.AddEntity(tl.NewText(0, 2, fmt.Sprintf("Final Score: %d", score), tl.ColorWhite, tl.ColorDefault))
	screen.AddEntity(tl.NewText(0, 4, "Press Ctrl+C to Close", tl.ColorWhite, tl.ColorDefault))
	game.SetScreen(screen)
}
