package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/limacat76/aoc2015/code"
)

func main() {
	filename := os.Getenv("HOME") + "/2015/01/data.txt"
	data, _ := ioutil.ReadFile(filename)
	s := string(data)

	fmt.Println(code.CalculateDestination(&s))
}
