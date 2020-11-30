package main

import (
	"fmt"
)

func main() {
	k := foo("Jamie", 1, 2, 3, 4, 5, 6, 67, 8, "James") // //EVERRYTHING IN GO IS PASS BY VALUE  these are arguments
	fmt.Println("The sum is:", k)
	fmt.Println("Hello, playground from main")
}

//func( r receivers) identifier (parameters)(return(s)) {...}
func foo(x ...int, s string) int { //x ...int is parameter
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("For index value %d we add value %d sum is %d\n", i, v, sum)
		fmt.Println(s)
	}
	return sum
}
