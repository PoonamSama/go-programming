package main

import "fmt"

// Zero Value is default value when we donot assign any value to our variable.
var x int
var y string

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y) //feedback: when printing on console, this just prints a blank line and then program exits. So it will be more informative to print something like this
	fmt.Println("the default value of an empty string: ", y)

}
