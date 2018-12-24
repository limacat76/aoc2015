package code

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type registry interface {
	fetch(name string) Gate
}

type Circuit struct {
	Dir map[string]Gate
}

var ccache = make(map[string]uint16)

func (r Circuit) SetSignal(gate string, signal uint16) {
	r.Dir[gate] = input{fmt.Sprint(signal)}
	ccache = make(map[string]uint16)
}

func (r Circuit) fetch(name string) Gate {
	res := r.Dir[name]
	// r.Dir[name] = nil
	return res
}

// gates
type Gate interface {
	Output(r registry) uint16
}

// Unary Gates
type input struct {
	x string
}

func (g input) Output(r registry) uint16 {
	return fetchAndExecute(r, g.x)
}

type not struct {
	x string
}

func fetchAndExecute(r registry, gateSpec string) uint16 {
	s2, err := strconv.Atoi(gateSpec)
	if err != nil {
		if val, ok := ccache[gateSpec]; ok {
			fmt.Printf("%v,", gateSpec)
			return val
		}
		// fmt.Printf("%v,", gateSpec)
		result := r.fetch(gateSpec).Output(r)
		ccache[gateSpec] = result
		return result
	}
	return uint16(s2)
}

func (g not) Output(r registry) uint16 {
	xg := fetchAndExecute(r, g.x)
	result := ^xg
	return result
}

// Binary Gates
type and struct {
	x, y string
}

func (g and) Output(r registry) uint16 {
	xg := fetchAndExecute(r, g.x)
	yg := fetchAndExecute(r, g.y)

	return xg & yg
}

type or struct {
	x, y string
}

func (g or) Output(r registry) uint16 {
	xg := fetchAndExecute(r, g.x)
	yg := fetchAndExecute(r, g.y)

	return xg | yg
}

// Shift Operations
type lshift struct {
	x    string
	bits uint
}

func (g lshift) Output(r registry) uint16 {
	xg := fetchAndExecute(r, g.x)
	return xg << g.bits
}

type rshift struct {
	x    string
	bits uint
}

func (g rshift) Output(r registry) uint16 {
	xg := fetchAndExecute(r, g.x)
	return xg >> g.bits
}

/*
	{"255 -> x", "x", reflect.TypeOf(input{}).Name()},
	{"x AND y -> d", "d", reflect.TypeOf(and{}).Name()},
	{"x OR y -> d", "d", reflect.TypeOf(or{}).Name()},
	{"x LSHIFT 2 -> d", "d", reflect.TypeOf(lshift{}).Name()},
	{"x RSHIFT 2 -> d", "d", reflect.TypeOf(rshift{}).Name()},
	{"NOT x -> h", "h", reflect.TypeOf(not{}).Name()},
*/

// MakeGate creates a single gate from a string specification
func MakeGate(row string) (Gate, string) {
	s := strings.Split(row, " -> ")

	outputWire := s[len(s)-1]
	gateRule := strings.Split(s[0], " ")
	var result Gate
	switch len(gateRule) {
	case 1:
		result = input{s[0]}
		break
	case 2:
		result = not{gateRule[1]}
		break
	case 3:
		first := gateRule[0]
		second := gateRule[2]
		switch gateRule[1] {
		case "AND":
			fmt.Printf("and mk %v %v\n", first, second)
			result = and{first, second}
			break
		case "OR":
			fmt.Printf("OR mk %v %v\n", first, second)
			result = or{first, second}
			break
		case "LSHIFT":
			fmt.Printf("LSHIFT mk %v %v\n", first, second)
			s2, _ := strconv.Atoi(second)
			result = lshift{first, uint(s2)}
			break
		case "RSHIFT":
			fmt.Printf("RSHIFT mk %v %v\n", first, second)
			s2, _ := strconv.Atoi(second)
			result = rshift{first, uint(s2)}
			break
		}
	}
	return result, outputWire
}

// MakeDirectory creates a circuit directory from the instruction booklet
func MakeDirectory(instructions io.Reader) Circuit {
	x := Circuit{
		Dir: make(map[string]Gate),
	}

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := scanner.Text()
		gate, s := MakeGate(row)
		fmt.Printf("%v, %v\n", s, gate)
		x.Dir[s] = gate
	}

	return x
}
