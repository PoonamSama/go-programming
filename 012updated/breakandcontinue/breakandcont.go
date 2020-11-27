package main

import (
	"fmt"
)

func main() {
	//a program to print even numbers only
	x := 1
	for {
		x++
		if x > 50 {
			break
		}
		if x%2 != 0 {
			continue // as soon as continue is reached, control goes back to for
		}

		fmt.Println(x)
	}
	fmt.Println("done")
}
