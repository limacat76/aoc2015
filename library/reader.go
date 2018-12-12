package library

import (
	"bufio"
	"io"
	"log"
)

// LoopRule signals to ReadRunes if it must continue or stop
type LoopRule bool

const (
	// Break ReadRunes stops on condition defined by algorithm
	Break LoopRule = false
	// Continue ReadRunes continues reading the runes
	Continue LoopRule = true
)

// ReadRunes reads the runes from the instructions reader and applies algorithm on the context interface
func ReadRunes(instructions io.Reader, context interface{}, algorithm func(c rune, context interface{}) LoopRule) {
	r := bufio.NewReader(instructions)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			signal := algorithm(c, context)
			if signal == Break {
				break
			}
		}
	}
}
