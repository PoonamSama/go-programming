package main

import (
	"fmt"
)

func main() {

	// first way to declare an array
	var x [5]int
	x[0] = 5
	x[3] = 7
	fmt.Println(x)

	// second way to declare an array

	y := [5]int{2, 3, 4, 5, 6}
	fmt.Println(y)
}
