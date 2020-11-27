package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 42 {
		fmt.Println("x condition satisfies")

	}
	if y := 62; y == 62 { //y has scope only here
		fmt.Println("y condition satisfies")
	}
	fmt.Println(x)
	//fmt.Println(y)
	//above line will give error if uncommented since y has scope only in the if stat it is declared

}
