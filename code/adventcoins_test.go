package code

import "testing"

func TestIsItCoin(t *testing.T) {
	cases := []struct {
		key      string
		expected bool
	}{
		{"000001dbbfaXXX", true},
		{"abcdef", false},
		{"001dbbfaXXX", false},
		{"0ABC", false},
		{"00000BABBONATALE", true},
	}
	for _, c := range cases {
		got := IsItCoin(c.key)
		if got != c.expected {
			t.Errorf("IsItCoin (X) == %v, want %v", got, c.expected)
		}
	}

}

func TestMine(t *testing.T) {
	cases := []struct {
		key       string
		iteration int
		expected  bool
	}{
		{"abcdef", 609042, false},
		{"abcdef", 609043, true},
		{"pqrstuv", 1048969, false},
		{"pqrstuv", 1048970, true},
	}
	for _, c := range cases {
		got := IsItCoin(CalculateHash(c.key, c.iteration))
		if got != c.expected {
			t.Errorf("CalculateHash (X) == %v, want %v", got, c.expected)
		}
	}
}
