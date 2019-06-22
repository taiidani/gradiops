package main

import "math/rand"

func initStageOne(lvl *BaseLevel) {
	const textGap = 100

	var textCandidates = []string{
		"DevOps",
		"Synergy",
		"ROI",
		"ShiftLeft",
		"DevSecOps",
		"CI/CD",
	}

	// Generate starting position
	screenWidth, screenHeight := game.Screen().Size()

	for i, text := range textCandidates {
		// Right edge of screen + Offset by previous slots + Random offset within its slot
		x := screenWidth + (i * textGap) + rand.Intn(textGap)
		y := rand.Intn(screenHeight-6) + 3 // Box it in by 3 so that ship bullets can actually hit it

		lvl.AddEntity(newText(text, x, y))
	}
}
