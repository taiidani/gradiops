package internal

import (
	tl "github.com/JoelOtter/termloop"
)

// hudOffset is the number of units that the HUD takes up at the top of the screen
const hudOffset = 2

// BaseLevel is the canvas that all level objects are written to
type BaseLevel struct {
	*tl.BaseLevel
	Ship     *Ship
	stageNum int
	stage    *Stage
}

func newBaseLevel() *BaseLevel {
	lvl := &BaseLevel{
		BaseLevel: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorDefault,
			Fg: tl.ColorBlack,
			Ch: '.',
		}),
	}

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
		m.stageNum++
		m.SetStage(m.stageNum)
	}

	m.BaseLevel.Draw(screen)
}

// SetStage will initialize the level to a given stage
func (m *BaseLevel) SetStage(stage int) {
	switch stage {
	case 1:
		m.stage = newStageOne()
	case 2:
		m.stage = newStageTwo()
	default:
		gameOver()
	}

	m.stage.init(m)
}
