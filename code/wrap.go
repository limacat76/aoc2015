package code

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func calculateWrap(measuresSi []int) int {
	area := 2*(measuresSi[0]*measuresSi[1]) +
		2*(measuresSi[1]*measuresSi[2]) +
		2*(measuresSi[2]*measuresSi[0])
	extra := measuresSi[0] * measuresSi[1]
	return area + extra
}

func calculateRibbon(measuresSi []int) int {
	ribbon := 2*measuresSi[0] +
		2*(measuresSi[1])
	bow := measuresSi[0] * measuresSi[1] * measuresSi[2]
	return ribbon + bow
}

func calculatorElf(instructions io.Reader, function func([]int) int) int {
	tally := 0

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		measures := strings.Split(word, "x")
		measuresSi := make([]int, 0, len(measures))
		for _, a := range measures {
			i, _ := strconv.Atoi(a)
			measuresSi = append(measuresSi, i)
		}
		sort.Ints(measuresSi)

		tally = tally + function(measuresSi)
	}
	return tally
}

// NeededPaper calculates the paper needed to wrap the present
func NeededPaper(instructions io.Reader) int {
	return calculatorElf(instructions, calculateWrap)
}

// NeededRibbon calculates the feet of ribbon needed to wrap the present
func NeededRibbon(instructions io.Reader) int {
	return calculatorElf(instructions, calculateRibbon)
}
