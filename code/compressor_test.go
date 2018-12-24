package code

import (
	"strings"
	"testing"
)

// "" is 2 characters of code (the two double quotes), but the string contains zero characters.
// "abc" is 5 characters of code, but 3 characters in the string data.
// "aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
// "\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.

func TestRamCalc(t *testing.T) {
	cases := []struct {
		instructions             string
		code, memory, difference int
	}{
		{`""`, 2, 0, 2},
		{`"abc"`, 5, 3, 2},
		{`"aaa\"aaa"`, 10, 7, 3},
		{`"\x27"`, 6, 1, 5},
	}

	for _, c := range cases {
		gotCode, gotMemory, gotDifference := RamCalc(c.instructions)
		if gotCode != c.code {
			t.Errorf("RamCalc (%v) code == %v, want %v", c.instructions, gotCode, c.code)
		} else if gotMemory != c.memory {
			t.Errorf("RamCalc (%v) memory == %v, want %v", c.instructions, gotMemory, c.memory)
		} else if gotDifference != c.difference {
			t.Errorf("RamCalc (%v) difference == %v, want %v", c.instructions, gotDifference, c.difference)
		}
	}

}

// "" encodes to , an increase from 2 characters to 6.
// "abc" encodes to , an increase from 5 characters to 9.
// "aaa\"aaa" encodes to "\"aaa\\\"aaa\"", an increase from 10 characters to 16.
// "\x27" encodes to , an increase from 6 characters to 11.

func TestRamEncode(t *testing.T) {
	cases := []struct {
		instructions, result string
	}{
		{`""`, `"\"\""`},
		{`"abc"`, `"\"abc\""`},
		{`"aaa\"aaa"`, `"\"aaa\\\"aaa\""`},
		{`"\x27"`, `"\"\\x27\""`},
	}

	for _, c := range cases {
		gotEncode := CompilerLiteral(c.instructions)
		if gotEncode != c.result {
			t.Errorf("CompilerLiteral (%v) code == %v, want %v", c.instructions, gotEncode, c.result)
		}
	}

}

func TestRamCalcStream(t *testing.T) {
	cases := []struct {
		instructions             string
		code, memory, difference int
	}{
		{
			`""
"abc"
"aaa\"aaa"
"\x27"`, 2 + 5 + 10 + 6, 0 + 3 + 7 + 1, 2 + 2 + 3 + 5},
	}

	for _, c := range cases {
		gotCode, gotMemory, gotDifference := RamCalcStream(strings.NewReader(c.instructions))
		if gotCode != c.code {
			t.Errorf("RamCalc (X) code == %v, want %v", gotCode, c.code)
		} else if gotMemory != c.memory {
			t.Errorf("RamCalc (X) memory == %v, want %v", gotMemory, c.memory)
		} else if gotDifference != c.difference {
			t.Errorf("RamCalc (X) difference == %v, want %v", gotDifference, c.difference)
		}
	}

}

func TestRamCalcEncodeStream(t *testing.T) {
	cases := []struct {
		instructions                 string
		newCode, oldCode, difference int
	}{
		{
			`""
"abc"
"aaa\"aaa"
"\x27"`, 6 + 9 + 16 + 11, 2 + 5 + 10 + 6, (6 + 9 + 16 + 11) - (2 + 5 + 10 + 6)},
	}

	for _, c := range cases {
		gotCode, gotOldCode, gotDifference := RamCalcEncodeStream(strings.NewReader(c.instructions))
		if gotCode != c.newCode {
			t.Errorf("RamCalcStream (X) newCode == %v, want %v", gotCode, c.newCode)
		} else if gotOldCode != c.oldCode {
			t.Errorf("RamCalcStream (X) oldCode == %v, want %v", gotOldCode, c.oldCode)
		} else if gotDifference != c.difference {
			t.Errorf("RamCalcStream (X) difference == %v, want %v", gotDifference, c.difference)
		}
	}

}
