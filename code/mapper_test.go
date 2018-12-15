package code

import (
	"strings"
	"testing"
)

func TestReadMap(t *testing.T) {
	cases := []struct {
		instructions string
		paper        int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, c := range cases {
		got := ReadMap(strings.NewReader(c.instructions))
		if got != c.paper {
			t.Errorf("ReadMap (X) == %v, want %v", got, c.paper)
		}
	}
}

func TestReadMap2(t *testing.T) {
	cases := []struct {
		instructions string
		paper        int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, c := range cases {
		got := ReadMap2(strings.NewReader(c.instructions))
		if got != c.paper {
			t.Errorf("ReadMap2 (X) == %v, want %v", got, c.paper)
		}
	}
}
