package code

import (
	"bufio"
	"io"
)

// SweetString a string is sweet if it has at least three vowels
func SweetString(s string) bool {
	vowels := 0
	for _, character := range s {
		if 'a' == character || 'e' == character || 'i' == character || 'o' == character || 'u' == character {
			vowels++
		}
		if vowels >= 3 {
			return true
		}
	}
	return false
}

// DiligentString a string is diligent if at least one letter appears twice in a row
func DiligentString(s string) bool {
	lastrune := '\n'
	for _, thisrune := range s {
		if lastrune == thisrune {
			return true
		}
		lastrune = thisrune
	}
	return false
}

// CleanString a string is clean if it does not contain ab, cd, pq or xy
func CleanString(s string) bool {
	lastrune := '\n'
	for _, thisrune := range s {
		if lastrune == 'a' && thisrune == 'b' {
			return false
		}
		if lastrune == 'c' && thisrune == 'd' {
			return false
		}
		if lastrune == 'p' && thisrune == 'q' {
			return false
		}
		if lastrune == 'x' && thisrune == 'y' {
			return false
		}
		lastrune = thisrune
	}
	return true
}

// NiceString a string is nice if it's sweet, diligent and clean
func NiceString(s string) bool {
	return SweetString(s) && DiligentString(s) && CleanString(s)
}

// NaughtyOrNice1 calculates how many nice strings are in the reader
func NaughtyOrNice1(instructions io.Reader) int {
	tally := 0

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		if NiceString(word) {
			tally++
		}
	}
	return tally
}

// SingingString tests if string contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
func SingingString(s string) bool {
	if len(s) < 4 {
		return false
	}

	// regexes could work better
	for idx := 0; idx < len(s)-2; idx++ {
		target := s[idx : idx+2]
		for idx2 := idx + 2; idx2 < len(s)-1; idx2++ {
			target2 := s[idx2 : idx2+2]
			if target == target2 {
				return true
			}
		}
	}
	return false
}

// ClearString tests if string contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
func ClearString(s string) bool {
	tailrune := '\n'
	lastrune := '\n'
	for _, thisrune := range s {
		if tailrune == thisrune {
			return true
		}
		tailrune = lastrune
		lastrune = thisrune
	}
	return false
}

// NiceString2 a string is nice if it's singing and clear
func NiceString2(s string) bool {
	return SingingString(s) && ClearString(s)
}

// NaughtyOrNice2 calculates how many nice strings are in the reader with the new rules the naughty old man gave us oh I hate working for this company
func NaughtyOrNice2(instructions io.Reader) int {
	tally := 0

	scanner := bufio.NewScanner(instructions)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		if NiceString2(word) {
			tally++
		}
	}
	return tally
}
