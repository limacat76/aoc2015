package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/limacat76/aoc2015/code"
)

func main() {
	filename := os.Getenv("HOME") + "/2015/04/data.txt"
	data, _ := ioutil.ReadFile(filename)
	key := strings.Trim(string(data), "\n")

	found := false
	iteration := 1
	for {
		found = code.IsItCoin6(code.CalculateHash(key, iteration))
		if found {
			break
		}
		iteration++
	}

	fmt.Println(iteration)
}
