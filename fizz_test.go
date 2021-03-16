package main

import "testing"

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

func TestFizzBuzz(t *testing.T) {
	for _, fb := range casesFizzBuzz {
		result := fizzBuzz(fb.Value)
		if result != fb.Expected {
			t.Errorf("Expected: %v Got:%v %v", fb.Expected, result, fb.Value)
		}
	}
}
