package internal

import (
	tl "github.com/JoelOtter/termloop"
)

// Bullet represents a bullet fired from the Ship
type Bullet struct {
	*tl.Entity
	IsTainted bool // If tainted, object will be removed during the next draw cycle
}

func newBullet(startX int, startY int) *Bullet {
	const width, height = 1, 1

	// Create the bullet
	bullet := &Bullet{
		Entity: tl.NewEntity(startX, startY, width, height),
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			bullet.SetCell(x, y, &tl.Cell{Fg: tl.RgbTo256Color(180, 180, 150), Bg: tl.ColorBlack, Ch: '='})
		}
	}

	return bullet
}

// Draw will draw the object to the screen
func (bullet *Bullet) Draw(screen *tl.Screen) {
	screenX, _ := screen.Size()
	x, y := bullet.Position()

	// Bullet advances!
	x++

	// Off screen means its no longer needed, taint it
	if x >= screenX {
		bullet.IsTainted = true
	}

	if bullet.IsTainted {
		screen.Level().RemoveEntity(bullet)
		return
	}

	bullet.SetPosition(x, y)
	bullet.Entity.Draw(screen)
}

// Collide is triggered whenever the entity runs into something
func (bullet *Bullet) Collide(collision tl.Physical) {
	// Check if it's evil Text we're colliding with
	if _, ok := collision.(*BuzzWord); ok {
		bullet.IsTainted = true
	}
}
