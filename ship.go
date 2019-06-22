package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Ship represents a controllable player ship
type Ship struct {
	*tl.Entity
}

func newShip() *Ship {
	const width, height = 10, 6
	const startX, startY = 1, 10
	ship := &Ship{
		Entity: tl.NewEntity(startX, startY, width, height),
	}

	makeCell := func(x int, y int) (cell *tl.Cell) {
		cell = &tl.Cell{}

		switch {
		// Cockpit
		case x == 7 && y == 3:
			cell.Bg = tl.ColorWhite
			cell.Ch = 'O'

		// Gun
		case x == width-1 && y == 3:
			cell.Bg = tl.ColorRed
			cell.Ch = '*'

		// Engines
		case x == 0:
			cell.Bg = tl.ColorYellow
			cell.Ch = '>'

		// Body
		default:
			cell.Fg = tl.RgbTo256Color(80, 80, 50)
			cell.Bg = tl.ColorBlack
			cell.Ch = '-'
		}

		return
	}

	// Build the ship canvas

	// Top/Bottom - Wingtips
	for x := 0; x < 2; x++ {
		ship.SetCell(x, 1, makeCell(x, 1))
		ship.SetCell(x, 5, makeCell(x, 5))
	}

	// Edges - Wings
	for x := 0; x < width/2; x++ {
		ship.SetCell(x, 2, makeCell(x, 2))
		ship.SetCell(x, 4, makeCell(x, 4))
	}

	// Middle - Fuselage
	for x := 0; x < width; x++ {
		ship.SetCell(x, 3, makeCell(x, 3))
	}

	return ship
}

// Tick is triggered in response to user input
func (ship *Ship) Tick(event tl.Event) {
	const constraintRight = 15
	constraintTop := 50 // TODO Get from screen

	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := ship.Position()
		width, height := ship.Size()

		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if x+width < constraintRight {
				ship.SetPosition(x+1, y)
			}
		case tl.KeyArrowLeft:
			if x > 0 {
				ship.SetPosition(x-1, y)
			}
		case tl.KeyArrowUp:
			if y > 0 {
				ship.SetPosition(x, y-1)
			}
		case tl.KeyArrowDown:
			if y+height < constraintTop {
				ship.SetPosition(x, y+1)
			}
		case tl.KeySpace:
			// FIRE DE CANNON!
			game.Log("Bullet fired from ship")
			b := newBullet(ship)
			level.AddEntity(b)
		}
	}
}

// Collide is triggered whenever the entity runs into something
func (ship *Ship) Collide(collision tl.Physical) {
	// Check if it's evil Text we're colliding with
	if _, ok := collision.(*BuzzWord); ok {
		gameOver()
	}
}
