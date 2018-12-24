package main

import (
	"fmt"
	"os"

	"github.com/limacat76/aoc2015/code"
)

func main() {
	filename := os.Getenv("HOME") + "/2015/08/data.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(code.RamCalcStream(file))
}
