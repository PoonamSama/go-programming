package main

import "fmt"

// Zero Value is default value when we donot assign any value to our variable.
var x int
var y string

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)

}
