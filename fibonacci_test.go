package main

import "testing"

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

func TestFibonacci(t *testing.T) {
	for _, f := range casesFibonacci {
		result := fibonacci(f.Value)
		if result != f.Expected {
			t.Errorf("Expected: %v Got:%v", f.Expected, result)
		}
	}
}
