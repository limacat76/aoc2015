package code

import (
	"strings"
	"testing"
)

func TestAddDistance(t *testing.T) {
	cases := []struct {
		a, b     string
		distance int
	}{
		{"Torino", "Milano", 200},
		{"Torino", "Roma", 400},
		{"Torino", "Palermo", 1400},
	}
	for _, c := range cases {
		citiesset := make(map[string]bool)
		distances := make(map[Edge]int)
		addDistance(c.a, c.b, c.distance, &citiesset, &distances)
		sizeCities := len(citiesset)
		if sizeCities != 2 {
			t.Errorf("addDistance (X) == does not add the %v-%v cities", c.a, c.b)
		} else if citiesset[c.a] == false || citiesset[c.b] == false {
			t.Errorf("addDistance (X) == does not add the %v-%v cities with their names", c.a, c.b)
		}
		size := len(distances)
		if size != 1 {
			t.Errorf("addDistance (X) == does not add the %v-%v edge", c.a, c.b)
		} else if distances[makeEdge(c.a, c.b)] != c.distance {
			t.Errorf("addDistance (X) == does not add the %v-%v edge with size %v", c.a, c.b, c.distance)
		}
	}
}

func TestDecodeDistanceString(t *testing.T) {
	cases := []struct {
		input, a, b string
		distance    int
	}{
		{"London to Dublin = 464", "Dublin", "London", 464},
		{"London to Belfast = 518", "Belfast", "London", 518},
		{"Dublin to Belfast = 141", "Belfast", "Dublin", 141},
	}
	for _, c := range cases {
		citiesset := make(map[string]bool)
		distances := make(map[Edge]int)
		decodeAndAddDistance(c.input, &citiesset, &distances)
		sizeCities := len(citiesset)
		if sizeCities != 2 {
			t.Errorf("addDistance (X) == does not add the %v-%v cities", c.a, c.b)
		} else if citiesset[c.a] == false || citiesset[c.b] == false {
			t.Errorf("addDistance (X) == does not add the %v-%v cities with their names", c.a, c.b)
		}
		size := len(distances)
		if size != 1 {
			t.Errorf("addDistance (X) == does not add the %v-%v edge", c.a, c.b)
		} else if distances[makeEdge(c.a, c.b)] != c.distance {
			t.Errorf("addDistance (X) == does not add the %v-%v edge with size %v", c.a, c.b, c.distance)
		}
	}
}

func TestFastestSanta(t *testing.T) {
	cases := []struct {
		instructions string
		distance     int
	}{
		{"London to Dublin = 464", 464},
		{"London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141", 605},
	}
	for _, c := range cases {
		got := TravellingSantaProblem(strings.NewReader(c.instructions))
		if got != c.distance {
			t.Errorf("the wrong distance, %v wants %v", got, c.distance)
		}
	}
}

func TestSlowestSanta(t *testing.T) {
	cases := []struct {
		instructions string
		distance     int
	}{
		// {"London to Dublin = 464", 464},
		{"London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141", 982},
		{"A to B = 1\nA to C = 99\nA to D = 1\nB to C = 1\nB to D = 2\nC to D = 1", 99 + 3},
	}
	for _, c := range cases {
		got := ShowOffSantaProblem(strings.NewReader(c.instructions))
		if got != c.distance {
			t.Errorf("the wrong distance, %v wants %v", got, c.distance)
		}
	}
}

// TODO
// 4a. test fastest santa
// 4a. fastest santa
