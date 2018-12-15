package library

import (
	"testing"
)

func TestPrint(t *testing.T) {

	cases := []struct {
		position *Point
		wish     string
	}{
		{&Point{X: 0, Y: 0}, "0,0"},
		{&Point{X: 2, Y: 3}, "2,3"},
		{&Point{X: -1, Y: -9}, "-1,-9"},
		{&Point{X: 0, Y: 3}, "0,3"},
	}

	for _, c := range cases {
		got := ToString(c.position)
		if got != c.wish {
			t.Errorf("ReadMap (X) == %v, want %v", got, c.wish)
		}
	}

}
