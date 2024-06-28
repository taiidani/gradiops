package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

// maxTextLength is the maximum length that text can be.
// It determines the "gutter" where text objects are off screen and can be recycled.
const maxTextLength = 50

// BuzzWord represents a text object that needs destruction
type BuzzWord struct {
	*tl.Text
	IsTainted bool // If tainted, object will be removed during the next draw cycle
}

func newText(word string, x int, y int) *BuzzWord {
	colors := []tl.Attr{
		tl.ColorBlack,
		tl.RgbTo256Color(40, 0, 0),
		tl.RgbTo256Color(0, 40, 0),
		tl.RgbTo256Color(0, 40, 40),
		tl.RgbTo256Color(40, 40, 0),
		tl.RgbTo256Color(40, 40, 40),
	}

	// Text will be rolled at first draw time
	return &BuzzWord{
		Text: tl.NewText(x, y, word, tl.ColorWhite, colors[rand.Intn(len(colors))]),
	}
}

// Draw will draw the object to the screen
func (m *BuzzWord) Draw(screen *tl.Screen) {
	x, y := m.Position()

	// Drift!
	x--

	// If passed out of view; bye
	if x < -maxTextLength {
		m.IsTainted = true
	}

	if m.IsTainted {
		screen.Level().RemoveEntity(m)
		return
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
			score++
			m.SetText(currentText[1:])
		} else {
			// It's dead!
			m.IsTainted = true
		}
	}
}
