package main

import (
	"fmt"
)

func main() {
	fmt.Println("1+2+3=", mySum(1, 2, 3))
	fmt.Println("22+2+3=", mySum(22, 2, 3))
	fmt.Println("1+1+3=", mySum(1, 1, 3))
}
func mySum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
