package code

import "testing"

func TestMyCircuit(t *testing.T) {
	x := circuit{
		dir: make(map[string]gate),
	}
	x.dir["A"] = input{10} // 1010
	x.dir["B"] = input{15} // 1111
	x.dir["C"] = and{"A", "B"}
	x.dir["D"] = input{15}
	x.dir["E"] = and{"B", "D"}
	if x.dir["A"].output(x) != 10 {
		t.Error("A is not 10")
	}
	ft := x.dir["C"].output(x)
	if ft != 10 {
		t.Errorf("C is not 10: %v", ft)
	}
	ft2 := x.dir["E"].output(x)
	if ft2 != 15 {
		t.Errorf("E is not 15: %v", ft)
	}
}

func TestBookletCircuit(t *testing.T) {
	x := circuit{
		dir: make(map[string]gate),
	}
	/*
			123 -> x
		456 -> y
		x AND y -> d
		x OR y -> e
		x LSHIFT 2 -> f
		y RSHIFT 2 -> g
		NOT x -> h
		NOT y -> i
	*/
	x.dir["x"] = input{123} // 1010
	x.dir["y"] = input{456} // 1010
	x.dir["d"] = and{"x", "y"}
	x.dir["e"] = or{"x", "y"}
	x.dir["f"] = lshift{"x", 2}
	x.dir["g"] = rshift{"y", 2}
	x.dir["h"] = not{"x"}
	x.dir["i"] = not{"y"}
	cases := []struct {
		key      string
		expected uint16
	}{
		{"d", 72},
		{"e", 507},
		{"f", 492},
		{"g", 114},
		{"h", 65412},
		{"i", 65079},
		{"x", 123},
		{"y", 456},
	}
	for _, c := range cases {
		got := x.dir[c.key].output(x)
		if got != c.expected {
			t.Errorf("x.dir[%x].output(x) == %v, want %v", c.key, got, c.expected)
		}
	}
}
