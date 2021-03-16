package main

import "testing"

var casesOrdering = []struct {
	Value    string
	Expected string
}{
	{"juan", "ajnu"},
	{"amigo", "agimo"},
	{"atolondrado", "aaddlnooort"},
	{"la4veces 1", " 14aceelsv"},
}

func TestOrdering(t *testing.T) {
	for _, r := range casesOrdering {
		result := orderingByte(r.Value)
		if result != r.Expected {
			t.Errorf("Expected: %s Got:%s", r.Expected, result)
		}
	}
}
