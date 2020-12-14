package main

import "testing"

func TestAnysum(t *testing.T) {
	type test struct {
		data   []int //data is a slice of (type)int
		answer int
	}
	z := []test{ //type compositeliteral{...put values in here}; also z is a slice of (type)test
		test{[]int{1, 2}, 3},
		test{[]int{11, 2}, 13},
		test{[]int{9, 2}, 11},
		test{[]int{90, 2}, 92},
	}
	for _, v := range z {
		x := mySum(v.data...) //unfurling v.data bcz mySum accepts(xi ...int)
		if x != v.answer {
			t.Error("Expected", v.answer, "but  got", x)
		}
	}

}
