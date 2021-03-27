package main

import (
	"testing"
)

func TestMustCalculateMontageMax(t *testing.T) {
	cases := []struct{
		tile string
		max int
	}{
		{"8x8", 64},
		{"4x4", 16},
		{"8x12", 96},
		{"256x256", 65536},
	}
	for i, c := range cases {
		got := mustCalculateMontageMax(c.tile)
		if got != c.max {
			t.Fatalf("case: %d, want: %d, got: %d", i, c.max, got)
		}
	}
}
