package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

func newStageTwo() *Stage {
	const textGap = 100

	var textCandidates = []string{
		"Transformation",
		"InfrastructureAsCode",
		"ContinuousIntegration",
		"ContinuousDeployment",
	}

	ret := &Stage{
		Title: "2",
		Background: tl.Cell{
			Bg: tl.RgbTo256Color(30, 70, 30),
			Fg: tl.ColorBlack,
			Ch: ' ',
		},
	}

	// Generate starting position
	screenWidth, screenHeight := game.Screen().Size()

	// Noise
	ret.Noise = makeNoise(len(textCandidates)*textGap+screenWidth, screenHeight, "⚘", ret.Background.Fg, ret.Background.Bg)

	// Text
	for i, text := range textCandidates {
		// Right edge of screen + Offset by previous slots + Random offset within its slot
		x := screenWidth + (i * textGap) + rand.Intn(textGap)
		y := rand.Intn(screenHeight-7) + 4 // Box it in by 3 so that ship bullets can actually hit it

		ret.Entities = append(ret.Entities, newText(text, x, y))
	}

	return ret
}
