package main

import (
	"fmt"
)

func main() {
	x := 33
	for x <= 122 {

		fmt.Printf("Decimal value is : %d\t andUTF8 : %#U\t and ASCII is: %c\n", x, x, x)
		x++
	}
}

// %#U will give utf8 code as well as the ascii character
