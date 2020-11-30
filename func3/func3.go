package main

import (
	"fmt"
)

func main() {
	foo(1, 2, 3, 4, 5, 6, 67, 8)
	fmt.Println("Hello, playground from main")
}

//func( r receivers) identifier (parameters)(return(s)) {...}
func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("For index value %d we add value %d sum is %d\n", i, v, sum)
	}
}
