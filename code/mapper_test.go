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

func TestPrint(t *testing.T) {

	cases := []struct {
		position *point
		wish     string
	}{
		{&point{0, 0}, "0:0"},
		{&point{2, 3}, "2:3"},
		{&point{-1, -9}, "-1:-9"},
		{&point{0, 3}, "0:3"},
	}

	for _, c := range cases {
		got := toString(c.position)
		if got != c.wish {
			t.Errorf("ReadMap (X) == %v, want %v", got, c.wish)
		}
	}

}
