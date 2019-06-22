package main

import "math/rand"

func initStageTwo(lvl *BaseLevel) {
	const levelWidth = 500

	var textCandidates = []string{
		"Transformation",
		"InfrastructureAsCode",
		"ContinuousIntegration",
		"ContinuousDeployment",
	}

	// Generate starting position
	screenWidth, screenHeight := game.Screen().Size()

	for _, text := range textCandidates {
		x := rand.Intn(levelWidth) + screenWidth
		y := rand.Intn(screenHeight-6) + 3 // Box it in by 3 so that ship bullets can actually hit it

		lvl.AddEntity(newText(text, x, y))
	}
}
