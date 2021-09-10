package main

import (
	"fmt"
)

var x int = 12

func main() {

	fmt.Printf("%d\t", x)
	fmt.Printf("%b\t", x)
	fmt.Printf("%#x\n", x)

	y := x << 1
	fmt.Printf("%d\t", y) // for each one left shift decimal number will multiply by 2
	fmt.Printf("%b\t", y)
	fmt.Printf("%#x\n", y)

	z := x >> 1
	fmt.Printf("%d\t", z) // for each one left shift decimal number will divide by 2
	fmt.Printf("%b\t", z)
	fmt.Printf("%#x\n", z)

}
