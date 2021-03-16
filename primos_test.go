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

func TestPrimo(t *testing.T) {
	for _, cc := range casesPrimos {
		if checkPrimo(cc.Value) != cc.Expected {
			t.Errorf("Expected: %v Got:%v", cc.Expected, !cc.Expected)
		}
	}
}
