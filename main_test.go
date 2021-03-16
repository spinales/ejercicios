package main

import "testing"

var casesPrimos = []struct {
	Value    int
	Expected bool
}{
	{3, true},
	{5, true},
	{7, true},
	{19, true},
	{97, true},
	{4, false},
	{25, false},
	{49, false},
	{91, false},
}

var casesFizzBuzz = []struct {
	Value    int
	Expected string
}{
	{12, "FIZZ"},
	{15, "FIZZ BUZZ"},
	{18, "FIZZ"},
	{20, "BUZZ"},
	{21, "FIZZ"},
	{24, "FIZZ"},
	{25, "BUZZ"},
	{30, "FIZZ BUZZ"},
	{35, "BUZZ"},
}

var casesFibonacci = []struct {
	Value    int
	Expected int
}{
	{1, 0},
	{2, 1},
	{3, 1},
	{4, 2},
	{5, 3},
	{6, 5},
	{7, 8},
	{8, 13},
	{9, 21},
	{10, 34},
	{11, 55},
	{12, 89},
	{13, 144},
	{14, 233},
	{15, 377},
	{16, 610},
	// {17, 800},
}

var casesReverse = []struct {
	Value    string
	Expected string
}{
	{"juan", "nauj"},
	{"oso", "oso"},
	{"World", "dlroW"},
	{"Hello", "olleH"},
	{"Geeks", "skeeG"},
	{"Lee", "eeL"},
}

func TestPrimo(t *testing.T) {
	for _, cc := range casesPrimos {
		if CheckPrimo(cc.Value) != cc.Expected {
			t.Errorf("Expected: %v Got:%v", cc.Expected, !cc.Expected)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	for _, fb := range casesFizzBuzz {
		result := FizzBuzz(fb.Value)
		if result != fb.Expected {
			t.Errorf("Expected: %v Got:%v %v", fb.Expected, result, fb.Value)
		}
	}
}

func TestFibonacci(t *testing.T) {
	for _, f := range casesFibonacci {
		result := Fibonacci(f.Value)
		if result != f.Expected {
			t.Errorf("Expected: %v Got:%v", f.Expected, result)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, r := range casesReverse {
		result := reverse(r.Value)
		if result != r.Expected {
			t.Errorf("Expected: %s Got:%s", r.Expected, result)
		}
	}
}
