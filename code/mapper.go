package code

import (
	"io"

	"github.com/limacat76/aoc2015/library"
)

func moveAndCalculate(who *library.Point, c rune, where map[string]bool) {
	if c == '>' {
		who.X = who.X + 1
	} else if c == '<' {
		who.X = who.X - 1
	} else if c == '^' {
		who.Y = who.Y + 1
	} else if c == 'v' {
		who.Y = who.Y - 1
	}
	where[library.ToString(who)] = true
}

type company struct {
	santasPosition library.Point
	roboPosition   library.Point
	covered        map[string]bool
	santaMoves     bool
}

func lastChristmas(c rune, context interface{}) library.LoopRule {
	original, ok := context.(*company)
	if ok {
		moveAndCalculate(&original.santasPosition, c, original.covered)
	}
	return library.Continue
}

func thisChristmas(c rune, context interface{}) library.LoopRule {
	original, ok := context.(*company)
	if ok {
		if original.santaMoves {
			moveAndCalculate(&original.santasPosition, c, original.covered)
		} else {
			moveAndCalculate(&original.roboPosition, c, original.covered)
		}
		original.santaMoves = !original.santaMoves
	}
	return library.Continue
}

// ReadMap reads the map to Santa and finds to how many houses he brought gifts
func ReadMap(instructions io.Reader) int {
	myCompany := &company{
		santasPosition: library.Point{X: 0, Y: 0},
		covered:        make(map[string]bool),
	}
	myCompany.covered[library.ToString(&myCompany.santasPosition)] = true

	library.ReadRunes(instructions, myCompany, lastChristmas)

	return len(myCompany.covered)
}

// ReadMap2 reads the map to Santa and RoboSanta and finds to how many houses they brought gifts
func ReadMap2(instructions io.Reader) int {
	myCompany := &company{
		santasPosition: library.Point{X: 0, Y: 0},
		roboPosition:   library.Point{X: 0, Y: 0},
		covered:        make(map[string]bool),
		santaMoves:     true,
	}
	myCompany.covered[library.ToString(&myCompany.santasPosition)] = true

	library.ReadRunes(instructions, myCompany, thisChristmas)

	return len(myCompany.covered)
}
