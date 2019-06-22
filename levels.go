package main

import (
	tl "github.com/JoelOtter/termloop"
)

// MainLevel is the canvas that all level objects are written to
type MainLevel struct {
	*tl.BaseLevel
}

func gameOver() {
	// Wipe the screen
	for _, entity := range level.Entities {
		level.RemoveEntity(entity)
	}

	level.AddEntity(tl.NewText(3, 3, "GAME OVER", tl.ColorWhite, tl.ColorBlack))
}

func newMainLevel() *MainLevel {
	return &MainLevel{
		BaseLevel: tl.NewBaseLevel(tl.Cell{
			Bg: tl.RgbTo256Color(10, 10, 50),
			Fg: tl.ColorBlack,
			Ch: '.',
		}),
	}
}

// Draw will lay out the level onto the screen
func (m *MainLevel) Draw(screen *tl.Screen) {
	m.drawText()
	m.BaseLevel.Draw(screen)
}

// drawText is a Text Reviver -- spawn new text if none left
func (m *MainLevel) drawText() {
	textFound := false
	for _, entity := range m.Entities {
		if _, ok := entity.(*BuzzWord); ok {
			textFound = true
		}
	}
	if !textFound {
		// Like a phoenix [project], bring a new buzz word to life
		m.AddEntity(newText())
	}
}
