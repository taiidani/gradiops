package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

// GameOverLevel is displayed when the game has ended.
type GameOverLevel struct {
	*tl.BaseLevel
}

func gameOver() {
	game.Log("Game ended :(")
	screen := tl.NewScreen()
	lvl := &GameOverLevel{
		BaseLevel: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorDefault,
			Fg: tl.ColorBlack,
			Ch: '.',
		}),
	}

	screen.SetLevel(lvl)
	lvl.AddEntity(tl.NewText(0, 1, "GAME OVER", tl.ColorWhite, tl.ColorDefault))
	lvl.AddEntity(tl.NewText(0, 2, fmt.Sprintf("Final Score: %d", score), tl.ColorWhite, tl.ColorDefault))
	lvl.AddEntity(tl.NewText(0, 4, "Press Return to Retry", tl.ColorWhite, tl.ColorDefault))
	lvl.AddEntity(tl.NewText(0, 5, "Press Ctrl+C to Close", tl.ColorWhite, tl.ColorDefault))
	game.SetScreen(screen)
}

// Tick is triggered in response to user input
func (l *GameOverLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyEnter:
			game.Log("User initiated new game")
			newGame()
		}
	}
}
