package main

import (
	tl "github.com/JoelOtter/termloop"
)

// shipGunOffset is the number of units from the top of the ship to where its gun is.
// Used for detetermining the gun position to build bullets from
const shipGunOffset = 2

// Ship represents a controllable player ship
type Ship struct {
	*tl.Entity
	IsFiring bool // If set to true, will emit a bullet during the next Draw cycle
}

func newShip() *Ship {
	const width, height = 10, 5
	const startX, startY = 1, 0
	ship := &Ship{
		Entity: tl.NewEntity(startX, startY, width, height),
	}

	sprite := tl.NewCanvas(width, height)

	makeCell := func(x int, y int) (cell *tl.Cell) {
		cell = &tl.Cell{}

		switch {
		// Cockpit
		case x == 7 && y == 2:
			cell.Bg = tl.ColorWhite
			cell.Ch = '옷'

		// Gun
		case x == width-1 && y == shipGunOffset:
			cell.Bg = tl.ColorRed
			cell.Ch = ' '

		// Engines
		case x == 0:
			cell.Bg = tl.ColorYellow
			cell.Ch = '►'

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
		sprite[x][0] = *makeCell(x, 0)
		sprite[x][4] = *makeCell(x, 4)
	}

	// Edges - Wings
	for x := 0; x < width/2; x++ {
		sprite[x][1] = *makeCell(x, 1)
		sprite[x][3] = *makeCell(x, 3)
	}

	// Middle - Fuselage
	for x := 0; x < width; x++ {
		sprite[x][shipGunOffset] = *makeCell(x, shipGunOffset)
	}

	ship.ApplyCanvas(&sprite)
	return ship
}

// Draw will draw the object to the screen
func (ship *Ship) Draw(screen *tl.Screen) {
	width, _ := ship.Size()
	x, y := ship.Position()

	if ship.IsFiring {
		b := newBullet(x+width, y+shipGunOffset)
		screen.Level().AddEntity(b)
		ship.IsFiring = false
	}

	ship.Entity.Draw(screen)
}

// Tick is triggered in response to user input
func (ship *Ship) Tick(event tl.Event) {
	constraintRight, constraintTop := game.Screen().Size()

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
			if y+height < constraintTop-hudOffset {
				ship.SetPosition(x, y+1)
			}
		case tl.KeySpace:
			// FIRE DE CANNON!
			game.Log("Bullet fired from ship")
			ship.IsFiring = true
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
