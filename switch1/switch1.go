package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 9):
		fmt.Println("This should not print either")
	case (3 == 3):
		fmt.Println("This should  print")
		fallthrough
	case (1 == 4):
		fmt.Println("not true, does it print?")
		fallthrough
	case (3 < 1):
		fmt.Println("This should not  print")
		fallthrough
	default:
		fmt.Println("This is default case")

	}
}
