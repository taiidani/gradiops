package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var level *MainLevel

func init() {
	// Seed the faker library
	rand.Seed(time.Now().Unix())
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(30)

	level = newMainLevel()
	ship := newShip()
	level.AddEntity(ship)

	g.Screen().SetLevel(level)
	g.Start()
}
