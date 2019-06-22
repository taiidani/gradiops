package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Bullet represents a bullet fired from the Ship
type Bullet struct {
	*tl.Entity
}

func newBullet(ship *Ship) *Bullet {
	const width, height = 1, 1

	// Get Ship point for where bullet emits from
	shipX, shipY := ship.Position()
	shipWidth, _ := ship.Size()
	startX, startY := shipX+shipWidth, shipY+3

	// Create the bullet
	bullet := &Bullet{
		Entity: tl.NewEntity(startX, startY, width, height),
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			bullet.SetCell(x, y, &tl.Cell{Fg: tl.RgbTo256Color(80, 80, 50), Bg: tl.ColorBlack, Ch: '='})
		}
	}

	return bullet
}

// Draw will draw the object to the screen
func (bullet *Bullet) Draw(screen *tl.Screen) {
	screenX, _ := screen.Size()
	x, y := bullet.Position()

	// If not visible, reroll
	if x < screenX {
		x++
	}

	bullet.SetPosition(x, y)
	bullet.Entity.Draw(screen)
}

// Collide is triggered whenever the entity runs into something
func (bullet *Bullet) Collide(collision tl.Physical) {
	// Check if it's evil Text we're colliding with
	if _, ok := collision.(*BuzzWord); ok {
		level.RemoveEntity(bullet)
	}
}
