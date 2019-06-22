package main

import (
	"testing"
)

func TestMaxTextLength(t *testing.T) {
	for _, text := range textCandidates {
		if len(text) > maxTextLength {
			t.Errorf("%s is beyond the maximum length of %d characters", text, maxTextLength)
		}
	}
}
