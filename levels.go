package main

import (
	tl "github.com/JoelOtter/termloop"
)

// hudOffset is the number of units that the HUD takes up at the top of the screen
const hudOffset = 2

// BaseLevel is the canvas that all level objects are written to
type BaseLevel struct {
	*tl.BaseLevel
	Ship  *Ship
	Stage int
}

func newBaseLevel() *BaseLevel {
	lvl := &BaseLevel{
		BaseLevel: tl.NewBaseLevel(tl.Cell{
			Bg: tl.RgbTo256Color(10, 10, 50),
			Fg: tl.ColorBlack,
			Ch: '.',
		}),
	}

	// Offset the level to handle HUD
	lvl.SetOffset(0, hudOffset)

	// Lay out persistent resources
	lvl.Ship = newShip()
	lvl.AddEntity(lvl.Ship)

	return lvl
}

// Draw will lay out the level onto the screen
func (m *BaseLevel) Draw(screen *tl.Screen) {
	// If at stage 0 (not started), or all items have been destroyed, advance to the next stage
	levelRunning := false
	for _, entity := range m.Entities {
		if _, ok := entity.(*BuzzWord); ok {
			levelRunning = true
		}
	}

	if !levelRunning {
		m.Stage++
		m.SetStage(m.Stage)
	}

	m.BaseLevel.Draw(screen)
}

// SetStage will initialize the level to a given stage
func (m *BaseLevel) SetStage(stage int) {
	// Clear the level
	for _, entity := range m.Entities {
		if _, ok := entity.(*Ship); ok {
			continue
		}

		m.RemoveEntity(entity)
	}

	// And populate the new one
	m.Stage = stage

	switch stage {
	case 1:
		initStageOne(m)
	case 2:
		initStageTwo(m)
	default:
		gameOver()
	}
}
