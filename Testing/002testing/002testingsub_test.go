package main

import "testing"

func TestSubtraction(t *testing.T) {
	x := 8 - 7
	if x != 1 {
		t.Error("Expected 1 but got", x)
	}
}
