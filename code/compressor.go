package code

import (
	"bufio"
	"io"
	"strings"

	"github.com/limacat76/aoc2015/library"
)

type compressor struct {
	code, memory, difference, skip int
}

// "" is 2 characters of code (the two double quotes), but the string contains zero characters.
// "abc" is 5 characters of code, but 3 characters in the string data.
// "aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
// "\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.

func calculateString(c rune, context interface{}) library.LoopRule {
	original, ok := context.(*compressor)
	if ok {
		// code is always updated
		original.code++
		if '"' == c && original.skip == 0 {
			original.difference++
		} else if '\\' == c && original.skip == 0 {
			// la sequenza di escape vale almeno un carattere
			original.memory++
			original.skip++
		} else {
			if original.skip > 0 {
				if original.skip == 1 && 'x' == c {
					original.skip = 2
				} else {
					original.skip--
				}
				original.difference++
			} else {
				original.memory++
			}
		}
	}
	return library.Continue
}

// CalculateDestination calculates the destinations of elevator instructions
func RamCalc(aString string) (int, int, int) {
	cfT := &compressor{0, 0, 0, 0}
	library.ReadRunes(strings.NewReader(aString), cfT, calculateString)
	return cfT.code, cfT.memory, cfT.difference
}

type holder struct {
	result string
}

func calculateLiteral(c rune, context interface{}) library.LoopRule {
	original, ok := context.(*holder)
	if ok {
		if '"' == c {
			original.result = original.result + `\"`
		} else if '\\' == c {
			original.result = original.result + `\\`
		} else {
			original.result = original.result + string(c)
		}
	}
	return library.Continue
}

func CompilerLiteral(aString string) string {
	cfT := &holder{`"`}
	library.ReadRunes(strings.NewReader(aString), cfT, calculateLiteral)
	return cfT.result + `"`
}

func RamCalcStream(instructions io.Reader) (int, int, int) {
	tA := 0
	tB := 0
	tC := 0

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		a, b, c := RamCalc(word)
		tA += a
		tB += b
		tC += c
	}
	return tA, tB, tC
}

func RamCalcEncodeStream(instructions io.Reader) (int, int, int) {
	tNewCode := 0
	tOldCode := 0

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		a, _, _ := RamCalc(word)
		tOldCode += a

		b, _, _ := RamCalc(CompilerLiteral(word))
		tNewCode += b
	}
	return tNewCode, tOldCode, (tNewCode - tOldCode)
}
