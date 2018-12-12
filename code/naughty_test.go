package code

import (
	"strings"
	"testing"
)

func TestSweet(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"aei", true},
		{"eee", true},
		{"xabiwe", true},
		{"cde", false},
		{"kde", false},
		{"msn", false},
		// COMMON TESTS
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"dvszwmarrgswjxmb", false},
		{"jchzalrnumimnmhp", true},
		{"haegwjzuvuyypxyu", true},
	}
	for _, c := range cases {
		got := SweetString(c.key)
		if got != c.expected {
			t.Errorf("SweetString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}
}

func TestDiligent(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"xx", true},
		{"abcdde", true},
		{"aabbccdd", true},
		{"cde", false},
		{"kde", false},
		{"msn", false},
		// COMMON TESTS
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"dvszwmarrgswjxmb", true},
	}
	for _, c := range cases {
		got := DiligentString(c.key)
		if got != c.expected {
			t.Errorf("DiligentString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}
}

func TestClean(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"aei", true},
		{"xab", false},
		{"zcd", false},
		{"spq", false},
		{"wxy", false},
		// COMMON TESTS
		{"ugknbfddgicrmopn", true},
		{"haegwjzuvuyypxyu", false},
	}
	for _, c := range cases {
		got := CleanString(c.key)
		if got != c.expected {
			t.Errorf("CleanString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}
}

func TestNiceString(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, c := range cases {
		got := NiceString(c.key)
		if got != c.expected {
			t.Errorf("CleanString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}

}

func TestNaughtyOrNice1(t *testing.T) {
	cases := []struct {
		key      string
		expected int
	}{
		{"ugknbfddgicrmopn\naaa\njchzalrnumimnmhp\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb", 2},
	}
	for _, c := range cases {
		got := NaughtyOrNice1(strings.NewReader(c.key))
		if got != c.expected {
			t.Errorf("CleanString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}

}

func TestSinging(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"w", false},
		{"ww", false},
		{"www", false},
		{"wwww", true},
		{"wweww", true},
		{"wweef", false},
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
		// COMMON TESTS
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"ieodomkazucvgmuy", false},
	}
	for _, c := range cases {
		got := SingingString(c.key)
		if got != c.expected {
			t.Errorf("SingingString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}

}

func TestClear(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
		{"xx", false},
		{"wxy", false},
		// COMMON TESTS
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
	}
	for _, c := range cases {
		got := ClearString(c.key)
		if got != c.expected {
			t.Errorf("ClearString (%v) == %v, want %v", c.key, got, c.expected)
		}
	}
}

func TestNiceString2(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"nxxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}
	for _, c := range cases {
		got := NiceString2(c.key)
		if got != c.expected {
			t.Errorf("NiceString2 (%v) == %v, want %v", c.key, got, c.expected)
		}
	}

}

func TestNaughtyOrNice2(t *testing.T) {
	cases := []struct {
		key      string
		expected int
	}{
		{"qjhvhtzxzqqjkmpb\nxxyxx\nuurcxstgmygtbstg\nieodomkazucvgmuy", 2},
	}
	for _, c := range cases {
		got := NaughtyOrNice2(strings.NewReader(c.key))
		if got != c.expected {
			t.Errorf("NaughtyOrNice2 (X) == %v, want %v", got, c.expected)
		}
	}

}
