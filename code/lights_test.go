package code

import (
	"testing"

	"github.com/limacat76/aoc2015/library"
)

func TestTurnOn(t *testing.T) {
	// turn on 0,0 through 999,999
	cases := []struct {
		start, end library.Point
		count      int
	}{
		{library.Point{X: 0, Y: 0}, library.Point{X: 999, Y: 999}, 1000000},
	}

	// var lm [][]bool
	for _, c := range cases {
		lm := library.MakeValueMapDefValue(1000, 1000, false)

		got := TurnOn(&lm, c.start, c.end)
		if got != c.count {
			t.Errorf("TurnOn (X) == %v, want %v", got, c.count)
		}
		got2 := Tally(&lm, library.Point{X: 0, Y: 0}, library.Point{X: 999, Y: 999}, switchAdder)
		if got2 != c.count {
			t.Errorf("Tally on TurnOn (X) == %v, want %v", got, c.count)
		}
	}
}

func TestToggle(t *testing.T) {
	// turn on 0,0 through 999,999
	cases := []struct {
		start, end library.Point
		count      int
	}{
		{library.Point{X: 0, Y: 0}, library.Point{X: 999, Y: 0}, 1000},
	}

	for _, c := range cases {
		lm := library.MakeValueMapDefValue(1000, 1000, false)

		got := Toggle(&lm, c.start, c.end)
		if got != c.count {
			t.Errorf("Toggle (X) == %v, want %v", got, c.count)
		}
	}
}

func TestTurnOff(t *testing.T) {
	// turn off 0,0 through 999,999
	cases := []struct {
		start, end library.Point
		count      int
	}{
		{library.Point{X: 499, Y: 499}, library.Point{X: 500, Y: 500}, 4},
	}

	for _, c := range cases {
		lm := library.MakeValueMapDefValue(1000, 1000, false)
		TurnOn(&lm, library.Point{X: 0, Y: 0}, library.Point{X: 999, Y: 999})
		got := TurnOff(&lm, c.start, c.end)
		if got != c.count {
			t.Errorf("TurnOff (X) == %v, want %v", got, c.count)
		}
	}
}

func TestParseSwitch(t *testing.T) {

	// turn on 931,331 through 939,812
	// turn off 756,53 through 923,339
	// toggle 756,965 through 812,992
	cases := []struct {
		instructions string
		defvalue     bool
		start, end   library.Point
		count        int
	}{
		{"turn on 0,0 through 99,99", false, library.Point{X: 0, Y: 0}, library.Point{X: 99, Y: 99}, (100 * 100)},
		{"turn off 0,0 through 99,99", true, library.Point{X: 0, Y: 0}, library.Point{X: 99, Y: 99}, (100 * 100)},
		{"toggle 0,0 through 99,99", false, library.Point{X: 0, Y: 0}, library.Point{X: 99, Y: 99}, (100 * 100)},
	}

	for _, c := range cases {
		lm := library.MakeValueMapDefValue(1000, 1000, c.defvalue)
		method, start, end := ParseSwitch(c.instructions, EnglishRuleBook())
		if start.X != c.start.X || start.Y != c.start.Y {
			t.Errorf("start == %v, want %v", c.start, start)
			continue
		}
		if end.X != c.end.X || end.Y != c.end.Y {
			t.Errorf("end == %v, want %v", c.end, end)
			continue
		}

		got := method(&lm, start, end)
		if got != c.count {
			t.Errorf("ParseSwitch(%v).Method(...) == %v, want %v", c.instructions, got, c.count)
		}
	}

}
