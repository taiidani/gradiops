package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Stage contains layout and metadata information regarding the contents of a given stage
type Stage struct {
	Background tl.Cell       // Background contains the background information laid behind the objects
	Title      string        // Title is the name of the stage that is displayed
	Noise      []*Noise      // Entities contains all objects drawn to the stage
	Entities   []tl.Drawable // Entities contains all objects drawn to the stage
}

func (s *Stage) init(lvl *BaseLevel) {
	// Set the background
	lvl.BaseLevel = tl.NewBaseLevel(s.Background)

	// Offset the level to handle HUD
	lvl.SetOffset(0, hudOffset)

	// Lay out background noise
	for _, entity := range s.Noise {
		lvl.AddEntity(entity)
	}

	// Lay out persistent resources
	lvl.Ship = newShip()
	lvl.AddEntity(lvl.Ship)

	// And populate with the new entities
	for _, entity := range s.Entities {
		lvl.AddEntity(entity)
	}
}
