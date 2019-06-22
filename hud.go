package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

// HUD is the canvas that all level objects are written to
type HUD struct {
	*tl.BaseLevel
	level *BaseLevel
}

func newHUD(level *BaseLevel) *HUD {
	hud := &HUD{
		BaseLevel: tl.NewBaseLevel(tl.Cell{
			Bg: tl.RgbTo256Color(10, 10, 10),
			Fg: tl.ColorBlack,
			Ch: '.',
		}),
		level: level,
	}

	hud.AddEntity(tl.NewText(0, 0, "Stage: ", tl.ColorWhite, tl.ColorDefault))
	hud.AddEntity(tl.NewText(20, 0, "Score: 0", tl.ColorWhite, tl.ColorDefault))
	return hud
}

// Draw will lay out the level onto the screen
func (hud *HUD) Draw(screen *tl.Screen) {
	hud.Entities[0].(*tl.Text).SetText(fmt.Sprintf("Stage: %d", hud.level.Stage))
	hud.Entities[1].(*tl.Text).SetText(fmt.Sprintf("Score: %d", score))

	hud.BaseLevel.Draw(screen)
}
