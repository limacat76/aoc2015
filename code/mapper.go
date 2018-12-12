package code

import (
	"fmt"
	"io"

	"github.com/limacat76/aoc2015/library"
)

type point struct {
	x, y int
}

func toString(position *point) string {
	return fmt.Sprintf("%v:%v", position.x, position.y)
}

func moveAndCalculate(who *point, c rune, where map[string]bool) {
	if c == '>' {
		who.x = who.x + 1
	} else if c == '<' {
		who.x = who.x - 1
	} else if c == '^' {
		who.y = who.y + 1
	} else if c == 'v' {
		who.y = who.y - 1
	}
	where[toString(who)] = true
}

type company struct {
	santasPosition point
	roboPosition   point
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
		santasPosition: point{0, 0},
		covered:        make(map[string]bool),
	}
	myCompany.covered[toString(&myCompany.santasPosition)] = true

	library.ReadRunes(instructions, myCompany, lastChristmas)

	return len(myCompany.covered)
}

// ReadMap2 reads the map to Santa and RoboSanta and finds to how many houses they brought gifts
func ReadMap2(instructions io.Reader) int {
	myCompany := &company{
		santasPosition: point{0, 0},
		roboPosition:   point{0, 0},
		covered:        make(map[string]bool),
		santaMoves:     true,
	}
	myCompany.covered[toString(&myCompany.santasPosition)] = true

	library.ReadRunes(instructions, myCompany, thisChristmas)

	return len(myCompany.covered)
}
