package main

import (
	"flag"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var game *tl.Game
var level *MainLevel

func init() {
	// Seed the faker library
	rand.Seed(time.Now().Unix())
}

func main() {
	debug := flag.Bool("debug", false, "Debug Mode")
	flag.Parse()

	game = tl.NewGame()
	game.SetDebugOn(*debug)
	game.Screen().SetFps(30)

	// Construct the level
	level = newMainLevel()
	ship := newShip()
	level.AddEntity(ship)

	game.Screen().SetLevel(level)
	game.Start()
}
