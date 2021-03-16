package main

import "testing"

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

func TestReverse(t *testing.T) {
	for _, r := range casesReverse {
		result := reverse(r.Value)
		if result != r.Expected {
			t.Errorf("Expected: %s Got:%s", r.Expected, result)
		}
	}
}
