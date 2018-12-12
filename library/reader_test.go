package library

import (
	"strings"
	"testing"
)

type contextForTest struct {
	floor int
}

func algorithmForTest(c rune, context interface{}) LoopRule {
	original, ok := context.(*contextForTest)
	if ok {
		if '(' == c {
			original.floor++
		} else if ')' == c {
			original.floor--
		}
	}
	return Continue
}

func TestReadRunes(t *testing.T) {
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
		cfT := &contextForTest{0}
		ReadRunes(strings.NewReader(c.instructions), cfT, algorithmForTest)
		got := cfT.floor
		if got != c.destination {
			t.Errorf("CalculateDestination (X) == %v, want %v", got, c.destination)
		}
	}
}
