package internal

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

func newStageOne() *Stage {
	const textGap = 100

	var textCandidates = []string{
		"DevOps",
		"Synergy",
		"ROI",
		"ShiftLeft",
		"DevSecOps",
		"CI/CD",
		"Guild",
		"Scrum",
		"Kanban",
		"Sprint",
		"Agile",
		"ChatOps",
		"GitOps",
		"Lambda",
		"Cloud",
		"Docker",
	}

	ret := &Stage{
		Title: "1",
		Background: tl.Cell{
			Bg: tl.RgbTo256Color(10, 10, 50),
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
		y := rand.Intn(screenHeight-7) + 4 // Box it in so that ship bullets can actually hit it

		ret.Entities = append(ret.Entities, newText(text, x, y))
	}

	return ret
}
