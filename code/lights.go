package code

import (
	"bufio"
	"io"
	"strings"

	"github.com/limacat76/aoc2015/library"
)

const (
	off = false
	on  = true
)

// TurnOn turns on the lights. Returns how many lights it acted on.
func TurnOn(lights *[][]library.Value, start library.Point, end library.Point) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		wasoff := (*m)[x][y].(bool)
		(*m)[x][y] = on
		return !wasoff
	}, switchAdder)
}

// TurnOff turns off the lights. Returns how many lights it acted on.
func TurnOff(lights *[][]library.Value, start library.Point, end library.Point) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		wason := (*m)[x][y]
		(*m)[x][y] = off
		return wason
	}, switchAdder)
}

// Toggle toggles the lights. Returns how many lights it acted on.
func Toggle(lights *[][]library.Value, start library.Point, end library.Point) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		if (*m)[x][y] == off {
			(*m)[x][y] = on
		} else {
			(*m)[x][y] = off
		}
		return true
	}, switchAdder)
}

// The phrase turn on actually
// The phrase turn off actually
// The phrase toggle actually means that you should increase the brightness of those lights by 2.

// ElvenTurnOn means that you should increase the brightness of those lights by 1.
func ElvenTurnOn(lights *[][]library.Value, start library.Point, end library.Point) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		(*m)[x][y] = (*m)[x][y].(int) + 1
		return on
	}, switchAdder)
}

// ElvenTurnOff means that you should decrease the brightness of those lights by 1, to a minimum of zero.
func ElvenTurnOff(lights *[][]library.Value, start library.Point, end library.Point) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		temp := (*m)[x][y].(int) - 1
		if temp < 0 {
			(*m)[x][y] = 0
		} else {
			(*m)[x][y] = temp
		}
		return on
	}, switchAdder)
}

// ElvenToggle means that you should increase the brightness of those lights by 2.
func ElvenToggle(lights *[][]library.Value, start library.Point, end library.Point) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		(*m)[x][y] = (*m)[x][y].(int) + 2
		return on
	}, switchAdder)
}

// Tally calculates the lights that are on
func Tally(lights *[][]library.Value, start library.Point, end library.Point, adder Adder) library.Result {
	return cycle(lights, start, end, func(m *[][]library.Value, x int, y int) library.Result {
		return (*m)[x][y]
	}, adder)
}

// LightSwitcher types out the kind of function that is returned by the ParseSwitch
type LightSwitcher func(*[][]library.Value, library.Point, library.Point) library.Result

type LightChecker func(*[][]library.Value, int, int) library.Result

// turn on 0
// turn off 0
// toggle 0

type RuleBook struct {
	turnon, turnoff, toggle LightSwitcher
}

// EnglishRuleBook returns the book with English rules
func EnglishRuleBook() RuleBook {
	return RuleBook{TurnOn, TurnOff, Toggle}
}

// ElvenRuleBook returns the book with Elven rules
func ElvenRuleBook() RuleBook {
	return RuleBook{ElvenTurnOn, ElvenTurnOff, ElvenToggle}
}

func ParseSwitch(row string, book RuleBook) (LightSwitcher, library.Point, library.Point) {
	start := 7
	method := book.toggle
	if row[start] == ' ' { // turn on
		method = book.turnon
		start = 8
	} else if row[start] == 'f' { // turn off
		method = book.turnoff
		start = 9
	} else {
		start = 7
	}

	row = row[start:]
	split := strings.Split(row, " ")

	return method, library.FromString(split[0]), library.FromString(split[2])
}

type Adder func(int, library.Result) int

func switchAdder(tally int, x library.Result) int {
	if x.(bool) {
		return tally + 1
	}
	return tally
}

func cumulativeAdder(tally int, x library.Result) int {
	return tally + x.(int)
}

func cycle(lights *[][]library.Value, start library.Point, end library.Point, algo LightChecker, tallyman Adder) library.Result {
	tally := 0
	for x := start.X; x <= end.X; x++ {
		for y := start.Y; y <= end.Y; y++ {
			tally = tallyman(tally, algo(lights, x, y))
		}
	}
	return tally
}

func SantasLightInstructions(instructions io.Reader) library.Result {
	lm := library.MakeValueMapDefValue(1000, 1000, false)

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := scanner.Text()
		f, s, e := ParseSwitch(row, EnglishRuleBook())
		f(&lm, s, e)
	}
	return Tally(&lm, library.Point{X: 0, Y: 0}, library.Point{X: 999, Y: 999}, switchAdder)
}

func SantasTranslatedLightInstructions(instructions io.Reader) library.Result {
	lm := library.MakeValueMapDefValue(1000, 1000, 0)

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := scanner.Text()
		f, s, e := ParseSwitch(row, ElvenRuleBook())
		f(&lm, s, e)
	}
	return Tally(&lm, library.Point{X: 0, Y: 0}, library.Point{X: 999, Y: 999}, cumulativeAdder)
}
