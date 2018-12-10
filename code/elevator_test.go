package code

import "testing"

func TestElevatorDestination(t *testing.T) {
	cases := []struct {
		instructions string
		destination  int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, c := range cases {
		got := CalculateDestination(&c.instructions)
		if got != c.destination {
			t.Errorf("CalculateDestination (X) == %v, want %v", got, c.destination)
		}
	}
}

func TestElevatorBasement(t *testing.T) {
	cases := []struct {
		instructions string
		destination  int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, c := range cases {
		got := CalculateBasement(&c.instructions)
		if got != c.destination {
			t.Errorf("CalculateBasement (X) == %v, want %v", got, c.destination)
		}
	}
}
