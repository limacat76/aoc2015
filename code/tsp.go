package code

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func addDistance(a string, b string, w int, x *map[string]bool, y *map[Edge]int) {
	addCity(a, x)
	addCity(b, x)
	addEdge(a, b, w, y)
}

func decodeAndAddDistance(c string, x *map[string]bool, y *map[Edge]int) {
	tokens := strings.Split(c, " ")
	distance, _ := strconv.Atoi(tokens[4])
	addDistance(tokens[0], tokens[2], distance, x, y)
}

func xXx(x *map[string]bool, y *map[Edge]int, z *string) int {
	temp := 9999
	for aCity := range *x {
		last := 0
		if z != nil {
			last = getEdge(*z, aCity, y)
		}
		if len(*x) == 1 {
			return last
		}
		x2 := make(map[string]bool)
		for anotherCity := range *x {
			if anotherCity != aCity {
				x2[anotherCity] = true
			}
		}
		tR := xXx(&x2, y, &aCity) + last
		if tR < temp {
			temp = tR
		}
	}
	return temp
}

func xXy(x *map[string]bool, y *map[Edge]int, z *string) int {
	temp := -1
	for aCity := range *x {
		last := 0
		if z != nil {
			last = getEdge(*z, aCity, y)
		}
		if len(*x) == 1 {
			return last
		}
		x2 := make(map[string]bool)
		for anotherCity := range *x {
			if anotherCity != aCity {
				x2[anotherCity] = true
			}
		}
		tR := xXy(&x2, y, &aCity) + last
		if tR > temp {
			temp = tR
		}
	}
	return temp
}

func TravellingSantaProblem(instructions io.Reader) int {
	citiesset := make(map[string]bool)
	distances := make(map[Edge]int)

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		decodeAndAddDistance(scanner.Text(), &citiesset, &distances)
	}
	return xXx(&citiesset, &distances, nil)
}

func ShowOffSantaProblem(instructions io.Reader) int {
	citiesset := make(map[string]bool)
	distances := make(map[Edge]int)

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		decodeAndAddDistance(scanner.Text(), &citiesset, &distances)
	}
	return xXy(&citiesset, &distances, nil)
}

func addCity(a string, c *map[string]bool) {
	(*c)[a] = true
}

type Edge struct {
	first, second string
}

func addEdge(a string, b string, w int, x *map[Edge]int) {
	(*x)[makeEdge(a, b)] = w
}

func makeEdge(a string, b string) Edge {
	if a < b {
		return Edge{a, b}
	}
	return Edge{b, a}
}

func getEdge(a string, b string, x *map[Edge]int) int {
	return (*x)[makeEdge(a, b)]
}
