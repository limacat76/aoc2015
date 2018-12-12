package code

import (
	"io"

	"github.com/limacat76/aoc2015/library"
)

type elevator struct {
	floor, position int
}

func calculateDest(c rune, context interface{}) library.LoopRule {
	original, ok := context.(*elevator)
	if ok {
		if '(' == c {
			original.floor++
		} else if ')' == c {
			original.floor--
		}
	}
	return library.Continue
}

func calculateBasement(c rune, context interface{}) library.LoopRule {
	original, ok := context.(*elevator)
	if ok {
		if '(' == c {
			original.floor++
		} else if ')' == c {
			original.floor--
		}
		if original.floor < 0 {
			// let's stop here
			return library.Break
		}
		original.position++
	}
	return library.Continue
}

// CalculateDestination calculates the destinations of elevator instructions
func CalculateDestination(instructions io.Reader) int {
	cfT := &elevator{0, 1}
	library.ReadRunes(instructions, cfT, calculateDest)
	return cfT.floor
}

// CalculateBasement calculates at which instruction step Santa enters the basement
func CalculateBasement(instructions io.Reader) int {
	cfT := &elevator{0, 1}
	library.ReadRunes(instructions, cfT, calculateBasement)
	return cfT.position
}
