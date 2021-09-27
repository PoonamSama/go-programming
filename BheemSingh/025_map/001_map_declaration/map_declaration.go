package main

import (
	"fmt"
)

func main() {

	// using short declarartion operator
	map_1 := map[int]string{
		90: "cow",
		91: "dog",
		92: "sparrow",
		93: "cat",
	}

	fmt.Println(map_1)

	// using var keyword empty map

	var map_2 map[int]int

	if map_2 == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
