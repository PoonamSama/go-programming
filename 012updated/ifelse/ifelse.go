package main

import "fmt"

func main() {
	x := 10
	if x == 10 {
		fmt.Println("if executed")
	} else if x == 20 {
		fmt.Println("elseif executed")
	} else {
		fmt.Println("else executed")
	}
}
