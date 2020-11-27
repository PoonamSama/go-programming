package main

import (
	"fmt"
)

const (
	y = iota
)

func main() {

	x := 1
	for {
		if x > 40 {
			break
		}
		fmt.Println(x)
		x *= 3
	}
	fmt.Println("done")
	fmt.Println(y)
}
