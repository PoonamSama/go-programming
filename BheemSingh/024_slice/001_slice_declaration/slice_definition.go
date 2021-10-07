package main

import "fmt"

func main() {

	// COMPOSITE literal use for declare slice, array,map etc. Syntax is X:= TYPE{VALUES}

	// A SLICE allows you to group together values of same type

	x := []int{2, 3, 4, 5, 6}

	fmt.Println(x)

	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)

	}
}
