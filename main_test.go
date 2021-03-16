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
	{20, "FIZZ"},
	{21, "FIZZ"},
	{24, "FIZZ"},
	{25, "BUZZ"},
	{30, "Fizz BUZZ"},
	{35, "BUZZ"},
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
			t.Errorf("Expected: %v Got:%v", fb.Expected, result)
		}
	}
}
