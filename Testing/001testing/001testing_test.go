package main

import "testing"

func TestAnysum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected 5 but  got", x)
	}
}
