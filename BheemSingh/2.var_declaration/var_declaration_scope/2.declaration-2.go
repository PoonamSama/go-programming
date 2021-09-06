package main

import "fmt"

// declaration of a variable out of the function body

var x int = 10

var y string = `hi we are on going to golang tour hope this will be amazing.`

var z string = "Hey, GO is all about TYPES. It is oriented about types."
var w = "Go  is static languge which reduce probability of errors."

// x:=10 will not work here this declaration is valid for func body

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(w)
}
