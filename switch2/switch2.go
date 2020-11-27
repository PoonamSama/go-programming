package main

import (
	"fmt"
)

func main() {
	s := "Bond"
	switch s {
	default:
		fmt.Println("this is default")
	case "James":
		fmt.Println("This should not print")
		fallthrough
	case "Jamie", "Bond", "lyla": //all these values will be checked
		fmt.Println("This is Bond")
	case "no":
		fmt.Println("This should not print")

	}
}
