package main

import (
	"fmt"
)

func main() {

	// Printing year from your birth to now using syntax  :  for{}
	year := 1996
	for {

		if year > 2021 {

			break
		} else {
			fmt.Println(year)
			year++
			continue
		}

	}
}
