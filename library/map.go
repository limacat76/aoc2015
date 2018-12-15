package library

import (
	"fmt"
	"strconv"
	"strings"
)

// Point maps a structure in space, not in time
type Point struct {
	X, Y int
}

// ToString represents a point into a string
func ToString(position *Point) string {
	return fmt.Sprintf("%v,%v", position.X, position.Y)
}

// FromString converts a string into a point
func FromString(aString string) Point {
	split := strings.Split(aString, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	return Point{X: x, Y: y}
}

// MakeBoolMap creates a map of bools of fixed dimensions. Hoping we can make this generic in a future release
func MakeBoolMap(X int, Y int) [][]bool {
	lm := make([][]bool, X)
	for i := 0; i < X; i++ {
		lm[i] = make([]bool, Y)
	}
	return lm
}

// MakeBoolMapDefValue creates a map of bools of fixed dimensions and a default value
func MakeBoolMapDefValue(X int, Y int, defvalue bool) *[][]bool {
	lm := make([][]bool, X)
	for i := 0; i < X; i++ {
		lm[i] = make([]bool, Y)
		for j := 0; j < Y; j++ {
			lm[i][j] = defvalue
		}
	}
	return &lm
}

// MakeIntMap creates a map of bools of fixed dimensions. Hoping we can make this generic in a future release
func MakeIntMap(X int, Y int) [][]int {
	lm := make([][]int, X)
	for i := 0; i < X; i++ {
		lm[i] = make([]int, Y)
	}
	return lm
}

// MakeIntMapDefValue creates a map of bools of fixed dimensions and a default value
func MakeIntMapDefValue(X int, Y int, defvalue int) *[][]int {
	lm := make([][]int, X)
	for i := 0; i < X; i++ {
		lm[i] = make([]int, Y)
		for j := 0; j < Y; j++ {
			lm[i][j] = defvalue
		}
	}
	return &lm
}

// MakeValueMapDefValue creates a map of bools of fixed dimensions and a default value
func MakeValueMapDefValue(X int, Y int, defvalue Value) [][]Value {
	lm := make([][]Value, X)
	for i := 0; i < X; i++ {
		lm[i] = make([]Value, Y)
		for j := 0; j < Y; j++ {
			lm[i][j] = defvalue
		}
	}
	return lm
}
