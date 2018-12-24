package main

import (
	"fmt"
	"os"

	"github.com/limacat76/aoc2015/code"
)

func main() {
	filename := os.Getenv("HOME") + "/2015/07/data.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	circuit := code.MakeDirectory(file)
	signal := circuit.Dir["a"].Output(circuit)
	circuit.SetSignal("b", signal)
	fmt.Println(circuit.Dir["a"].Output(circuit))
}
