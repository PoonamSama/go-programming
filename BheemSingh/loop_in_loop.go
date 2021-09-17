package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 6; i++ {
		fmt.Printf("outer loop no. %d\n", i)

		for j := 0; j < 2; j++ {
			fmt.Printf("\t inner loop no. %d\n", j)
		}

	}

}
