package code

import (
	"strings"
	"testing"
)

func TestNeededPaper(t *testing.T) {
	cases := []struct {
		instructions string
		paper        int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
		{"2x3x4\n1x1x10", 58 + 43},
	}

	for _, c := range cases {
		got := NeededPaper(strings.NewReader(c.instructions))
		if got != c.paper {
			t.Errorf("NeededPaper (X) == %v, want %v", got, c.paper)
		}
	}
}

func TestNeededRibbon(t *testing.T) {
	cases := []struct {
		instructions string
		ribbon       int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
		{"2x3x4\n1x1x10", 34 + 14},
	}

	for _, c := range cases {
		got := NeededRibbon(strings.NewReader(c.instructions))
		if got != c.ribbon {
			t.Errorf("NeededRibbon (X) == %v, want %v", got, c.ribbon)
		}
	}
}
