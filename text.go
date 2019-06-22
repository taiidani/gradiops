package main

import (
	"math/rand"
	"sync"

	tl "github.com/JoelOtter/termloop"
)

// maxTextLength is the maximum length that text can be.
// It determines the "gutter" where text objects are off screen and can be recycled.
const maxTextLength = 50

// textCandidates are possible choices for text objects
var textCandidates = []string{
	"DevOps",
	"Transformation",
	"Synergy",
	"ROI",
	"ShiftLeft",
	"DevSecOps",
	"CI/CD",
	"ContinuousIntegration",
	"ContinuousDeployment",
}

var deathMutex = sync.Mutex{}

// BuzzWord represents a text object that needs destruction
type BuzzWord struct {
	*tl.Text
}

func newText() *BuzzWord {
	colors := []tl.Attr{
		tl.ColorBlack,
		tl.RgbTo256Color(50, 0, 0),
		tl.RgbTo256Color(0, 50, 0),
		tl.RgbTo256Color(0, 50, 50),
		tl.RgbTo256Color(50, 50, 0),
		tl.RgbTo256Color(50, 50, 50),
	}

	// Text will be rolled at first draw time
	return &BuzzWord{
		Text: tl.NewText(-100, -100, "", tl.ColorWhite, colors[rand.Intn(len(colors))]),
	}
}

// Draw will draw the object to the screen
func (m *BuzzWord) Draw(screen *tl.Screen) {
	x, y := m.Position()

	// If new, roll the starting position
	if len(m.Text.Text()) == 0 {
		// Generate starting position
		screenX, screenY := screen.Size()

		x = screenX
		y = rand.Intn(screenY-6) + 3 // Box it in by 3 so that ship bullets can actually hit it

		m.SetText(textCandidates[rand.Intn(len(textCandidates))])
	} else if x < -maxTextLength {
		// Passed out of view; create a new one
		m.die()
		return
	} else {
		// Drift!
		x--
	}

	m.SetPosition(x, y)
	m.Text.Draw(screen)
}

// Collide is triggered whenever the entity runs into something
func (m *BuzzWord) Collide(collision tl.Physical) {
	// Check if it's a bullet we're colliding with
	if _, ok := collision.(*Bullet); ok {
		currentText := m.Text.Text()
		game.Log("Bullet hit text. Current text value:" + currentText)

		if len(currentText) > 2 {
			// Reduce the text by one
			m.SetText(currentText[2:])
		} else {
			// It's dead!
			m.die()
		}
	}
}

// die triggers the death and rebirth of a BuzzWord
// It can be called from multiple separate goroutines such as Collision methods, and employs a mutex to ensure
// processing is done in an orderly fashion.
func (m *BuzzWord) die() {
	// Prevent multiple deaths (such as collisions) from triggering this action more than once
	deathMutex.Lock()
	defer deathMutex.Unlock()

	for _, entity := range level.Entities {
		// Only remove the entity if it's present, in case a previous concurrent death already did it
		if entity == m {
			level.RemoveEntity(m)
		}
	}
}
