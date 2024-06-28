package internal

import (
	tl "github.com/JoelOtter/termloop"
)

type Game struct {
	*tl.Game
	debug      bool
	displayFps bool
	fps        float64
}

// Global state
var game *Game
var score int

func SetupGame(debug bool, displayFps bool, fps float64) *Game {
	game = &Game{
		Game:       tl.NewGame(),
		debug:      debug,
		displayFps: displayFps,
		fps:        fps,
	}

	game.SetDebugOn(debug)
	return game
}

func (g *Game) Restart() {
	score = 0
	game.Log("Starting new game")

	if g.displayFps {
		game.Screen().AddEntity(tl.NewFpsText(60, 0, tl.ColorWhite, tl.ColorBlack, 1))
	}
	game.Screen().SetFps(g.fps)

	// Construct the level
	level := newBaseLevel()
	game.Screen().AddEntity(newHUD(level))
	game.Screen().SetLevel(level)
}
