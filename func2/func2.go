package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	e, d := add(3, 4) //EVERRYTHING IN GO IS PASS BY VALUE  3,4  are arguments
	fmt.Println("sum and product is:", d, e)
}

//func( r receivers) identifier (parameters){return(s)} {...}
func add(x int, y int) (int, int) { // x and y these are parameters; int and int are returns
	c := x + y
	a := x * y

	fmt.Println("from func add here is", a, c)
	return c, a
}
