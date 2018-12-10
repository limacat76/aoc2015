package code

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

// ReadMap reads the map to Santa and finds to how many houses he brought gifts
func ReadMap(instructions io.Reader) int {
	covered := make(map[string]bool)

	// starting home
	santasPosition := &point{0, 0}
	covered[toString(santasPosition)] = true
	r := bufio.NewReader(instructions)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			moveAndCalculate(santasPosition, c, covered)
		}
	}

	return len(covered)
}

// ReadMap2 reads the map to Santa and RoboSanta and finds to how many houses they brought gifts
func ReadMap2(instructions io.Reader) int {
	covered := make(map[string]bool)

	// starting home
	santasPosition := &point{0, 0}
	roboPosition := &point{0, 0}
	covered[toString(santasPosition)] = true
	santaMoves := true

	r := bufio.NewReader(instructions)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if santaMoves {
				moveAndCalculate(santasPosition, c, covered)
			} else {
				moveAndCalculate(roboPosition, c, covered)
			}
			santaMoves = !santaMoves
		}
	}

	return len(covered)
}
