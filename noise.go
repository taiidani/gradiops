package main

import (
	"math"
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

// Noise represents a background filler object
type Noise struct {
	*tl.Text
	speed int
}

func makeNoise(width int, height int, text string, fg tl.Attr, bg tl.Attr) []*Noise {
	ret := []*Noise{}

	// Slow noise
	for i := 0; i < width; i++ {
		// 10% chance to occur
		if rand.Intn(100) <= 10 {
			x := i
			y := rand.Intn(height)
			ret = append(ret, newNoise(x, y, 1, text, fg, bg))
		}
	}

	// Fast noise
	for i := 0; i < width*2; i++ {
		// 5% chance to occur
		if rand.Intn(100) <= 5 {
			x := i
			y := rand.Intn(height)
			ret = append(ret, newNoise(x, y, 2, text, fg, bg))
		}
	}

	return ret
}

func newNoise(x int, y int, speed int, text string, fg tl.Attr, bg tl.Attr) *Noise {
	// Text will be rolled at first draw time
	return &Noise{
		Text:  tl.NewText(x, y, text, fg, bg),
		speed: speed,
	}
}

// Draw will draw the object to the screen
func (m *Noise) Draw(screen *tl.Screen) {
	x, y := m.Position()

	// Drift!
	x = x - int(math.Abs(float64(m.speed)))

	m.SetPosition(x, y)
	m.Text.Draw(screen)
}
