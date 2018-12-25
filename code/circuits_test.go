package code

import (
	"reflect"
	"strings"
	"testing"
)

func TestMyCircuit(t *testing.T) {
	x := Circuit{
		Dir: make(map[string]Gate),
	}
	x.Dir["A"] = input{"10"} // 1010
	x.Dir["B"] = input{"15"} // 1111
	x.Dir["C"] = and{"A", "B"}
	x.Dir["D"] = input{"15"}
	x.Dir["E"] = and{"B", "D"}
	if x.Dir["A"].Output(x) != 10 {
		t.Error("A is not 10")
	}
	ft := x.Dir["C"].Output(x)
	if ft != 10 {
		t.Errorf("C is not 10: %v", ft)
	}
	ft2 := x.Dir["E"].Output(x)
	if ft2 != 15 {
		t.Errorf("E is not 15: %v", ft)
	}
}

func TestBookletCircuit(t *testing.T) {
	x := Circuit{
		Dir: make(map[string]Gate),
	}
	x.Dir["x"] = input{"123"}
	x.Dir["y"] = input{"456"}
	x.Dir["d"] = and{"x", "y"}
	x.Dir["e"] = or{"x", "y"}
	x.Dir["f"] = lshift{"x", 2}
	x.Dir["g"] = rshift{"y", 2}
	x.Dir["h"] = not{"x"}
	x.Dir["i"] = not{"y"}
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
		got := x.Dir[c.key].Output(x)
		if got != c.expected {
			t.Errorf("x.Dir[%x].output(x) == %v, want %v", c.key, got, c.expected)
		}
	}
}

func TestMakeGate(t *testing.T) {
	cases := []struct {
		rule, expectedKey string
		expectedType      string
	}{
		{"255 -> x", "x", reflect.TypeOf(input{}).Name()},
		{"x AND y -> d", "d", reflect.TypeOf(and{}).Name()},
		{"x OR y -> d", "d", reflect.TypeOf(or{}).Name()},
		{"x LSHIFT 2 -> d", "d", reflect.TypeOf(lshift{}).Name()},
		{"x RSHIFT 2 -> d", "d", reflect.TypeOf(rshift{}).Name()},
		{"NOT x -> h", "h", reflect.TypeOf(not{}).Name()},
	}

	for _, c := range cases {
		gotType, gotKey := MakeGate(c.rule)
		if gotKey != c.expectedKey {
			t.Errorf("MakeGate() key == %v, want %v", gotKey, c.expectedKey)
			continue
		}
		gotTypeName := reflect.TypeOf(gotType).Name()
		if gotTypeName != c.expectedType {
			t.Errorf("MakeGate() tyoe == %v, want %v", gotTypeName, c.expectedType)
			continue
		}
	}

}

func TestCircuitRules(t *testing.T) {

	cases := []struct {
		rule, key string
		expected  uint16
	}{
		{"255 -> x", "x", 255},
		{"255 -> x\n255 -> y\nx AND y -> z", "z", 255},
		{"1 -> x\n2 -> y\nx OR y -> z", "z", 3},
	}

	for _, c := range cases {
		circuit := MakeDirectory(strings.NewReader(c.rule))
		got := circuit.Dir[c.key].Output(circuit)
		if got != c.expected {
			t.Errorf("x.Dir[%v].output(x) == %v, want %v", c.key, got, c.expected)
		}
	}
}
